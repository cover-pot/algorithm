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
