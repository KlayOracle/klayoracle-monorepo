package reducer

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
	"reflect"
	"strconv"
)

func PARSE(data interface{}, path string) (interface{}, error) {
	var jsonData interface{}

	err := json.Unmarshal([]byte(data.(string)), &jsonData)
	if err != nil {
		return nil, err
	}

	value, err := jsonpath.Get(path, jsonData)
	if err != nil {
		return nil, err
	}

	if reflect.TypeOf(value).Kind() == 23 { //23 slice
		return value.([]interface{})[0], nil
	}

	return value, nil
}

// PARSE_DEPRECATED use PARSE instead
func PARSE_DEPRECATED(data interface{}, arg protonode.Reducer) (interface{}, error) {
	var (
		jsonData interface{}
		value    interface{}
	)

	err := json.Unmarshal([]byte(data.(string)), &jsonData)
	if err != nil {
		return nil, err
	}

	jsonType := fmt.Sprintf("%T", jsonData)

	if jsonType == "[]interface {}" {
		var result []interface{}

		result = jsonData.([]interface{})

		for i := 0; i < len(arg.Args); i++ {
			index, _ := strconv.ParseInt(arg.Args[i], 10, 64)

			if i == len(arg.Args)-1 {
				value = result[index]
			} else {
				result = result[index].([]interface{})
			}
		}
	}

	if jsonType == "map[string]interface {}" {
		var result map[string]interface{}

		result = jsonData.(map[string]interface{})

		for i := 0; i < len(arg.Args); i++ {

			if i == len(arg.Args)-1 {
				value = result[arg.Args[i]]
			} else {
				result = result[arg.Args[i]].(map[string]interface{})
			}
		}
	}

	return value, nil
}
