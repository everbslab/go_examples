package main

import (
	"fmt"
	emath "github.com/everbslab/go_examples/algorithms/internals/math"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	fs := make([]float64, len(args))

	for i, v := range args {
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			fs[i] = fv
		} else {
			fmt.Println("Warning: ", err)
		}
	}

	ra := emath.AverageFloats(fs)
	rs := emath.SumFloats(fs)

	fmt.Printf("Avg of %v = %.2f\n", args, ra)
	fmt.Printf("Sum of %v = %.2f\n", args, rs)
}
