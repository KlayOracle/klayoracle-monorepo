package reducer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/PaesslerAG/jsonpath"
)

func Parse(args []interface{}) (interface{}, error) {
	data := args[0].(interface{})
	path := args[1].(string)

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

func StringToFloat64(args []interface{}) (interface{}, error) {
	value := args[0].(string)

	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func StringToFloat32(args []interface{}) (interface{}, error) { //float32
	value := args[0].(string)

	result, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func StringToUint64(args []interface{}) (interface{}, error) { //float64
	value := args[0].(string)

	result, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

//Note that the reason the interface values are float64 is because that's the default setting of
//the encoding/json decoder when unmarshaling JSON numbers into interface{} values.
//https://stackoverflow.com/a/70706082/2827325

func Float64MulByUint(args []interface{}) (interface{}, error) { //*big.Int
	var (
		times, value interface{}
		err          error
	)

	if reflect.TypeOf(args[0]).Kind() == 24 { //24 string
		value, err = strconv.ParseFloat(args[0].(string), 64)

		if err != nil {
			return nil, err
		}
	}

	if reflect.TypeOf(args[0]).Kind() == 14 {
		value = args[0].(float64)
	}

	if reflect.TypeOf(args[1]).Kind() == 24 { //24 string
		times, err = strconv.ParseUint(args[1].(string), 10, 64)

		if err != nil {
			return nil, err
		}
	}

	if reflect.TypeOf(args[1]).Kind() == 14 { //14 float64
		times = uint64(args[1].(float64))
	}

	newFloatVal := new(big.Float).SetFloat64(value.(float64))
	timesFloatVal := new(big.Float).SetUint64(times.(uint64))

	newFloatVal.Abs(newFloatVal)
	newFloatVal.Mul(newFloatVal, timesFloatVal)

	result := new(big.Int)
	newFloatVal.Int(result)

	if result.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), fmt.Errorf("zero value returned")
	}

	return result, nil
}

func NumberDivByUint(args []interface{}) (interface{}, error) { //*big.Int
	value := uint64(args[0].(float64))
	var (
		times interface{}
		err   error
	)

	if reflect.TypeOf(args[1]).Kind() == 24 { //24 string
		times, err = strconv.ParseUint(args[1].(string), 10, 64)

		return nil, err
	}

	if reflect.TypeOf(args[1]).Kind() == 14 { //14 float64
		times = uint64(args[1].(float64))
	}

	newUnitVal := new(big.Int).SetUint64(value)
	divUnitVal := new(big.Int).SetUint64(times.(uint64))

	result := new(big.Int).Div(newUnitVal, divUnitVal)

	if result.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), fmt.Errorf("zero value returned")
	}

	return result, nil
}

func NumberMulByUint(args []interface{}) (interface{}, error) { //*big.Int
	value := uint64(args[0].(float64))
	var (
		times interface{}
		err   error
	)

	if reflect.TypeOf(args[1]).Kind() == 24 { //23 string
		times, err = strconv.ParseUint(args[1].(string), 10, 64)

		return nil, err
	}

	if reflect.TypeOf(args[1]).Kind() == 14 { //23 float64
		times = uint64(args[1].(float64))
	}

	newUnitVal := new(big.Int).SetUint64(value)
	divUnitVal := new(big.Int).SetUint64(times.(uint64))

	result := new(big.Int).Mul(newUnitVal, divUnitVal)

	if result.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), fmt.Errorf("zero value returned")
	}

	return result, nil
}
