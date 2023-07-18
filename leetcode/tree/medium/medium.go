package medium

import "github.com/cover-pot/algorithm/leetcode/tree"

// ConstructMaximumBinaryTree 654. Maximum Binary Tree
func ConstructMaximumBinaryTree(nums []int) *tree.TreeNode {
	var constructTree func([]int, int, int) *tree.TreeNode

	constructTree = func(nums []int, left int, right int) *tree.TreeNode {
		if left > right {
			return nil
		}

		max := left

		for i := left; i <= right; i++ {
			if nums[i] > nums[max] {
				max = i
			}
		}

		node := &tree.TreeNode{Val: nums[max]}
		node.Left = constructTree(nums, left, max-1)
		node.Right = constructTree(nums, max+1, right)

		return node
	}

	return constructTree(nums, 0, len(nums)-1)
}
