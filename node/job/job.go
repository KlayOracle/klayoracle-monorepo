package job

import "github.com/klayoracle/klayoracle-monorepo/node/reducer"

var reducers = map[string]interface{}{
	"PARSE":              reducer.Parse,
	"STRING_TO_FLOAT32":  reducer.StringToFloat32,
	"STRING_TO_FLOAT64":  reducer.StringToFloat64,
	"STRING_TO_UINT64":   reducer.StringToUint64,
	"FLOAT64_MUL_UINT64": reducer.Float64MulByUint,
	"UINT_DIV_UINT":      reducer.NumberDivByUint,
}
