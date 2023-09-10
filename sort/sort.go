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

// InsertionSort sort by insert
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}

// MergeSort merge
func MergeSort(arr []int) {

}

// HeapSort heap
// parent index = (i - 1 / 2)
// left child = 2 * i + 1
// right child = 2 * i + 2
func HeapSort(arr []int) {

	// build heap
	heapInsert := func(arr []int, i int) {
		for i > 0 {
			if arr[i] > arr[(i-1)/2] {
				swap(arr, i, (i-1)/2)
				i = (i - 1) / 2
			} else {
				break
			}
		}
	}

	heapify := func(arr []int, i, k int) {
		for i < k && 2*i+1 < k {
			// max(left child, right child)
			maxIndex := 2*i + 1
			if 2*i+2 < k && arr[2*i+1] < arr[2*i+2] {
				maxIndex++
			}
			if arr[i] < arr[maxIndex] {
				swap(arr, i, maxIndex)
				i = maxIndex
			} else {
				break
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}

	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		heapify(arr, 0, i)
	}
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
