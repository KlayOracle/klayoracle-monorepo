package adapter

import (
	"testing"

	"github.com/klayoracle/klayoracle-monorepo/data-provider/protoadapter"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestAdapterMarshalOk(t *testing.T) {

	newAdapter := &protoadapter.Adapter{}
	byt := []byte(`
				{
				"active": true,
				"name": "KLAY_USD",
				"job_type": "DATA_FEED",
				"adapter_id":"efbdab54419511edb8780242ac120002",
				"oracle_address": "0xCC4377b912c4517Fe895817c6a7c6937D92A70B3",
				"category": 2,
				"feeds": [
						{
						"url": "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=KLAY&tsyms=USD",
						"request_type": 0,
						"headers": [{"field": {"Content-Type" : "application/json"}},{"field": {"Authorization" : "Bearer eyJraWQiOiJoeDhjMU5tRHJmV0U5dzFCRjNr..."}}],
						"reducers": [{"function": "PARSE","args": ["RAW","KLAY","USD","PRICE"]},{"function": "MUL","args": ["1000000000"]}]
						},
						{
						"url": "https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
						"request_type": 1,
						"headers": [{"field": {"X-CoinAPI-Key": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXX"}}],
						"reducers": [{"function": "PARSE","args": ["rate"]},{"function": "MUL","args": ["1000000000"]}]
						}
					]
				}
				`)
	err := protojson.Unmarshal(byt, newAdapter)
	if err != nil {
		panic(err)
	}

	assert.True(t, newAdapter.Active)
	assert.Equal(t, newAdapter.Category, protoadapter.FeedCategory_PRICE_FEED)
	assert.Equal(t, newAdapter.Name, "KLAY_USD")
	assert.Equal(t, newAdapter.JobType, protoadapter.JobTypes_DATA_FEED)
	assert.Equal(t, newAdapter.AdapterId, "efbdab54419511edb8780242ac120002")
	assert.Equal(t, newAdapter.OracleAddress, "0xCC4377b912c4517Fe895817c6a7c6937D92A70B3")

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
						"Authorization": "Bearer eyJraWQiOiJoeDhjMU5tRHJmV0U5dzFCRjNr...",
					},
				},
			},
			Reducers: []*protoadapter.Reducer{
				{
					Function: "PARSE",
					Args:     []string{"RAW", "KLAY", "USD", "PRICE"},
				},
				{
					Function: "MUL",
					Args:     []string{"1000000000"},
				},
			},
		},
		{
			Url:         "https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
			RequestType: protoadapter.RequestType_POST,
			Headers: []*protoadapter.Feed_Header{
				{
					Field: map[string]string{
						"X-CoinAPI-Key": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXX",
					},
				},
			},
			Reducers: []*protoadapter.Reducer{
				{
					Function: "PARSE",
					Args:     []string{"rate"},
				},
				{
					Function: "MUL",
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
			assert.Equal(t, headerValue, importedHeaderValue)
		}

		for i := 0; i < len(feed.Reducers); i++ {
			reducer := feed.Reducers[i]
			importedReducer := importedFeed.Reducers[i]

			assert.Equal(t, reducer.Function, importedReducer.Function)
			assert.EqualValues(t, reducer.Args, importedReducer.Args)
		}
	}

}
