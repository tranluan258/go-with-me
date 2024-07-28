package arrays

func qs(arr []int, low int, high int) {
	if low >= high {
		return
	}

	pivotIdx := partition(arr, low, high)
	qs(arr, 0, pivotIdx-1)
	qs(arr, pivotIdx+1, high)
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]

	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			idx++
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}

	idx++
	arr[idx], arr[high] = arr[high], arr[idx]

	return idx
}

func QuickSort(arr []int) []int {
	qs(arr, 0, len(arr)-1)
	return arr
}
