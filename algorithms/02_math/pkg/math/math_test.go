package math

import "testing"

type tablepairs struct {
	Example     []float64
	ExpectedAvg float64
	ExpectedSum float64
}

var tests = []tablepairs{
	{[]float64{1, 2}, 1.5, 3},
	{[]float64{1, 2, 3}, 2, 5},
	{[]float64{0, 1}, 0.5, 1},
	{[]float64{-1, 1}, 0, 0},
	{[]float64{1}, 1, 1},
	{[]float64{0}, 0, 0},
	{[]float64{}, 0, 0},
}

func TestAverageFloats(t *testing.T) {
	for _, pair := range tests {
		result := AverageFloats(pair.Example)

		if result != pair.ExpectedAvg {
			t.Error(
				"For example", pair.Example,
				"expected AVG", pair.ExpectedAvg,
				"got", result,
			)
		}
	}
}

func TestSumFloats(t *testing.T) {
	for _, pair := range tests {
		result := AverageFloats(pair.Example)

		if result != pair.ExpectedAvg {
			t.Error(
				"For example", pair.Example,
				"expected SUM", pair.ExpectedSum,
				"got", result,
			)
		}
	}
}

func BenchmarkAverageFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			AverageFloats(pair.Example)
		}
	}
}

func BenchmarkSumFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			SumFloats(pair.Example)
		}
	}
}
