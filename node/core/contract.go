package core

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/klayoracle/klayoracle-monorepo/node/config"
	"github.com/klayoracle/klayoracle-monorepo/node/contracts/oracle"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
)

func nodeSigningKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		config.Loaded.Logger.Fatalw("error with node signing key", "error", err)
	}

	return privateKey
}

func getNonce(address common.Address) uint64 {
	nonce, err := KlaytnClient.NonceAt(KlaytnClientCtx, address, nil)
	if err != nil {
		config.Loaded.Logger.Warnw("cannot determine nonce", "address", address.String())
	}

	return nonce
}

func getAddress(privateKey *ecdsa.PrivateKey) common.Address {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		config.Loaded.Logger.Warnw("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func UpdateRoundAnswer(adapter *protonode.Adapter, roundAnswer int64) (error, common.Hash) {
	oracleAddress := common.HexToAddress(adapter.OracleAddress)

	oracleDP, err := oracle.NewOracleProviderSample(oracleAddress, KlaytnClient)
	if err != nil {
		config.Loaded.Logger.Warnw("cannot instantiate oracle", "error", err)
	} else {

		gasPrice, _ := KlaytnClient.SuggestGasPrice(KlaytnClientCtx)
		nodePrivateKey := nodeSigningKey()
		nodeAddr := getAddress(nodePrivateKey)
		transactor := bind.NewKeyedTransactor(nodePrivateKey)

		//transactor.Nonce = big.NewInt(int64(getNonce(nodeAddr)))
		transactor.Value = big.NewInt(0)
		transactor.GasPrice = gasPrice

		roundTime := time.Now().UnixMilli()

		hash := crypto.Keccak256Hash(
			common.HexToAddress(nodeAddr.Hex()).Bytes(),
			common.LeftPadBytes(big.NewInt(roundAnswer).Bytes(), 32),
			common.LeftPadBytes(big.NewInt(roundTime).Bytes(), 32),
		)

		prefixedHashMsg := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Klaytn Signed Message:\n%d%s", len(hash), hash)))
		signature, _ := crypto.Sign(prefixedHashMsg.Bytes(), nodePrivateKey)

		var (
			roundAnswerByte32 [32]byte
		)

		copy(roundAnswerByte32[:], common.LeftPadBytes(big.NewInt(roundAnswer).Bytes(), 32))

		if err != nil {
			msg := "unable to sign round data"
			config.Loaded.Logger.Warnw(msg, "error", err)

			return fmt.Errorf(msg), common.Hash{}
		} else {
			contractABI, _ := abi.JSON(strings.NewReader(oracle.KlayOracleABI))
			payload, err := contractABI.Pack("newRoundData", big.NewInt(roundTime), roundAnswerByte32, signature)

			if err != nil {
				msg := "unable to pack round data"

				config.Loaded.Logger.Warnw(msg, "error", err)

				return fmt.Errorf(msg), common.Hash{}
			} else {

				gasEstimate, err := KlaytnClient.EstimateGas(KlaytnClientCtx, ethereum.CallMsg{
					From:     nodeAddr,
					To:       &oracleAddress,
					Data:     payload,
					GasPrice: gasPrice,
				})
				if err != nil {
					msg := "unable to estimate gas"
					config.Loaded.Logger.Warnw(msg, "error", err)

					return fmt.Errorf(msg), common.Hash{}
				} else {

					auth := bind.NewKeyedTransactor(nodePrivateKey)
					//auth.Nonce = big.NewInt(int64(getNonce(nodeAddr)))
					auth.GasLimit = gasEstimate
					auth.Context = KlaytnClientCtx
					auth.GasPrice = gasPrice

					trx, err := oracleDP.NewRoundData(auth, big.NewInt(roundTime), roundAnswerByte32, signature)
					if err != nil {
						config.Loaded.Logger.Warnw("error while updating round data", "error", err)
					} else {
						config.Loaded.Logger.Infow("successfully updated round data", "adapter", adapter.AdapterId, "transaction", trx.Hash())
					}

					return nil, trx.Hash()
				}
			}
		}

	}

	return fmt.Errorf("unknown error"), common.Hash{}
}

func DeployNewOracleProviderSample(nodeAddress, adapterId string) {
	nodePrivateKey := nodeSigningKey()
	nodeAddr := getAddress(nodePrivateKey)
	gasPrice, _ := KlaytnClient.SuggestGasPrice(KlaytnClientCtx)

	auth := bind.NewKeyedTransactor(nodePrivateKey)
	auth.Nonce = big.NewInt(int64(getNonce(nodeAddr)))
	auth.GasLimit = uint64(3800000)
	auth.Context = KlaytnClientCtx
	auth.GasPrice = gasPrice

	var adapter32 [32]byte

	copy(adapter32[:], adapterId)

	addr, tx, _, err := oracle.DeployOracleProviderSample(auth, KlaytnClient, common.HexToAddress(nodeAddress), adapter32)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(addr.Bytes()))
	fmt.Println(tx)
}
