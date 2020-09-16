package math

func AverageFloats(args []float64) float64 {
	var sum = SumFloats(args)

	if len(args) == 0 {
		return 0
	}

	return sum / float64(len(args))
}

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
