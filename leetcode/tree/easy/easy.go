package easy

import "github.com/cover-pot/algorithm/leetcode/tree"

// PreorderTraversal
// 144. Binary Tree Preorder Traversal
func PreorderTraversal(root *tree.TreeNode) []int {
	res := make([]int, 0)
	var preorderHelp func(root *tree.TreeNode)
	preorderHelp = func(root *tree.TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorderHelp(root.Left)
		preorderHelp(root.Right)

	}

	preorderHelp(root)
	return res
}

// InorderTraversal
// 94. Binary Tree Inorder Traversal
func InorderTraversal(root *tree.TreeNode) []int {
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

// PostorderTraversal
// 145. Binary Tree Postorder Traversal
func PostorderTraversal(root *tree.TreeNode) []int {
	res := make([]int, 0)
	var postorderHelp func(root *tree.TreeNode)
	postorderHelp = func(root *tree.TreeNode) {
		if root == nil {
			return
		}
		postorderHelp(root.Left)
		postorderHelp(root.Right)
		res = append(res, root.Val)

	}

	postorderHelp(root)
	return res
}
