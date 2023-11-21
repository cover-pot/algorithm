package hoot100

import "fmt"

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
