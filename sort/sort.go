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

}
