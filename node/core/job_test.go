package core

import (
	"os"
	"path"
	"testing"

	"github.com/klayoracle/klayoracle-monorepo/node/boot"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/protobuf/encoding/protojson"
)

var byt = []byte(`
				{
				"active": true,
				"name": "KLAY_USD",
				"job_type": "DATA_FEED",
				"adapter_id":"0x8b7460cccfa0aca303ee85c3fb81c344faad2fbab415adc32b2984008b7efd76",
				"oracle_address": "0x7055108d9B5AB0765191A570f8F22A4ccEffd950",
				"category": 2,
				"frequency": 30000000000,
				"feeds": [
						{
						"url": "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=KLAY&tsyms=USD",
						"request_type": 0,
						"headers": [{"field": {"Content-Type" : "application/json"}},{"field": {"Authorization" : "${BEARER_TOKEN}"}}],
						"reducers": [{"function": "PARSE","args": ["$.RAW.KLAY.USD.PRICE"]},{"function": "FLOAT64_MUL_UINT64","args": ["1000000000"]}],
      					"payload": ""
						}
					]
				}
				`)

func TestJob(t *testing.T) {

	wd, _ := os.Getwd()
	boot.Boot(path.Join(wd, ".."), path.Join(wd, "..", "config.yaml"), path.Join(wd, "..", ".env"))

	newAdapter := protonode.Adapter{}

	stream := []byte(os.ExpandEnv(string(byt)))
	protojson.Unmarshal(stream, &newAdapter)

	Run(newAdapter)
}

func TestUpdateRoundAnswer(t *testing.T) {
	wd, _ := os.Getwd()
	boot.Boot(path.Join(wd, ".."), path.Join(wd, "..", "config.yaml"), path.Join(wd, "..", ".env"))

	NewWssClient()

	newAdapter := protonode.Adapter{}

	stream := []byte(os.ExpandEnv(string(byt)))
	protojson.Unmarshal(stream, &newAdapter)

	//DeployNewOracleProviderSample(os.Getenv("PUBLIC_ADDRESS"), "0x8b7460cccfa0aca303ee85c3fb81c344faad2fbab415adc32b2984008b7efd76")
	//err, hash := UpdateRoundAnswer(&newAdapter, 10)
	//
	//fmt.Println("Transaction hash: ", hash.String())
	//
	//assert.Nil(t, err, err)
}
