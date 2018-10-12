package algorithm

func SelectionSort(arr []int) []int {
	for key := range arr {
		min := key
		for i := key + 1; i < len(arr); i++ {
			if arr[i] < arr[min] {
				min = i
			}
		}
		arr[key], arr[min] = arr[min], arr[key]
	}
	return arr
}
