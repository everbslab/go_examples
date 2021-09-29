package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type tablepairs struct {
	input     []float64
	expectAvg float64
	expectSum float64
}

var tests = []tablepairs{
	{[]float64{1, 2}, 1.5, 3},
	{[]float64{1, 2, 3}, 2, 6},
	{[]float64{0, 1}, 0.5, 1},
	{[]float64{-1, 1}, 0, 0},
	{[]float64{1}, 1, 1},
	{[]float64{0}, 0, 0},
	{[]float64{}, 0, 0},
}

func TestAverageFloats(t *testing.T) {
	for _, tp := range tests {
		fr := AverageFloats(tp.input)

		assert.Equal(t, tp.expectAvg, fr)
	}
}

func TestSumFloats(t *testing.T) {
	for _, tp := range tests {
		assert.Equal(t, tp.expectSum, SumFloats(tp.input))
	}
}

func BenchmarkAverageFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			AverageFloats(pair.input)
		}
	}
}

func BenchmarkSumFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			SumFloats(pair.input)
		}
	}
}
