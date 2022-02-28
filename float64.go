package utils

import "math"

type Float64Functions struct{}

func Float64() Float64Functions {
	return Float64Functions{}
}

func (f Float64Functions) GetInt64(fl float64, decimals int) int64 {
	multiplier := math.Pow(10, float64(decimals))
	return int64(math.Round(fl * multiplier))
}
