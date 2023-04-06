package job

import (
	"github.com/klayoracle/klayoracle-monorepo/node/boot"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
	"path"
	"testing"
)

func TestJob(t *testing.T) {

	wd, _ := os.Getwd()
	boot.Boot(path.Join(wd, ".."), path.Join(wd, "..", "config.yaml"), path.Join(wd, "..", ".env"))

	newAdapter := protonode.Adapter{}

	byt := []byte(`
				{
				"active": true,
				"name": "KLAY_USD",
				"job_type": "DATA_FEED",
				"adapter_id":"0x8b7460cccfa0aca303ee85c3fb81c344faad2fbab415adc32b2984008b7efd76",
				"oracle_address": "0xCC4377b912c4517Fe895817c6a7c6937D92A70B3",
				"category": 2,
				"frequency": 30000000000,
				"feeds": [
						{
						"url": "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=KLAY&tsyms=USD",
						"request_type": 0,
						"headers": [{"field": {"Content-Type" : "application/json"}},{"field": {"Authorization" : "${BEARER_TOKEN}"}}],
						"reducers": [{"function": "PARSE","args": ["$.RAW.KLAY.USD.PRICE"]},{"function": "MUL","args": ["1000000000"]}],
      					"payload": "{\"type\":\"limit\",\"side\":\"buy\",\"price\":1.058e-9,\"size\":100,\"currency\":[\"USDT_BTC\",\"WEMIX_KLAY\",\"KLAY_USDT\"]}"
						},
						{
						"url": "https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
						"request_type": 0,
						"headers": [{"field": {"X-CoinAPI-Key": "${X_COIN_API_KEY}"}}],
						"reducers": [{"function": "PARSE","args": ["$.rate"]},{"function": "MUL","args": ["1000000000"]}]
						}
					]
				}
				`)

	stream := []byte(os.ExpandEnv(string(byt)))
	protojson.Unmarshal(stream, &newAdapter)

	Run(newAdapter)
}
