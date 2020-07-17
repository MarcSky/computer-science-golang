package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var a []int

func init() {
	for i := 0; i < 1000000; i++ {
		a = append(a, i)
	}
}

type binarySearchData struct {
	value    []int
	item     int
	expected int
	err      error
}

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	params := []binarySearchData{
		{a, 17, 17, nil},
		{a, 1000001, 0, notFound},
	}

	for _, p := range params {
		res, err := BinarySearch(p.value, p.item)
		assert.Equal(t, p.expected, res)
		assert.Equal(t, p.err, err)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = BinarySearch(a, 1000000)
	}
}

func stupidSearch(a []int, item int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == item {
			return i
		}
	}

	return -1
}

func BenchmarkStupidSearch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = stupidSearch(a, 1000000)
	}
}
