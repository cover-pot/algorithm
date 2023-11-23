package medium

import (
	"github.com/cover-pot/algorithm/leetcode"
)

// ConstructMaximumBinaryTree 654. Maximum Binary Tree
func ConstructMaximumBinaryTree(nums []int) *leetcode.TreeNode {
	var constructTree func([]int, int, int) *leetcode.TreeNode

	constructTree = func(nums []int, left int, right int) *leetcode.TreeNode {
		if left > right {
			return nil
		}

		max := left

		for i := left; i <= right; i++ {
			if nums[i] > nums[max] {
				max = i
			}
		}

		node := &leetcode.TreeNode{Val: nums[max]}
		node.Left = constructTree(nums, left, max-1)
		node.Right = constructTree(nums, max+1, right)

		return node
	}

	return constructTree(nums, 0, len(nums)-1)
}

// FindKthLargest 215. Kth Largest Element in an Array
func FindKthLargest(nums []int, k int) int {
	partation := func(arr []int, l, r int) int {
		mid := l + (r-l)/2
		arr[l], arr[mid] = arr[mid], arr[l]
		left, right := l+1, r

		for left <= right {
			for left <= right && arr[left] < arr[l] {
				left++
			}

			for left <= right && arr[right] > arr[l] {
				right--
			}

			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}

		arr[right], arr[l] = arr[l], arr[right]

		return right
	}

	p := partation(nums, 0, len(nums)-1)
	for {

		if p+1 < k {
			p = partation(nums, 0, p)
		} else if p+1 > k {
			p = partation(nums, p+1, len(nums)-1)
		} else {
			break
		}
	}

	return nums[p]
}

// 106. 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *leetcode.TreeNode {
	m := make(map[int]int, len(inorder))
	for i, val := range inorder {
		m[val] = i
	}
	var helpBuildTree func(inorder, postorder []int, il, ir, pl, pr int) *leetcode.TreeNode

	helpBuildTree = func(inorder, postorder []int, il, ir, pl, pr int) *leetcode.TreeNode {
		if il > ir {
			return nil
		}
		root := &leetcode.TreeNode{Val: postorder[pr]}
		r := m[root.Val]
		root.Left = helpBuildTree(inorder, postorder, il, r-1, pl, pr+r-ir-1)
		root.Right = helpBuildTree(inorder, postorder, r+1, ir, pr+r-ir+1, pr-1)
		return root
	}

	return helpBuildTree(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

// 2 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
func addTwoNumbers(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	dummy := &leetcode.ListNode{Val: 0}
	cur := dummy

	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry

		carry = sum / 10
		sum = sum % 10

		cur.Next = &leetcode.ListNode{Val: sum}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		cur.Next = &leetcode.ListNode{Val: carry}
	}
	return dummy.Next
}

// 46 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	var backtrack func(nums []int, tmp []int, used []bool)

	backtrack = func(nums []int, tmp []int, used []bool) {
		if len(tmp) == len(nums) {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			tmp = append(tmp, nums[i])
			used[i] = true
			backtrack(nums, tmp, used)

			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}

	backtrack(nums, tmp, used)
	return res
}

// 72 编辑距离
// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1:i] == word2[j-1:j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j] + 1 // 删除
				if dp[i][j-1]+1 < dp[i][j] {
					dp[i][j] = dp[i][j-1] + 1 // 插入
				}
				if dp[i-1][j-1]+1 < dp[i][j] {
					dp[i][j] = dp[i-1][j-1] + 1 // 替换
				}
			}
		}
	}

	return dp[m][n]
}

// 75 颜色分类
func sortColors(nums []int) {
	l, r, i := -1, len(nums), 0

	for i < r {
		if nums[i] == 0 {
			l++
			nums[l], nums[i] = nums[i], nums[l]
			i++
		} else if nums[i] == 2 {
			r--
			nums[r], nums[i] = nums[i], nums[r]
		} else {
			i++
		}
	}
}
