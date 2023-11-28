package hot100

import (
	"fmt"
	"sort"
)

// 3、无重复最长子串
func lengthOfLongestSubstring(s string) int {

	window := make(map[string]int, 0)

	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right : right+1]

		right++
		window[c]++
		// 窗口有重复字符
		for window[c] > 1 {
			d := s[left : left+1]
			left++
			window[d]--
		}
		if res < (right - left) {
			res = right - left
		}
	}

	return res
}

// 19 删除链表倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	fast := head
	slow := dummy

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

// 22. 括号生成
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}
	track := ""

	var backtrack func(left, right int, track string)
	backtrack = func(left, right int, track string) {
		// 右括号少 不合法
		if right < left {
			return
		}
		// 小于0的括号不合法
		if left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			tmpByte := make([]byte, len(track))
			copy(tmpByte, track)
			fmt.Println(string(tmpByte))
			res = append(res, string(tmpByte))
			return
		}

		// 放一个左括号
		track = fmt.Sprintf("%s(", track)
		backtrack(left-1, right, track)
		track = track[:len(track)-1]

		// 放一个右括号
		track = fmt.Sprintf("%s)", track)
		backtrack(left, right-1, track)
		track = track[:len(track)-1]
	}

	backtrack(n, n, track)

	return res
}

// 11.盛水最多的容器
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	min := func(x, y int) int {
		if height[x] < height[y] {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}
	for i < j {
		minVal := min(i, j)
		res = max(res, height[minVal]*(j-i))
		if i == minVal {
			i++
		}
		if j == minVal {
			j--
		}
	}

	return res
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			break
		}
	}

	l, r = mid, mid

	for l > 0 && nums[l] == target {
		l--
	}

	for r < len(nums) && nums[r] == target {
		r++
	}

	return []int{l + 1, r - 1}
}

type Node struct {
	Id  int
	Pid int
}

// 39、组合总数
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := make([][]int, 0)
	var backtrack func([]int, int, int, []int)

	backtrack = func(cand []int, begin, target int, tmp []int) {
		if target < 0 {
			return
		}
		if target == 0 {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}

		for i := begin; i < len(cand); i++ {
			tmp = append(tmp, cand[i])

			backtrack(cand, i, target-candidates[i], tmp)

			tmp = tmp[:len(tmp)-1]

		}
	}
	backtrack(candidates, 0, target, []int{})

	return res
}
