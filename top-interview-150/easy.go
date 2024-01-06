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

/*
*26. 删除有序数组中的重复项
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, idx := 0, 1

	for idx < len(nums) {
		// 不相等交换
		if nums[l] != nums[idx] {
			nums[l+1] = nums[idx]
			l++
		}
		// 右侧指针移动
		idx++
	}

	return l + 1
}

/*
*169. 多数元素 摩尔投票
 */
func majorityElement(nums []int) int {
	candidateNum, candidate := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidateNum {
			candidate++
		} else {
			candidate--
		}

		if candidate == 0 {
			candidateNum = nums[i]
			candidate = 1
		}
	}

	return candidateNum
}
