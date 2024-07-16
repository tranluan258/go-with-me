package arrays

func BinarySearch(arr []int, low, high, v int) bool {
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] == v {
			return true
		} else if arr[mid] < v {
			low = mid + 1
		} else if arr[mid] > v {
			high = mid - 1
		}
	}
	return false
}
