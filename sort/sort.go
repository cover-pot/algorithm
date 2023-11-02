package sort

type Sort struct {
}

// SelectSort sort by select
func (*Sort) SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		idx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[idx] {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
		//swap(arr, i, idx)
	}
}

// BubbleSort sort by bubble
func (*Sort) BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// InsertionSort sort by insert
func (*Sort) InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

// MergeSort merge
func (*Sort) MergeSort(arr []int) {
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
func (*Sort) HeapSort(arr []int) {

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
				arr[k], arr[j] = arr[j], arr[k]
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
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(0, i, arr)
	}
}

// QuickSort quick sort three ways
func (*Sort) QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	lt, gt := partition(arr, left, right)

	quickSort(arr, left, lt)
	quickSort(arr, gt, right)

}

// arr[l+1, lt] < v
// arr[gt, r] > v
func partition(arr []int, left, right int) (int, int) {

	mid := left + (right-left)/2
	arr[mid], arr[left] = arr[left], arr[mid]

	lt, i, gt := left, left+1, right+1
	for i < gt {
		// arr[i] < arr[left]
		if arr[i] < arr[left] {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > arr[left] {
			// 扩充右边界限
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			// 相等
			i++
		}
	}
	arr[lt], arr[left] = arr[left], arr[lt]
	return lt - 1, gt
}
