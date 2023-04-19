package job

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/klayoracle/klayoracle-monorepo/node/config"

	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"github.com/klayoracle/klayoracle-monorepo/node/reducer"
)

var reducers = map[string]func([]interface{}) (interface{}, error){
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

	answers, err := aggregate(feed)
	if err != nil {
		return err
	}

	fmt.Println(answers)

	return nil
}

// DryRun for dev purposes to test if everythign will work on production
func DryRun(feed protonode.Adapter) ([]*big.Int, error) {
	return aggregate(feed)
}

func aggregate(feed protonode.Adapter) ([]*big.Int, error) {
	var answers []*big.Int

	for _, feed := range feed.Feeds {
		response, err := request(feed)
		if err != nil {
			config.Loaded.Logger.Warnw("error processing feeds request", "url", feed.Url, "body", feed.Payload, "header", feed.Headers)
			return nil, err
		}

		output, err := reduce(feed, response)
		if err != nil {
			config.Loaded.Logger.Warnw("error processing while reducing", "error", err)
			return nil, err
		}

		answers = append(answers, output)
	}

	return answers, nil
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

func reduce(f *protonode.Feed, jsonString string) (*big.Int, error) {

	var (
		result    interface{}
		reducer   *protonode.Reducer
		parseFunc func([]interface{}) (interface{}, error)
		err       error
	)

	reducer = f.Reducers[0]
	parseFunc = reducers[reducer.Function]
	args := []interface{}{jsonString, reducer.Args[0]}

	config.Loaded.Logger.Infow("running", "reducer", "PARSE", "args", args)

	result, err = parseFunc(args)
	if err != nil {
		return nil, err
	}

	config.Loaded.Logger.Infow("result", "reducer", "PARSE", "output", result)

	otherReducers := f.Reducers[1:]

	for _, reducer = range otherReducers {
		var args []interface{}

		parseFunc = reducers[reducer.Function]
		args = append(args, result)

		for _, arg := range reducer.Args {
			args = append(args, interface{}(arg))
		}

		config.Loaded.Logger.Infow("running ", "reducer", reducer.Function, "args", args)

		result, err = parseFunc(args)
		if err != nil {
			return nil, err
		}

		config.Loaded.Logger.Infow("result", "reducer", reducer.Function, "output", result)

	}

	return result.(*big.Int), err
}
