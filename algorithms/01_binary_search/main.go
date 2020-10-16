package main

import (
	"fmt"
	"github.com/everbslab/go_examples/algorithms/search"
	"math"
	"math/rand"
	"time"
)

func main() {
	binarySearchExample()
}

func binarySearchExample() {
	fmt.Println("Binary Search algorithm example:")

	sn := 1000000
	haystack := make([]int, sn)
	for i := 0; i <= len(haystack)-1; i++ {
		haystack[i] = i
	}

	logRes := math.Round(math.Log2(float64(sn)))
	fmt.Printf("Algo complexity: O^(log n). Log2(%d) = %.0f\n", sn, logRes)

	rand.Seed(time.Now().Unix())

	needle := haystack[rand.Intn(len(haystack))] // Pick random needle value

	if ok, index := search.BinarySearch(haystack, needle); ok {
		fmt.Printf("Value `%d` was found by index `%d`.\n", needle, index)
	} else {
		fmt.Printf("Value `%d` was NOT found.\n", needle)
	}
}
