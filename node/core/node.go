package core

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/pborman/uuid"

	"github.com/klaytn/klaytn/common"

	"github.com/jackc/pgtype"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/adapter"
	bootstrap2 "github.com/klayoracle/klayoracle-monorepo/node/bootstrap"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/klayoracle/klayoracle-monorepo/node/storage"

	"golang.org/x/exp/slices"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/bootstrap"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const TxTimeout = time.Minute * 1

var (
	errMissingMetadata            = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken               = status.Errorf(codes.Unauthenticated, "invalid token")
	errDataProviderNotWhitelisted = status.Errorf(codes.Unknown, "you need to be whitelisted")
	errAddingToKnownPeer          = status.Errorf(codes.Unknown, "cannot add to known DP known peers")
)

type RoundQueue struct {
	adapter *protonode.Adapter
	answer  int64
}

type Node struct {
	protonode.UnimplementedNodeServiceServer
	Sever         *grpc.Server
	Organization  config.Organization
	dataProviders map[string]*protonode.DPInfo
	mu            sync.Mutex
	Jobs          map[string]*protonode.Adapter  //Using multiple routines to R&W to channel is safe without mutex.
	bootstraps    map[string]*protonode.NodeInfo // A list of 3 default bootstrap node nodes
	knownPeers    map[string]*protonode.NodeInfo
	listenAddress string
	LastTxHash    string
	LastNonce     *big.Int
	RoundQueue    *[]RoundQueue
}

