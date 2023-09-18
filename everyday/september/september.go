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
