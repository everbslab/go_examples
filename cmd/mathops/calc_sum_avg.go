package main

import (
	"fmt"
	"github.com/everbslab/go_examples/pkg/calc"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Insufficient arguments passed")
		os.Exit(1)
	}

	fs := make([]float64, len(args))

	for i, v := range args {
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			fs[i] = fv
		} else {
			fmt.Println("Warning: ", err)
		}
	}

	ra, rs := calc.AverageFloats(fs), calc.SumFloats(fs)

	fmt.Printf("Avg of %v = %.2f\n", args, ra)
	fmt.Printf("Sum of %v = %.2f\n", args, rs)
}
