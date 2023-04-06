package reducer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/PaesslerAG/jsonpath"
)

func Parse(data interface{}, path string) (interface{}, error) {
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

func StringToFloat64(value string) (float64, error) {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func StringToFloat32(value string) (float32, error) {
	result, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func StringToUint64(value string) (uint64, error) {
	result, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func Float64MulByUint(value float64, times uint64) (*big.Int, error) {
	newFloatVal := new(big.Float).SetFloat64(value)
	timesFloatVal := new(big.Float).SetUint64(times)

	newFloatVal.Abs(newFloatVal)
	newFloatVal.Mul(newFloatVal, timesFloatVal)

	result := new(big.Int)
	newFloatVal.Int(result)

	if result.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), fmt.Errorf("zero value returned")
	}

	return result, nil
}

func NumberDivByUint(value uint64, times uint64) (*big.Int, error) {
	newUnitVal := new(big.Int).SetUint64(value)
	divUnitVal := new(big.Int).SetUint64(times)

	result := new(big.Int).Div(newUnitVal, divUnitVal)

	if result.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), fmt.Errorf("zero value returned")
	}

	return result, nil
}

func NumberMulByUint(value uint64, times uint64) (*big.Int, error) {
	newUnitVal := new(big.Int).SetUint64(value)
	divUnitVal := new(big.Int).SetUint64(times)

	result := new(big.Int).Mul(newUnitVal, divUnitVal)

	if result.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), fmt.Errorf("zero value returned")
	}

	return result, nil
}
