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

// CheckTree 2236. Root Equals Sum of Children
func CheckTree(root *tree.TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}

// RangeSumBST 938. Range Sum of BST
func RangeSumBST(root *tree.TreeNode, low int, high int) int {
	//var inorderSum func(*tree.TreeNode, int, int)
	//
	//sum := 0
	//inorderSum = func(node *tree.TreeNode, low int, high int) {
	//	if node == nil {
	//		return
	//	}
	//
	//	inorderSum(node.Left, low, high)
	//
	//	if node.Val >= low && node.Val <= high {
	//		sum += node.Val
	//	}
	//
	//	inorderSum(node.Right, low, high)
	//}
	//
	//inorderSum(root, low, high)
	//
	//return sum

	// optimal solution
	if root == nil {
		return 0
	}

	var sum = 0

	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	l := 0
	r := 0

	if low <= root.Val {
		l = RangeSumBST(root.Left, low, high)
	}

	if high >= root.Val {
		r = RangeSumBST(root.Right, low, high)
	}

	return sum + l + r
}

// TwoSum 1. Two Sum
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if k, ok := m[target-num]; ok {
			return []int{i, k}
		} else {
			m[num] = i
		}
	}

	return []int{-1, -1}
}
