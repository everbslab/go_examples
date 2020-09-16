package search

import (
	"math/rand"
	"testing"
	"time"
)

type testpair struct {
	haystack []int
	needle   int
	Expected bool
}

var tests = []testpair{
	{[]int{}, 0, false},    // empty haystack validation
	{[]int{1, 2}, 1, true}, // lower check
	{[]int{1, 2}, 2, true},
	{[]int{1, 2}, 0, false},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 3, true},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 4, true},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 5, true},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 6, true}, // higher check
	{[]int{0, 1, 2, 3, 4, 5, 6}, 7, false},
	{[]int{0, 1, 2, 3, 4, 5, 6}, -1, false},
}

func TestBinarySearch(t *testing.T) {
	for _, pair := range tests {
		ok, _ := BinarySearch(pair.haystack, pair.needle)

		if ok != pair.Expected {
			t.Error(
				"For haystack", pair.haystack, "and needle", pair.needle,
				"expected", pair.Expected,
				"got", ok,
			)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			BinarySearch(pair.haystack, pair.needle)
		}
	}
}

func BenchmarkBinarySearch1Mln(b *testing.B) {
	sn := 1000000
	haystack := make([]int, sn)
	for i := 0; i <= len(haystack)-1; i++ {
		haystack[i] = i
	}

	rand.Seed(time.Now().Unix())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		needle := haystack[rand.Intn(len(haystack))] // Pick random needle value
		BinarySearch(haystack, needle)
	}
}

func BenchmarkBinarySearch100mln(b *testing.B) {
	sn := 100000000
	haystack := make([]int, sn)
	for i := 0; i <= len(haystack)-1; i++ {
		haystack[i] = i
	}

	rand.Seed(time.Now().Unix())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		needle := haystack[rand.Intn(len(haystack))] // Pick random needle value
		BinarySearch(haystack, needle)
	}
}
