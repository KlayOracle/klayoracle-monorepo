package reducer

import (
	"encoding/json"

	"github.com/klayoracle/klayoracle-monorepo/node/protonode"
)

func PARSE(data interface{}, arg protonode.Reducer) (interface{}, error) {
	var jsonData interface{}

	err := json.Unmarshal([]byte(data.(string)), &jsonData)
	if err != nil {
		return nil, err
	}

	var (
		result map[string]interface{}
		value  interface{}
	)

	result = jsonData.(map[string]interface{})

	for i := 0; i < len(arg.Args); i++ {

		if i == len(arg.Args)-1 {
			value = result[arg.Args[i]]
		} else {
			result = result[arg.Args[i]].(map[string]interface{})
		}
	}

	return value, nil
}
