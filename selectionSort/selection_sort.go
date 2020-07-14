package selectionSort

func SelectionSort(arr []int) []int {
	sizeArr := len(arr)

	for i := 0; i < sizeArr; i++ {
		minValueIndex := i

		for j := i; j < sizeArr; j++ {
			if arr[j] < arr[minValueIndex] {
				minValueIndex = j
			}
		}

		arr[i], arr[minValueIndex] = arr[minValueIndex], arr[i]
	}

	return arr
}
