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
func buildTree(inorder []int, postorder []int) *tree.TreeNode {
	m := make(map[int]int, len(inorder))
	for i, val := range inorder {
		m[val] = i
	}
	var helpBuildTree func(inorder, postorder []int, il, ir, pl, pr int) *tree.TreeNode

	helpBuildTree = func(inorder, postorder []int, il, ir, pl, pr int) *tree.TreeNode {
		if il > ir {
			return nil
		}
		root := &tree.TreeNode{Val: postorder[pr]}
		r := m[root.Val]
		root.Left = helpBuildTree(inorder, postorder, il, r-1, pl, pr+r-ir-1)
		root.Right = helpBuildTree(inorder, postorder, r+1, ir, pr+r-ir+1, pr-1)
		return root
	}

	return helpBuildTree(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}