func (n *Node) HandShake(ctx context.Context, provider *protonode.DPInfo) (*protonode.HandShakeStatus, error) {

	//@Todo confirm organization is whitelisted in DB
	//Node runner uses binary/make command to add and remove dp from serving organization
	if isDPWhitelist(provider) != true {
		return nil, errDataProviderNotWhitelisted
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	config.Loaded.Logger.Infow("attempting to add provider to all bootstrap dp known peers ", "listen address", provider.ListenAddress)

	//if slices.Contains(bootstrap.Nodes(), provider.ListenAddress) {
	//	config.Loaded.Logger.Warnw("ignoring adding to known peer", "reason", "bootstrap dp", "listen address", provider.ListenAddress)
	//} else {
	err := addDataProviderPeer(provider)

	if err != nil {
		config.Loaded.Logger.Fatal("error adding to known peer: ", err)
		return nil, errAddingToKnownPeer
	}
	//}

	config.Loaded.Logger.Infow("registering provider with service node ", "provider", provider.ListenAddress, "service node", os.Getenv("HOST_IP"))

	n.dataProviders[provider.ListenAddress] = provider

	//Check DP Peers to find out Org is not yet registered
	return &protonode.HandShakeStatus{Status: true, ErrorMsg: ""}, nil
}

func (n *Node) AddToKnownPeers(ctx context.Context, info *protonode.NodeInfo) (*protonode.Null, error) {

	config.Loaded.Logger.Info("adding ", info.ListenAddress, " to known peers of bootstrap node ", n.listenAddress)

	n.mu.Lock()
	defer n.mu.Unlock()

	n.knownPeers[info.ListenAddress] = info

	return new(protonode.Null), nil
}

func (n *Node) ListKnownPeers(context.Context, *protonode.Null) (*protonode.NodeInfos, error) {

	nds := &protonode.NodeInfos{}

	for _, nd := range n.knownPeers {
		nds.List = append(nds.List, &protonode.NodeInfo{
			ListenAddress: nd.ListenAddress,
			Name:          nd.Name,
			KOrgId:        nd.KOrgId,
			Website:       nd.Website,
			Uptime:        time.Now().UnixNano(),
		})
	}

	return nds, nil
}

func (n *Node) Peer() {
	p := &protonode.NodeInfo{
		ListenAddress: n.listenAddress,
		Name:          n.Organization.Name,
		KOrgId:        n.Organization.ID,
		Website:       n.Organization.Website,
		KnownPeers:    map[string]*protonode.NodeInfo{},
		Bootstraps:    map[string]*protonode.NodeInfo{},
	}

	//bn := bootstrap2.BT{Addr: p.ListenAddress, OrgID: p.KOrgId, Domain: p.Website, Name: p.Name}

	//if slices.Contains(bootstrap2.Nodes(), bn) {
	//	config.Loaded.Logger.Warnw("ignoring adding to known peer", "reason", "bootstrap node", "listen address", p.ListenAddress)

	//index := slices.Index(bootstrap2.Nodes(), bn)
	//
	//otherBts := slices.Delete(bootstrap2.Nodes(), index, index+1)
	//
	//for _, bt := range otherBts {
	//	err := addNodePeer(&protonode.NodeInfo{
	//		ListenAddress: bt.Addr,
	//		Name:          bt.Name,
	//		KOrgId:        bt.OrgID,
	//		Website:       bt.Domain,
	//		KnownPeers:    map[string]*protonode.NodeInfo{},
	//		Bootstraps:    map[string]*protonode.NodeInfo{},
	//	})
	//	if err != nil {
	//		config.Loaded.Logger.Warnw("node peering failed", "error", err)
	//	}
	//}
	//} else {
	err := addNodePeer(p)
	if err != nil {
		config.Loaded.Logger.Warnw("node peering failed", "error", err)
	}

	//}
}

func (n *Node) QueueJob(ctx context.Context, adapter *protonode.Adapter) (*protonode.RequestStatus, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	if len(md["provider"]) > 0 {
		provider := md["provider"][0]

		_, ok := n.dataProviders[provider]

		if ok {
			config.Loaded.Logger.Infow("new queue request", "data provider", provider, "request", adapter)
			n.mu.Lock()
			defer n.mu.Unlock()

			jobId := uuid.New()

			n.Jobs[jobId] = adapter

			err := storage.StoreJob(adapter)
			if err != nil {
				config.Loaded.Logger.Warnw("error storing job", "err", err)
			}

			config.Loaded.Logger.Infow("job in queue: ", "jobId", jobId)
			config.Loaded.Logger.Infow("total job in queue: ", "size", len(n.Jobs))

			go func() {
				n.execJob(adapter, jobId)
			}()
		} else {
			config.Loaded.Logger.Warnw("adapter has not registered an handshake", "service node",
				os.Getenv("HOST_IP"), "data provider", provider)

			return &protonode.RequestStatus{
				Status: 1,
			}, fmt.Errorf("adapter has not registered an handshake")
		}

	}

	return &protonode.RequestStatus{
		Status: 0,
	}, nil
}

func (n *Node) execJob(adapter *protonode.Adapter, jobId string) {
	defer func() {
		delete(n.Jobs, jobId)
	}()

	dp := &protonode.Adapter{}

	err := castBtwDPInfo(adapter, dp)
	if err != nil {
		config.Loaded.Logger.Warnw("conversion between protoadapter.Adapter{} and protonode.Adapter{}  gone wrong", "error", err)
	}

	roundAnswer, err := Run(*dp)

	if err != nil {
		config.Loaded.Logger.Warnw("dropping job, an error occurred", "error", err)
	} else if roundAnswer == 0 {
		config.Loaded.Logger.Warnw("cannot update oracle with 0 dropping", "oracle", adapter.OracleAddress)
	} else {
		config.Loaded.Logger.Infow("updating oracle with round answer", "oracle", adapter.OracleAddress, "answer", roundAnswer)

		// We need to queue round due to issue arising from submitting parallel transaction
		// causing nonce to clash, and keep track of nonce then doing +1 can block all ahead txs if a lower tx fails.
		// Solution:
		// - Add round details to a queue on the node using mutex lock to ensure concurrency
		// - Keep checking queue for new transaction hash
		// - If its empty, log new time, pick next round using FIFO
		// - Send round details to Oracle contract, and update last transaction hash
		// - Keep checking if last transaction using its hash is done or its timeout (i.e. More than {TxTimeout} in sec)
		// - If transaction is taking longer than {TxTimeout} secs, send transaction with same nonce and higher gas to cancel it (send 0 token to null address with higher gas).
		// - If transaction is mined clear last hash & last time

		n.queueRound(adapter, roundAnswer)
	}

}

func (n *Node) queueRound(adapter *protonode.Adapter, roundAnswer int64) {
	n.mu.Lock()
	defer n.mu.Unlock()

	*n.RoundQueue = append(*n.RoundQueue, RoundQueue{
		adapter: adapter,
		answer:  roundAnswer,
	})
}

func (n *Node) WatchRoundQueue() {

	//If queue is free, and new answer is available
	//update round on oracle contract
	if n.LastTxHash == "" && len(*n.RoundQueue) > 0 {
		adapter := (*n.RoundQueue)[0].adapter
		roundAnswer := (*n.RoundQueue)[0].answer

		err, hash := UpdateRoundAnswer(adapter, roundAnswer)

		if err != nil || hash.String() == "0x0000000000000000000000000000000000000000000000000000000000000000" {
			config.Loaded.Logger.Warn(err)
		} else {
			n.LastTxHash = hash.String()

			for {
				timeout := <-time.After(TxTimeout)

				config.Loaded.Logger.Infow("wait time for round is up, checking transaction", "hash", n.LastTxHash, "after", timeout.Local())

				client, err := NewHttpClient()
				go func() {
					defer client.Close()
				}()

				if err != nil {
					config.Loaded.Logger.Fatalw("https connection to blockchain client failed", "error", err)
				} else {

					transaction, pending, err := client.TransactionByHash(KlaytnClientCtx, common.HexToHash(n.LastTxHash))

					if err != nil {
						config.Loaded.Logger.Warnw("error fetching transaction for round answer, transaction probably failed", "hash", n.LastTxHash, "error", err)

						//@Todo cover more reasons for this case. e.g client connection error e.t.c
						nonce, err := client.NonceAt(KlaytnClientCtx, common.HexToAddress(""), nil)
						if err != nil {
							config.Loaded.Logger.Fatal("cannot determine nonce for node address: %v", err)
						} else {
							n.LastNonce = big.NewInt(int64(nonce))
							n.LastTxHash = ""
						}

					} else if pending == true {
						config.Loaded.Logger.Warnw("transaction to update round answer is taking longer than TxTimeout configured, canceling trx to free nonce", "hash", n.LastTxHash)

						//@Todo
						//Cancel pending transaction
						//Update new nonce +1
					} else {
						config.Loaded.Logger.Infow("successful transaction response from oracle", "transaction", transaction)

						n.mu.Lock()

						n.LastNonce = n.LastNonce.Add(n.LastNonce, big.NewInt(1))
						n.LastTxHash = ""

						*n.RoundQueue = (*n.RoundQueue)[1:]

						n.mu.Unlock()
					}
				}

				return
			}
		}

	}
}

func addNodePeer(p *protonode.NodeInfo) error {

	//Refer to addDataProviderPeer(p *protonode.DPInfo) function comment for better understanding
	var (
		longestChain     = 0                                 //Count of the longest chain
		longestChainNode = ""                                //Listen Ip of the longest Node's peer-chain
		peerList         = map[string]*protonode.NodeInfos{} //Mapping of bootstrap node listen address to their node peers
		peers            = &protonode.NodeInfos{}            //
		err              error
	)

	bootstraps := bootstrap2.Nodes()

	for _, bt := range bootstraps {

		config.Loaded.Logger.Info("fetching known peers for bootstrap node: ", bt.Addr)
		peers, err = getPeerListNode(bt)
		if err != nil {
			config.Loaded.Logger.Warnw("bootstrap node is unresponsive on ", "node", bt.Addr, "error", err)
		} else {
			peerList[bt.Addr] = peers

			if len(peers.List) >= longestChain {
				longestChain = len(peers.List)
				longestChainNode = bt.Addr
			}

			config.Loaded.Logger.Info("peers found for ", bt.Addr)
			if len(peers.List) > 0 {
				config.Loaded.Logger.Info(peers)
			} else {
				config.Loaded.Logger.Warn("no known peers discovered")
			}
		}
	}

	list := append(peerList[longestChainNode].List, p)

	for _, bt := range bootstraps {

		//Due to defer behaviour in loop, we do this instead
		err := func() error {

			//Bootstrap node can't add self as peer
			client, conn, err := NewNodeServiceClient(bt)
			if err != nil {
				return fmt.Errorf("failed connecting to client: %v", err)
			}

			sbt, _ := json.Marshal(bt)
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) //Enough time to authenticate and add DP to known peers
			md := metadata.Pairs("node", string(sbt))
			ctx = metadata.NewOutgoingContext(ctx, md)

			defer cancel()
			defer func() {
				defer conn.Close()
			}()

			//Add all node services in list as this bootstrap node know peer
			for _, newNd := range list {
				_, err = client.AddToKnownPeers(ctx, newNd)

				if err != nil {
					config.Loaded.Logger.Warnw("node.AddToKnownPeers(_) failed", "reason", err, "skipping", bt.Addr)
				} else {
					config.Loaded.Logger.Infow("node registered to bootstrap node known peer ", "boostrap node", bt.Addr, "added peer", newNd.ListenAddress)
				}
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func addDataProviderPeer(p *protonode.DPInfo) error {

	//Send a grpc request to all bootstrap nodes for Known peers
	//Identify the bootstrap dp node with the longest known peers chain
	//Add new dp Info to its peers and loop through the all bootstrap nodes to update the list of known peers

	//Gather known peers from all bootstrap dp node
	//Identify the longest peer chain
	var (
		longestChain   = 0                                  //Count of the longest chain
		longestChainDP = ""                                 //Listen Ip of the longest DP's peer-chain
		peerList       = map[string]*protoadapter.DPInfos{} //Mapping of bootstrap dp listen address to their dp peers
		peers          = &protoadapter.DPInfos{}            //
		err            error
	)

	bootstraps := bootstrap.Nodes()

	//@Todo If a bootstrap node does not respond, it should not stop the service, just warn and move on
	for _, listenAddr := range bootstraps {
		config.Loaded.Logger.Info("fetching known peers for bootstrap DP: ", listenAddr)
		peers, err = getPeerListDP(listenAddr)
		if err != nil {
			config.Loaded.Logger.Warn("bootstrap DP is unresponsive on ", listenAddr)
		} else {

			peerList[listenAddr] = peers

			if len(peers.List) >= longestChain {
				longestChain = len(peers.List)
				longestChainDP = listenAddr
			}

			config.Loaded.Logger.Info("peers found for ", listenAddr)
			if len(peers.List) > 0 {
				config.Loaded.Logger.Info(peers)
			} else {
				config.Loaded.Logger.Warn("no known peers discovered")
			}
		}
	}

	config.Loaded.Logger.Infow("DP bootstrap with the longest peer", longestChainDP, longestChain)

	dp := &protoadapter.DPInfo{}

	err = castBtwDPInfo(p, dp)

	if err != nil {
		config.Loaded.Logger.Warn("conversion between DPInfo gone wrong: ", err)

		return err
	} else {

		list := append(peerList[longestChainDP].List, dp)

		for _, bt := range bootstraps {

			//Due to defer behaviour in loop, we do this instead
			err := func() error {

				//Bootstrap node can't add self as peer
				client, _, err := adapter.NewDPServiceClient(bt)
				if err != nil {
					return fmt.Errorf("failed connecting to client: %v", err)
				}

				ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) //Enough time to authenticate and add DP to known peers
				defer cancel()

				//Add all DP services in list as this bootstrap DP know peer
				for _, newDp := range list {
					_, err = client.AddToKnownPeers(ctx, newDp)

					if err != nil {
						config.Loaded.Logger.Warnw("dp.AddToKnownPeers(_) failed", "reason", err, "skipping", bt)
					} else {
						config.Loaded.Logger.Infow("Added dp to bootstrap dp known peer ", "boostrap dp", bt, "added peer", newDp.ListenAddress)
					}
				}

				return nil
			}()

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getPeerListDP(listenAddr string) (*protoadapter.DPInfos, error) {

	client, _, err := adapter.NewDPServiceClient(listenAddr)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Enough time to authenticate and add DP to known peers
	defer cancel()

	res, err := client.ListKnownPeers(ctx, &protoadapter.Null{})

	if err != nil {
		return nil, fmt.Errorf("dp.ListKnowPeers(_) failed: %v", err)
	}

	return res, nil
}

func getPeerListNode(bt bootstrap2.BT) (*protonode.NodeInfos, error) {
	client, conn, err := NewNodeServiceClient(bt)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to client: %v", err)
	}

	sbt, _ := json.Marshal(bt)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Enough time to authenticate and add DP to known peers
	md := metadata.Pairs("node", string(sbt))
	ctx = metadata.NewOutgoingContext(ctx, md)

	defer cancel()
	defer func() {
		defer conn.Close()
	}()

	res, err := client.ListKnownPeers(ctx, &protonode.Null{})

	if err != nil {
		return nil, fmt.Errorf("node.ListKnowPeers(_) failed: %v", err)
	}

	return res, nil
}

// NewNodeServiceClient
// @Todo this can only be used by bootstrap node to create client conn. other node will fail with auth err
func NewNodeServiceClient(bt bootstrap2.BT) (protonode.NodeServiceClient, *grpc.ClientConn, error) {
	oauthKey := os.Getenv("OAUTH_TOKEN")

	//@Todo, use boostrap oAuth
	rpcCredentials := oauth.TokenSource{TokenSource: adapter.OauthTokenSource{
		AuthKey: oauthKey,
	}}

	nodeCertFilePath := path.Join(config.Loaded.RootWD, "certs/bootstraps/x509", fmt.Sprintf("%s.pem", bt.OrgID))

	credentials, err := credentials.NewClientTLSFromFile(nodeCertFilePath, bt.Domain)
	if err != nil {
		return nil, &grpc.ClientConn{}, fmt.Errorf("failed to load credentials: %v", err)
	}
	var opts []grpc.DialOption

	opts = []grpc.DialOption{
		grpc.WithPerRPCCredentials(rpcCredentials),
		grpc.WithTransportCredentials(credentials),
	}

	conn, err := grpc.Dial(bt.Addr, opts...)

	if err != nil {
		return nil, &grpc.ClientConn{}, fmt.Errorf("did not connect: %v", err)
	}

	return protonode.NewNodeServiceClient(conn), conn, nil
}

func NewNodeServiceServer(n *Node) (*grpc.Server, error) {
	certFilePath := path.Join(config.Loaded.RootWD, "certs/host", config.Loaded.SSL.Certificate)
	keyFilePath := path.Join(config.Loaded.RootWD, "certs/host", config.Loaded.SSL.Key)

	cert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return new(grpc.Server), fmt.Errorf("failed to load credentials: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)

	n.dataProviders = make(map[string]*protonode.DPInfo)
	n.Sever = s
	n.Organization = config.Loaded.Organization
	n.listenAddress = os.Getenv("HOST_IP")
	n.knownPeers = make(map[string]*protonode.NodeInfo)
	n.bootstraps = make(map[string]*protonode.NodeInfo)

	config.Loaded.Logger.Info("node organization: ", n.Organization)

	for _, bt := range bootstrap2.Nodes() {
		p := &protonode.NodeInfo{
			ListenAddress: bt.Addr,
			KOrgId:        bt.OrgID,
			Website:       bt.Domain,
			KnownPeers:    map[string]*protonode.NodeInfo{},
			Bootstraps:    map[string]*protonode.NodeInfo{},
		}
		n.bootstraps[p.ListenAddress] = p
	}

	protonode.RegisterNodeServiceServer(s, n)

	config.Loaded.Logger.Info("starting Node service on ", os.Getenv("HOST_IP"))

	return s, nil
}

// valid validates the authorization.
func valid(authorization []string) bool {

	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	return token == os.Getenv("OAUTH_TOKEN")
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	if md["node"] != nil {
		sbt := md["node"][0]

		if sbt != "" {
			var bt bootstrap2.BT

			json.Unmarshal([]byte(sbt), &bt)

			if slices.Contains(bootstrap2.Nodes(), bt) {
				return handler(ctx, req)
			}
		}
	}

	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

func isDPWhitelist(dpi *protonode.DPInfo) bool {
	return true
}

func castBtwDPInfo(from, to interface{}) error {
	b, err := json.Marshal(from)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, to)
}

func (n *Node) NodeJobsCount(ctx context.Context, null *protonode.Null) (*protonode.JobLength, error) {

	var roundSize int

	for _, dp := range n.dataProviders {
		client, conn, err := adapter.NewDPServiceClient(dp.ListenAddress)
		if err != nil {
			config.Loaded.Logger.Warnw("error in NodeJobsCount", "error", err)
			return nil, err
		}

		feeds, err := client.ListFeeds(ctx, &protoadapter.Null{})
		if err != nil {
			config.Loaded.Logger.Warnw("error in NodeJobsCount", "error", err)
			return nil, err
		}

		var ids []string

		for _, feed := range feeds.Adapters {
			ids = append(ids, feed.AdapterId)
		}

		rows, err := storage.FetchRoundSizeAll(ids)
		if err != nil {
			config.Loaded.Logger.Warnw("error in NodeJobsCount", "error", err)
			return nil, err
		}

		var len int

		for rows.Next() {
			rows.Scan(&len)
		}

		roundSize += len

		conn.Close()
	}

	return &protonode.JobLength{Value: int64(roundSize)}, nil
}

func (n *Node) TotalNetworkRequests(ctx context.Context, null *protonode.Null) (*protonode.JobLength, error) {
	sumJobs := &protonode.JobLength{}

	for _, peer := range n.knownPeers {
		bt := bootstrap2.BT{
			Addr:   peer.ListenAddress,
			Name:   peer.Name,
			OrgID:  peer.KOrgId,
			Domain: peer.Website,
		}

		//@Todo only bootstrap Nodes can be accessed with known cert file
		if !slices.Contains(bootstrap2.Nodes(), bt) {
			continue
		}

		client, conn, err := NewNodeServiceClient(bt)
		if err != nil {
			return nil, fmt.Errorf("failed connecting to client: %v", err)
		}

		sbt, _ := json.Marshal(bt)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Enough time to authenticate and add DP to known peers
		md := metadata.Pairs("node", string(sbt))
		ctx = metadata.NewOutgoingContext(ctx, md)

		reqs, err := client.NodeJobsCount(ctx, &protonode.Null{})
		if err != nil {
			return nil, err
		}

		sumJobs.Value = sumJobs.Value + reqs.Value

		cancel()
		conn.Close()
	}

	return sumJobs, nil
}

func (n *Node) FeedHistory(ctx context.Context, req *protonode.FeedHistoryRequest) (*protonode.FeedHistories, error) {

	histories := &protonode.FeedHistories{}

	rows, err := storage.FeedHistory(req)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		rowValues, err := rows.Values()
		if err != nil {
			config.Loaded.Logger.Warnw("error returning feed history", "error", err)
			return nil, err
		}

		var rs []int64

		for _, element := range rowValues[0].(pgtype.Int8Array).Elements {
			rs = append(rs, element.Int)
		}

		history := &protonode.FeedHistory{
			Responses:   rs,
			RoundAnswer: uint64(rowValues[1].(int64)),
			AdapterId:   rowValues[2].(string),
			Period:      uint64(rowValues[3].(time.Time).Unix()),
		}

		histories.Feed = append(histories.Feed, history)
	}

	return histories, nil
}
