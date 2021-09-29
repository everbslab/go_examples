package calc

// AverageFloats calculates AVG of slice of floats
func AverageFloats(args []float64) float64 {
	if len(args) == 0 {
		return 0
	}

	var sum = SumFloats(args)

	return sum / float64(len(args))
}

// SumFloats calculates SUM of slice of floats
func SumFloats(args []float64) float64 {
	var sum float64
	sz := len(args)

	if sz == 0 {
		return 0
	}

	for _, val := range args {
		sum += val
	}

	return sum
}
