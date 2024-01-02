package top_interview_150

/*
*
88. 合并两个有序数组
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	idx := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]

			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
		idx--
	}

	for i >= 0 {
		nums1[idx] = nums1[i]
		idx--
		i--
	}
	for j >= 0 {
		nums1[idx] = nums2[j]
		idx--
		j--
	}
}

/*
*
27. 移除元素
*/
func removeElement(nums []int, val int) int {
	r := len(nums) - 1

	for i := 0; i <= r; i++ {
		if nums[i] == val {
			nums[i], nums[r] = nums[r], nums[i]
			i--
			r--
		}
	}

	return r + 1
}
