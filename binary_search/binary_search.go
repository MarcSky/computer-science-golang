package binary_search

import (
	"errors"
)

var (
	notFound = errors.New("element not found")
)

// array m must be sorted
// result is index found element in array or error
func BinarySearch(m []int, item int) (int, error) {
	low := 0
	high := len(m) - 1

	for low <= high {
		mid := (high + low) / 2
		guess := m[mid]
		if guess == item {
			return mid, nil
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, notFound
}
