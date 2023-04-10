package job

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"github.com/klayoracle/klayoracle-monorepo/node/reducer"
)

var reducers = map[string]interface{}{
	"PARSE":              reducer.Parse,
	"STRING_TO_FLOAT32":  reducer.StringToFloat32,
	"STRING_TO_FLOAT64":  reducer.StringToFloat64,
	"STRING_TO_UINT64":   reducer.StringToUint64,
	"FLOAT64_MUL_UINT64": reducer.Float64MulByUint,
	"UINT_DIV_UINT":      reducer.NumberDivByUint,
	"UINT_MUL_UINT":      reducer.NumberMulByUint,
}

// Aggregate Off-chain
// Sort prices in DESC order
// Get Median / Mean
// Store answers in DB
// Use TWAP (Time weighted average price) to get final answer
// Send answer Onchain
// Monitor deviation and update answer based on deviation

// Run should not halt execution of program if error occurred
// simply pass error back to go routine anonymous function to log
func Run(feed protonode.Adapter) error {
	var answers []*big.Int

	for _, feed := range feed.Feeds {
		response, err := request(feed)
		if err != nil {
			return err
		}

		reduce(feed, response)
		fmt.Println(response)
	}

	fmt.Println(answers)

	return nil
}

func request(f *protonode.Feed) (string, error) {
	method := http.MethodGet

	if f.RequestType == 1 {
		method = http.MethodPost
	}

	bodyReader := bytes.NewReader([]byte(f.Payload))

	req, err := http.NewRequest(method, f.Url, bodyReader)
	defer req.Body.Close()
	if err != nil {
		return "", err
	}

	var (
		headerKey, headerValue string
	)

	for _, header := range f.Headers {
		for headerKey, headerValue = range header.Field {
		}
		req.Header.Set(headerKey, headerValue)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)

	return string(body), nil
}

func reduce(f *protonode.Feed, jsonString string) {

	//for _, reducer := range collection {
	//
	//
	//}
}

//Setup postgres database:
//- store adapter request sent to node (node)
//- oracle response to adapter (dp)
//- answer history to allow TWAP calculation (node)
