package sort

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// SelectSort sort by select
func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		idx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[idx] {
				idx = j
			}
		}
		swap(arr, i, idx)
	}
}

// BubbleSort sort by bubble
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

// InsertSort sort by insert
func InsertSort(arr []int) {

}

// MergeSort merge
func MergeSort(arr []int) {

}

// HeapSort heap
func HeapSort(arr []int) {

}

// QuickSort quick
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, left, right int) int {
	mid := left + (right-left)/2
	swap(arr, left, mid)

	// arr[left + 1, j] < v
	// arr[j + 1, i] >= v
	j := left
	for i := left + 1; i <= right; i++ {
		if arr[i] < arr[left] {
			j++
			swap(arr, i, j)
		}
	}
	swap(arr, left, j)

	return j
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	p := partition(arr, left, right)

	quickSort(arr, left, p-1)
	quickSort(arr, p+1, right)

}
