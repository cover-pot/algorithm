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
