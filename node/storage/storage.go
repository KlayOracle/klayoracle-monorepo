package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/protobuf/encoding/protojson"
)

//var Conn *pgx.Conn

//var ConnCtx context.Context

//func CreateDBConn() *pgx.Conn {
//	dns := os.Getenv("COCKROACH_DNS_URL")
//	ConnCtx = context.Background()
//	conn, err := pgx.Connect(ConnCtx, dns)
//	if err != nil {
//		config.Loaded.Logger.Fatalw("failed to connect database", err)
//	}
//
//	return conn
//}

//func dns() string {
//	dns := os.Getenv("COCKROACH_DNS_URL")
//}

//func conn() *pgx.Conn {
//	connCtx := context.Background()
//	conn, err := pgx.Connect(connCtx, os.Getenv("COCKROACH_DNS_URL"))
//	defer conn.Close(connCtx)
//	if err != nil {
//		config.Loaded.Logger.Fatalw("failed to connect database", err)
//	}
//
//	defer func() {
//		err := conn.Close(connCtx)
//		if err != nil {
//			log.Fatal("cannot close klaytn client conn: ", err)
//		}
//	}()
//
//	return conn
//}

func StoreJob(adapter *protonode.Adapter) error {
	connCtx := context.Background()
	conn, err := pgx.Connect(connCtx, os.Getenv("COCKROACH_DNS_URL"))
	defer conn.Close(connCtx)
	if err != nil {
		config.Loaded.Logger.Warnw("failed to connect database", err)
	}

	jsonData, err := protojson.Marshal(adapter)
	if err != nil {
		config.Loaded.Logger.Warnw("cannot store job info", "reason", err)
	}

	id := uuid.New()

	err = crdbpgx.ExecuteTx(connCtx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(connCtx,
			"INSERT INTO node_jobs (id, adapter_id, oracle_address, request, period) VALUES ($1, $2, $3, $4, $5)", id, adapter.AdapterId, adapter.OracleAddress, string(jsonData), time.Now()); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		config.Loaded.Logger.Warn(err)
	}

	return nil
}

func StoreResponses(adapter *protonode.Adapter, responses []int64, roundAnswer int64) error {
	connCtx := context.Background()
	conn, err := pgx.Connect(connCtx, os.Getenv("COCKROACH_DNS_URL"))
	defer conn.Close(connCtx)
	if err != nil {
		config.Loaded.Logger.Warnw("failed to connect database", err)
	}

	id := uuid.New()

	err = crdbpgx.ExecuteTx(connCtx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(connCtx,
			"INSERT INTO node_responses (id, responses, round_answer, adapter_id, period) VALUES ($1, $2, $3, $4, $5)", id, responses, roundAnswer, adapter.AdapterId, time.Now()); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		config.Loaded.Logger.Warnw("error storing aggregated responses", "feed", adapter.Name, "err", err)
	}

	return nil
}

func FetchResponsesFromT0(adapter *protonode.Adapter, t0 time.Time) (pgx.Rows, error) {
	connCtx := context.Background()
	conn, err := pgx.Connect(connCtx, os.Getenv("COCKROACH_DNS_URL"))
	defer conn.Close(connCtx)
	if err != nil {
		config.Loaded.Logger.Warnw("failed to connect database", err)
	}

	query := fmt.Sprintf("SELECT responses from node_responses WHERE adapter_id='%s' AND CAST(period as TIMESTAMP) >= '%s'", adapter.AdapterId, t0.Format("2006-01-02 15:04:05.840"))

	rows, err := conn.Query(connCtx, query)
	if err != nil {
		config.Loaded.Logger.Warnw("error fetching TWAP", "feed", adapter.Name, "err", err)

		return nil, err
	}

	return rows, nil
}

func FeedHistory(req *protonode.FeedHistoryRequest) (pgx.Rows, error) {
	connCtx := context.Background()
	conn, err := pgx.Connect(connCtx, os.Getenv("COCKROACH_DNS_URL"))
	defer conn.Close(connCtx)
	if err != nil {
		config.Loaded.Logger.Warnw("failed to connect database", err)
	}

	t0 := time.Unix(int64(req.FromTimestamp), 0).Format("2006-01-02 15:04:05.840")
	t1 := time.Unix(int64(req.ToTimestamp), 0).Format("2006-01-02 15:04:05.840")

	query := fmt.Sprintf("SELECT responses,round_answer,adapter_id,period from node_responses WHERE adapter_id='%s' AND CAST(period as TIMESTAMP) >= '%s' AND CAST(period as TIMESTAMP) <= '%s'", req.AdapterId, t0, t1)

	rows, err := conn.Query(connCtx, query)
	if err != nil {
		config.Loaded.Logger.Warnw("error fetching feed history", "adapter_id", req.AdapterId, "err", err)

		return nil, err
	}

	return rows, nil
}

func FetchRoundSize(adapter *protonode.Adapter) (pgx.Rows, error) {
	connCtx := context.Background()
	conn, err := pgx.Connect(connCtx, os.Getenv("COCKROACH_DNS_URL"))
	defer conn.Close(connCtx)
	if err != nil {
		config.Loaded.Logger.Warnw("failed to connect database", err)
	}

	query := fmt.Sprintf("SELECT COUNT(*) as count from node_responses WHERE adapter_id='%s'", adapter.AdapterId)

	rows, err := conn.Query(connCtx, query)
	if err != nil {
		config.Loaded.Logger.Warnw("error fetching round size", "feed", adapter.Name, "err", err)
	}

	return rows, nil
}

func FetchRoundSizeAll(adapters []string) (pgx.Rows, error) {
	connCtx := context.Background()
	conn, err := pgx.Connect(connCtx, os.Getenv("COCKROACH_DNS_URL"))
	defer conn.Close(connCtx)
	if err != nil {
		config.Loaded.Logger.Warnw("failed to connect database", err)
	}

	var fetch string

	for index, adapter := range adapters {
		if index == 0 {
			fetch += "('"
		} else {
			fetch += "'"
		}

		fetch += string(adapter)

		if index+1 == len(adapters) {
			fetch += "')"
		} else {
			fetch += "', "
		}
	}

	query := fmt.Sprintf("SELECT COUNT(*) as count from node_responses WHERE adapter_id in %s", fetch)

	fmt.Println(query)
	rows, err := conn.Query(connCtx, query)
	if err != nil {
		config.Loaded.Logger.Warnw("error fetching round size", "feeds", adapters, "err", err)
	}

	return rows, nil
}
