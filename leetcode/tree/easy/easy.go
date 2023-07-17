package easy

import "github.com/cover-pot/algorithm/leetcode/tree"

func inorderTraversal(root *tree.TreeNode) []int {
	res := make([]int, 0)
	var inorderHelp func(root *tree.TreeNode)
	inorderHelp = func(root *tree.TreeNode) {
		if root == nil {
			return
		}
		inorderHelp(root.Left)
		res = append(res, root.Val)
		inorderHelp(root.Right)

	}

	inorderHelp(root)
	return res
}
