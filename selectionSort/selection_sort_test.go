package selectionSort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var a []int

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		a = append(a, rand.Intn(999)-rand.Intn(999))
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelectionSort(a)
	}
}

type selectionSortData struct {
	value    []int
	expected []int
}

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	params := []selectionSortData{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}

	for _, p := range params {
		res := SelectionSort(p.value)
		assert.Equal(t, p.expected, res)
	}
}
