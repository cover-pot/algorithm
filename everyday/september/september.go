package september

// KWeakestRows 给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。
//
// 请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。
//
// 如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。
//
// 军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
// 9.18
func KWeakestRows(mat [][]int, k int) []int {
	var res []int

	m := make(map[int][]int)
	for i := 0; i < len(mat); i++ {
		cnt := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				break
			}
			cnt++
		}

		m[cnt] = append(m[cnt], i)
	}

	for i := 0; i <= len(mat[0]); i++ {
		if v, ok := m[i]; ok {
			res = append(res, v...)
		}
		if len(res) >= k {
			break
		}
	}

	return res[:k]
}

// FindDuplicate 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
//
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
//
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
// 9.19   287
// solution：
// 第一次遇到x将其放到索引x处
// 第二次遇到x时 一定有索引x处的值和x相等 即arr[x] = x 此时直接返回
// step 1: if nums[i] == i 已经正确归位 继续往下走
// step 2: 如果当前位置的值：nums[i] 和索引为nums[i] 处的值相等： 即nums[i] == nums[nums[i]] 那么代表这个元素重复 直接返回 nums[i]
// step 3: 交换索引i处的值 和索引i的值为索引处的值 即 swap(nums[i], nums[nums[i]])
func FindDuplicate(nums []int) int {
	//m := make(map[int]struct{})
	//
	//for _, num := range nums {
	//	if _, ok := m[num]; ok {
	//		return num
	//	} else {
	//		m[num] = struct{}{}
	//	}
	//}
	//return -1

	idx := 0
	for idx < len(nums) {
		for nums[idx] == idx {
			idx++
			continue
		}
		if nums[nums[idx]] == nums[idx] {
			return nums[idx]
		}

		tmp := nums[idx]
		nums[idx] = nums[tmp]
		nums[tmp] = tmp
	}

	return -1
}
