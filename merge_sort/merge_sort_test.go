package merge_sort

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

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(a)
	}
}

func BenchmarkMergeMultiSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortMulti(a)
	}
}

func TestMergeSort(t *testing.T) {
	t.Parallel()

	type mergeSortData struct {
		value    []int
		expected []int
	}

	params := []mergeSortData{
		{[]int{5, 2, 3, 1, 6, 8, 9, 5, 4, 3, 3}, []int{1, 2, 3, 3, 3, 4, 5, 5, 6, 8, 9}},
	}

	for _, p := range params {
		res := MergeSort(p.value)
		assert.Equal(t, p.expected, res)
	}

	for _, p := range params {
		res := MergeSortMulti(p.value)
		assert.Equal(t, p.expected, res)
	}

}
