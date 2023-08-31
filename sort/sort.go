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

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	partition := func(arr []int, left, right int) int {
		mid := left + (right-left)/2
		swap(arr, left, mid)
		l := left + 1
		for l < right {
			for l < right && arr[l] < arr[left] {
				l += 1
			}
			for l < right && arr[right] > arr[left] {
				right -= 1
			}
			if l < right {
				swap(arr, l, right)
				l += 1
				right -= 1
			}
		}
		swap(arr, left, right)
		return right
	}

	p := partition(arr, left, right)

	quickSort(arr, left, p-1)
	quickSort(arr, p+1, right)

}
