package algo

import "sort"

// BinarySearch implements Binary Search Algorithm.
// Time complexity: O^(log n)
// Returns -1 if value wasn't found and boolean false
// Return N if value was found and boolean true
func BinarySearch(haystack []int, needle int) (int, bool) {
	var lower int
	var counter int

	if !sort.SliceIsSorted(haystack, func(i, j int) bool {
		return haystack[i] < haystack[j]
	}) {
		sort.Ints(haystack)
	}

	higher := len(haystack) - 1

	// In case if slice is empty
	if higher < 1 {
		return -1, false
	}

	// lower bound check
	if needle == haystack[lower] {
		return lower, true
	}

	// higher bound check
	if needle == haystack[higher] {
		return higher, true
	}

	for ok := true; ok; ok = lower <= higher {
		counter += 1
		median := (lower + higher) / 2

		if needle == haystack[median] {
			return median, true
		}

		if needle > haystack[median] {
			lower = median + 1
		} else {
			higher = median - 1
		}
	}

	return -1, false
}
