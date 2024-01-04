package top_interview_150

/*
*
80. 删除有序数组中的重复项 II
*/
func removeDuplicatesTwo(nums []int) int {
	var process func([]int, int) int

	process = func(nums []int, k int) int {
		n := 0

		for _, num := range nums {
			if n < k || nums[n-k] != num {
				nums[n] = num
				n++
			}
		}
		return n
	}

	return process(nums, 2)
}
