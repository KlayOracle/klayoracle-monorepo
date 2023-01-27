package adapter

import (
	"fmt"
	"github.com/klayoracle/klayoracle-monorepo/data-provider/proto"
	pb "google.golang.org/protobuf/proto"
	"testing"
)

const jsonAdapter = `
{
  "active": true,
  "name": "KLAY_USD",
  "job_type": "DATA_FEED",
  "adapter_id": "efbdab54419511edb8780242ac120002",
  "oracle": "0xCC4377b912c4517Fe895817c6a7c6937D92A70B3",
  "feeds": [
    {
      "url": "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=KLAY&tsyms=USD",
      "headers": [
        {"Content-Type" :  "application/json"}
      ],
      "request_type": "GET",
      "reducers": [
        {
          "function": "PARSE",
          "args": ["RAW","KLAY","USD","PRICE"]
        },
        {
          "function": "MUL",
          "args": [1000000000]
        }
      ]
    },
    {
      "url": "https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
      "headers": [
        {"X-CoinAPI-Key": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXX"}
      ],
      "reducers": [
        {
          "function": "PARSE",
          "args": ["rate"]
        },
        {
          "function": "MUL",
          "args": [1000000000]
        }
      ]
    }
  ]
}
`

func TestAdapterMarshalOk(t *testing.T) {
	adapter := &proto.Adapter{}

	pb.Unmarshal([]byte(jsonAdapter), adapter)

	//assert.True(t, adapter.Active)
	fmt.Printf("%+v%T\n", adapter, adapter.Active)
	//fmt.Println(adapter.Name)
}
