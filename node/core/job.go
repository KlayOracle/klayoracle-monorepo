package core

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/klayoracle/klayoracle-monorepo/node/storage"

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

const TimeWeight = 1 //Time duration to check for cumulative price

// Aggregate Off-chain
// Fetch answers in DB from last time.Now() - time.Hour * TimeWeight
// Use TWAP (Time weighted average price) to get round answer
// Store round answer
// Send round answer On-chain
// Monitor deviation and update answer based on deviation (Not done)

// Run should not halt execution of program if error occurred
// simply pass error back to go routine anonymous function to log
func Run(feed protonode.Adapter) (int64, error) {

	//Aggregate
	answers, err := aggregate(feed)
	if err != nil {
		return 0, err
	} else {

		t0 := time.Now().Add(-time.Hour * TimeWeight)

		rows, err := storage.FetchResponsesFromT0(&feed, t0)
		if err != nil {
			return 0, err
		} else {
			var priceHistory []int64

			for rows.Next() {
				var responses []int64

				if err := rows.Scan(&responses); err != nil {
					return 0, err
				}

				priceHistory = append(priceHistory, responses...)
			}

			var int64Answers []int64

			for _, answer := range answers {
				int64Answers = append(int64Answers, answer.Int64())
			}

			priceHistory = append(priceHistory, int64Answers...)

			//TWAP
			roundAnswer := getTWAP(priceHistory)

			//Store Round
			err := storage.StoreResponses(&feed, int64Answers, roundAnswer)
			if err != nil {
				return 0, err
			}

			rows, err := storage.FetchRoundSize(&feed)
			if err != nil {
				return 0, err
			}

			var roundSize int

			for rows.Next() {
				rows.Scan(&roundSize)
			}

			config.Loaded.Logger.Infow("round", "number", roundSize, "answer", roundAnswer)

			return roundAnswer, nil
		}
	}

	return 0, nil
}

func getTWAP(prices []int64) int64 {
	var sum int64

	for _, price := range prices {
		sum += price
	}

	len := len(prices)

	return sum / int64(len)
}

// DryRun for dev purposes to test if everything will work on production
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

	fmt.Println(string(body))

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
