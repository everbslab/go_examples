package math

// AverageFloats calculates AVG of slice of floats
func AverageFloats(args []float64) float64 {
	var sum = SumFloats(args)

	if len(args) == 0 {
		return 0
	}

	return sum / float64(len(args))
}

// SumFloats calculates SUM of slice of floats
func SumFloats(args []float64) float64 {
	var sum float64
	size := len(args)

	if size == 0 {
		return 0
	}

	for _, val := range args {
		sum += val
	}

	return sum
}
