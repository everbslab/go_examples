package algo

// BinarySearch implements Binary Search Algorithm.
// Time complexity: O^(log n)
func BinarySearch(haystack []int, needle int) (bool, int) {
	var lower int
	var counter int

	higher := len(haystack) - 1

	// In case if slice is empty
	if higher < 1 {
		return false, 0
	}

	// lower bound check
	if needle == haystack[lower] {
		return true, lower
	}

	// higher bound check
	if needle == haystack[higher] {
		return true, higher
	}

	for ok := true; ok; ok = lower <= higher {
		counter += 1
		median := (lower + higher) / 2

		if needle == haystack[median] {
			return true, median
		}

		if needle > haystack[median] {
			lower = median + 1
		} else {
			higher = median - 1
		}
	}

	return false, 0
}
