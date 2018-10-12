package algorithm

func InsertionSort(arr []int) []int {
	for key := range arr {
		for i := key - 1; i >= 0; i-- {
			if arr[key] < arr[i] {
				arr[key], arr[i] = arr[i], arr[key]
				key--
			} else {
				break
			}
		}
	}
	return arr
}
