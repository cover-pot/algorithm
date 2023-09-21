package september

import "math"

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

// MinOperations 给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。
//
// 如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1
// 9.20 1658
func MinOperations(nums []int, x int) int {

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	l := len(nums)
	sum := 0

	for _, num := range nums {
		sum += num
	}

	if sum < x {
		return -1
	}

	right := 0
	lsum := 0
	rsum := sum
	ans := l + 1
	for left := -1; left < l; left++ {
		if left != -1 {
			lsum += nums[left]
		}

		for right < l && lsum+rsum > x {
			rsum -= nums[right]
			right++
		}
		if lsum+rsum == x {
			ans = min(ans, (left+1)+(l-right))
		}

	}
	if ans > l {
		return -1
	}
	return ans
}

// FindMedianSortedArrays 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
// 9.21  4
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		nums_im1 := math.MinInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		nums_i := math.MaxInt32
		if i != m {
			nums_i = nums1[i]
		}
		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}
		if nums_im1 <= nums_j {
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
