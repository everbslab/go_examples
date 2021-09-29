package main

import (
	"fmt"
	"github.com/everbslab/go_examples/pkg/algo"
	"math"
	"math/rand"
	"time"
)

func main() {
	rndRangeSz := 1_000_000
	binarySearchExample(rndRangeSz)
}

func binarySearchExample(sz int) {
	fmt.Println("Binary Search algorithm example:")

	rand.Seed(time.Now().UTC().UnixNano())

	haystack := make([]int, sz)
	for i := 0; i <= len(haystack)-1; i++ {
		haystack[i] = i
	}

	logRes := math.Round(math.Log2(float64(sz)))
	fmt.Printf("Algo complexity: O^(log n). Log2(%d) = %.0f\n", sz, logRes)

	needle := haystack[rand.Intn(len(haystack))] // Pick random needle value

	if ok, index := algo.BinarySearch(haystack, needle); ok {
		fmt.Printf("Value `%d` was found by index `%d`.\n", needle, index)
	} else {
		fmt.Printf("Value `%d` was NOT found.\n", needle)
	}
}
