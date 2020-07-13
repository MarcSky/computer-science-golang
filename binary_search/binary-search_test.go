package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type binarySearchData struct {
	value    []int
	item     int
	expected int
	err      error
}

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	params := []binarySearchData{
		{[]int{1, 5, 7, 8, 9, 11, 14, 17, 20}, 17, 7, nil},
		{[]int{1, 5, 7, 8, 9, 11, 14, 17, 20}, 22, 0, notFound},

	}

	for _, p := range params {
		res, err := BinarySearch(p.value, p.item)
		assert.Equal(t, p.expected, res)
		assert.Equal(t, p.err, err)
	}

}
