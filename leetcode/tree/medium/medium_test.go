package medium

import (
	"fmt"
	"github.com/cover-pot/algorithm/leetcode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	root := &leetcode.TreeNode{Val: 3}
	root.Left = &leetcode.TreeNode{Val: 9}
	root.Right = &leetcode.TreeNode{Val: 20}
	root.Right.Left = &leetcode.TreeNode{Val: 15}
	root.Right.Right = &leetcode.TreeNode{Val: 7}
	res := levelOrder(root)
	fmt.Println(res)
}

func Test_sumNumbers(t *testing.T) {
	root := &leetcode.TreeNode{
		Val:   1,
		Left:  &leetcode.TreeNode{Val: 2},
		Right: &leetcode.TreeNode{Val: 3},
	}

	res := sumNumbers(root)
	assert.Equal(t, res, 25)
}

func Test_rightSideView(t *testing.T) {
	root := &leetcode.TreeNode{
		Val: 1,
		Left: &leetcode.TreeNode{
			Val:   2,
			Right: &leetcode.TreeNode{Val: 5},
		},
		Right: &leetcode.TreeNode{
			Val:   3,
			Right: &leetcode.TreeNode{Val: 4},
		},
	}

	res := rightSideView(root)
	fmt.Println(res)

	tmp := []int{1, 2, 3}

	tt := make([]int, len(tmp))
	copy(tt, tmp)
	fmt.Println(tmp)
	fmt.Println(tt)
}

func Test_lowestCommonAncestor(t *testing.T) {
	root := &leetcode.TreeNode{
		Val: 6,
		Left: &leetcode.TreeNode{
			Val:  2,
			Left: &leetcode.TreeNode{Val: 0},
			Right: &leetcode.TreeNode{
				Val:   4,
				Left:  &leetcode.TreeNode{Val: 3},
				Right: &leetcode.TreeNode{Val: 5},
			},
		},
		Right: &leetcode.TreeNode{
			Val:   8,
			Left:  &leetcode.TreeNode{Val: 7},
			Right: &leetcode.TreeNode{Val: 9},
		},
	}
	p := &leetcode.TreeNode{Val: 2}
	q := &leetcode.TreeNode{Val: 8}
	ancestor := lowestCommonAncestor(root, p, q)
	fmt.Println(ancestor.Val)
}
