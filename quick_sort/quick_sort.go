package quick_sort

func QuickSort(arr []int, l, r int) []int {
	if l >= r {
		return arr
	}

	lHold := l
	rHold := r
	pivot := arr[(l+r)/2] // O(n * log(n)) or pivot = arr[l] its O(n^2)

	for ; l < r; {
		for ; arr[l] < pivot; l++ {
		}

		for ; arr[r] > pivot; r-- {
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	QuickSort(arr, lHold, r)
	QuickSort(arr, l, rHold)

	return arr
}
