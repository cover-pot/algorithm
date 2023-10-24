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
	var merge func(l, mid, r int, arr []int)
	var mergeSort func(l, r int, arr []int)

	merge = func(l, mid, r int, arr []int) {
		tmp := arr[l : r+1]

		i, j := l, mid+1

		// 每轮循环为arr[k]赋值
		for k := l; k <= r; k++ {
			// 左边越界
			if i > mid {
				arr[k] = tmp[j-l]
				j++
			} else if j > r { //右边越界
				arr[k] = tmp[i-l]
				i++
			} else if tmp[i-l] < tmp[j-l] { // 左边的值小
				arr[k] = tmp[i-l]
				i++
			} else {
				arr[k] = tmp[j-l]
				j++
			}
		}

	}

	mergeSort = func(l, r int, arr []int) {
		if l >= r {
			return
		}
		mid := l + (r-l)/2

		mergeSort(l, mid, arr)
		mergeSort(mid+1, r, arr)
		merge(l, mid, r, arr)

	}

	mergeSort(0, len(arr)-1, arr)
}

// HeapSort heap
// parent index = (i - 1） / 2
// left child = 2 * i + 1
// right child = 2 * i + 2
func HeapSort(arr []int) {

	if len(arr) < 2 {
		return
	}

	siftDown := func(k, n int, arr []int) {
		for 2*k+1 < n {
			j := 2*k + 1
			if j+1 < n && arr[j] > arr[j+1] {
				j++
			}

			if arr[k] > arr[j] {
				swap(arr, k, j)
				k = j
			} else {
				break
			}
		}
	}

	// 从第一个非叶子结点开始 复杂度O（n）
	// 最后一个元素为 len(arr) - 1  父亲节点就是 (len(arr) - 1 - 1) / 2
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		siftDown(i, len(arr), arr)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		siftDown(0, i, arr)
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
