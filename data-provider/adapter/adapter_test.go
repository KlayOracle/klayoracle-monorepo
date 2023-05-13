package adapter

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/boot"
	"github.com/stretchr/testify/assert"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
)

func TestAdapterMarshalOk(t *testing.T) {

	wd, _ := os.Getwd()
	boot.Boot(path.Join(wd, ".."), path.Join(wd, "..", "config.yaml"), path.Join(wd, "..", ".env"))

	newAdapter := &protoadapter.Adapter{}
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
						"reducers": [{"function": "PARSE","args": ["$.RAW.KLAY.USD.PRICE"]},{"function": "FLOAT64_MUL_UINT64","args": ["1000000000"]}],
      					"payload": "{\"type\":\"limit\",\"side\":\"buy\",\"price\":1.058e-9,\"size\":100,\"currency\":[\"USDT_BTC\",\"WEMIX_KLAY\",\"KLAY_USDT\"]}"
						},
						{
						"url": "https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
						"request_type": 1,
						"headers": [{"field": {"X-CoinAPI-Key": "${X_COIN_API_KEY}"}}],
						"reducers": [{"function": "PARSE","args": ["$.rate"]},{"function": "FLOAT64_MUL_UINT64","args": ["1000000000"]}]
						}
					]
				}
				`)

	Import(byt, newAdapter, true)

	assert.True(t, newAdapter.Active)
	assert.Equal(t, newAdapter.Category, protoadapter.FeedCategory_PRICE_FEED)
	assert.Equal(t, newAdapter.Name, "KLAY_USD")
	assert.Equal(t, newAdapter.JobType, protoadapter.JobTypes_DATA_FEED)
	assert.Equal(t, newAdapter.AdapterId, "0x8b7460cccfa0aca303ee85c3fb81c344faad2fbab415adc32b2984008b7efd76")
	assert.Equal(t, newAdapter.OracleAddress, "0xCC4377b912c4517Fe895817c6a7c6937D92A70B3")
	assert.Equal(t, newAdapter.Frequency, int64(30000000000))

	importedAdapter := []*protoadapter.Feed{
		{
			Url:         "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=KLAY&tsyms=USD",
			RequestType: protoadapter.RequestType_GET,
			Headers: []*protoadapter.Feed_Header{
				{
					Field: map[string]string{
						"Content-Type": "application/json",
					},
				},
				{
					Field: map[string]string{
						"Authorization": "${BEARER_TOKEN}",
					},
				},
			},
			Reducers: []*protoadapter.Reducer{
				{
					Function: "PARSE",
					Args:     []string{"$.RAW.KLAY.USD.PRICE"},
				},
				{
					Function: "FLOAT64_MUL_UINT64",
					Args:     []string{"1000000000"},
				},
			},
			Payload: `{"type":"limit","side":"buy","price":1.058e-9,"size":100,"currency":["USDT_BTC","WEMIX_KLAY","KLAY_USDT"]}`,
		},
		{
			Url:         "https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
			RequestType: protoadapter.RequestType_POST,
			Headers: []*protoadapter.Feed_Header{
				{
					Field: map[string]string{
						"X-CoinAPI-Key": "${X_COIN_API_KEY}",
					},
				},
			},
			Reducers: []*protoadapter.Reducer{
				{
					Function: "PARSE",
					Args:     []string{"$.rate"},
				},
				{
					Function: "FLOAT64_MUL_UINT64",
					Args:     []string{"1000000000"},
				},
			},
		},
	}

	for i := 0; i < len(newAdapter.Feeds); i++ {
		feed := newAdapter.Feeds[i]
		importedFeed := importedAdapter[i]

		assert.Equal(t, feed.Url, importedFeed.Url)
		assert.Equal(t, feed.RequestType, importedFeed.RequestType)

		for i := 0; i < len(feed.Headers); i++ {
			header := feed.Headers[i].Field
			importedHeader := importedFeed.Headers[i].Field

			var headerKey, headerValue, importedHeaderKey, importedHeaderValue string

			for headerKey, headerValue = range header {
			}

			for importedHeaderKey, importedHeaderValue = range importedHeader {
			}

			assert.Equal(t, headerKey, importedHeaderKey)
			assert.Equal(t, headerValue, os.ExpandEnv(importedHeaderValue))
		}

		for i := 0; i < len(feed.Reducers); i++ {
			reducer := feed.Reducers[i]
			importedReducer := importedFeed.Reducers[i]

			assert.Equal(t, reducer.Function, importedReducer.Function)
			assert.EqualValues(t, reducer.Args, importedReducer.Args)
		}

		var payload1, payload2 interface{}

		json.Unmarshal([]byte(feed.Payload), &payload1)
		json.Unmarshal([]byte(importedFeed.Payload), &payload2)

		assert.Equal(t, payload1, payload2)
	}

}
