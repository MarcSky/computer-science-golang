package quick_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var a []int

func init() {
	for i := 0; i < 100000; i++ {
		a = append(a, i)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		QuickSort(a, 0, len(a)-1)
	}
}

func TestQuickSort(t *testing.T) {
	t.Parallel()

	type quickSortData struct {
		value    []int
		expected []int
	}

	params := []quickSortData{
		{[]int{5, 2, 3, 1, 6, 8, 9, 5, 4, 3, 3}, []int{1, 2, 3, 3, 3, 4, 5, 5, 6, 8, 9}},
		{[]int{4, 9, 7, 6, 2, 3, 8}, []int{2, 3, 4, 6, 7, 8, 9}},
	}

	for _, p := range params {
		res := QuickSort(p.value, 0, len(p.value)-1)
		assert.Equal(t, p.expected, res)
	}
}
