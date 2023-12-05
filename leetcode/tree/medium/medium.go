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

// 102 二叉树层序遍历
func levelOrder(root *leetcode.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*leetcode.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		ql := len(queue)
		tmp := make([]int, 0, ql)
		for i := 0; i < ql; i++ {
			cur := queue[0]
			tmp = append(tmp, cur.Val)
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}

// LCR 054. 把二叉搜索树转换为累加树
// 给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和
func convertBST(root *leetcode.TreeNode) *leetcode.TreeNode {
	var convertBSTHelp func(root *leetcode.TreeNode)

	res := 0

	// 反中序遍历
	convertBSTHelp = func(root *leetcode.TreeNode) {
		if root == nil {
			return
		}
		convertBSTHelp(root.Right)
		res += root.Val
		root.Val = res
		convertBSTHelp(root.Left)
	}
	convertBSTHelp(root)
	return root
}

// 1302. 层数最深叶子节点的和
func deepestLeavesSum(root *leetcode.TreeNode) int {
	// 记录每层和
	m := make(map[int]int)
	// 记录最大的层数
	max := 0
	var dfs func(root *leetcode.TreeNode, depth int)

	maxFunc := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dfs = func(root *leetcode.TreeNode, depth int) {
		if root == nil {
			return
		}
		max = maxFunc(max, depth)

		if _, ok := m[depth]; !ok {
			m[depth] = root.Val
		} else {
			m[depth] += root.Val
		}

		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)

	}
	dfs(root, 0)

	return m[max]
}

// 1305. 两棵二叉搜索树中的所有元素
// 二叉树中序遍历+merge数组
func getAllElements(root1 *leetcode.TreeNode, root2 *leetcode.TreeNode) []int {

	inorder := func(root *leetcode.TreeNode) []int {
		l := make([]int, 0)

		var dfs func(root *leetcode.TreeNode)

		dfs = func(root *leetcode.TreeNode) {
			if root == nil {
				return
			}
			dfs(root.Left)
			l = append(l, root.Val)
			dfs(root.Right)

		}
		dfs(root)
		return l

	}
	l1 := inorder(root1)
	l2 := inorder(root2)

	res := make([]int, 0, len(l1)+len(l2))

	i, j := 0, 0

	for i < len(l1) || j < len(l2) {
		if i < len(l1) && j >= len(l2) {
			res = append(res, l1[i])
			i++
		} else if i >= len(l1) && j < len(l2) {
			res = append(res, l2[j])
			j++
		} else {
			if l1[i] < l2[j] {
				res = append(res, l1[i])
				i++
			} else {
				res = append(res, l2[j])
				j++
			}
		}
	}

	return res
}

// 105 从前序和中序遍历构造二叉树
func buildBSTTree(preorder []int, inorder []int) *leetcode.TreeNode {
	m := make(map[int]int, len(inorder))

	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}

	var buildPreTree func(preorder []int, pl, pr int, inorder []int, il, ir int) *leetcode.TreeNode

	buildPreTree = func(preorder []int, pl, pr int, inorder []int, il, ir int) *leetcode.TreeNode {
		if pl > pr {
			return nil
		}
		rootVal := preorder[pl]
		root := &leetcode.TreeNode{Val: rootVal}

		idx := m[rootVal]

		root.Left = buildPreTree(preorder, pl+1, idx+pl-il, inorder, il, idx-1)
		root.Right = buildPreTree(preorder, idx+pl-il+1, pr, inorder, idx+1, ir)

		return root
	}

	return buildPreTree(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

// 107. 二叉树的层序遍历 II
func levelOrderBottom(root *leetcode.TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*leetcode.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		ql := len(queue)

		tmp := make([]int, 0, ql)
		for i := 0; i < ql; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, tmp)
	}

	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

// 109 将有序链表转换为一个BST
func sortedListToBST(head *leetcode.ListNode) *leetcode.TreeNode {
	if head == nil {
		return nil
	}
	length := 0
	root := head
	for head != nil {
		length++
		head = head.Next
	}

	var buildBST func(start, end int) *leetcode.TreeNode

	buildBST = func(start, end int) *leetcode.TreeNode {
		if start > end {
			return nil
		}
		mid := start + (end-start)/2
		left := buildBST(start, mid-1)
		r := &leetcode.TreeNode{Val: root.Val}
		root = root.Next
		r.Left = left
		r.Right = buildBST(mid+1, end)
		return r
	}

	return buildBST(0, length-1)

}

// 114 二叉树展开为链表
func flatten(root *leetcode.TreeNode) {

	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}

// 129 求根节点到叶节点的数字只喝
func sumNumbers(root *leetcode.TreeNode) int {

	var sumTree func(node *leetcode.TreeNode, preSum int) int

	sumTree = func(node *leetcode.TreeNode, preSum int) int {
		if node == nil {
			return 0
		}
		sum := preSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}

		return sumTree(node.Left, sum) + sumTree(node.Right, sum)
	}

	return sumTree(root, 0)
}

// 199 二叉树的右视图
func rightSideView(root *leetcode.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*leetcode.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		ql := len(queue)

		for i := 0; i < ql; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i == ql-1 {
				res = append(res, cur.Val)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return res
}

// 113 路径总和 II
func pathSum(root *leetcode.TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)

	var dfs func(root *leetcode.TreeNode, targetSum int, tmp []int)

	dfs = func(root *leetcode.TreeNode, targetSum int, tmp []int) {
		if root == nil {
			return
		}
		tmp = append(tmp, root.Val)
		targetSum -= root.Val

		if targetSum == 0 && root.Left == nil && root.Right == nil {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}

		dfs(root.Left, targetSum, tmp)
		dfs(root.Right, targetSum, tmp)

		tmp = tmp[:len(tmp)-1]
		targetSum += root.Val
	}

	dfs(root, targetSum, []int{})

	return res
}

// Node 116. 填充每个节点的下一个右侧节点指针
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	cur := root
	if cur == nil {
		return cur
	}

	queue := make([]*Node, 0)
	queue = append(queue, cur)

	for len(queue) != 0 {
		ql := len(queue)
		var pre *Node
		for i := 0; i < ql; i++ {
			cur = queue[0]
			queue = queue[1:]
			if i == 0 {
				pre = cur
			} else {
				pre.Next = cur
				pre = cur
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return root
}

// 230 二叉树中第k小的元素
func kthSmallest(root *leetcode.TreeNode, k int) int {
	res := 0

	var unInorder func(root *leetcode.TreeNode)

	unInorder = func(root *leetcode.TreeNode) {
		if root == nil {
			return
		}

		unInorder(root.Left)
		if k > 0 {
			res = root.Val
			k--
		}
		unInorder(root.Right)

	}

	unInorder(root)

	return res
}

// 235 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	for ((root.Val - p.Val) * (root.Val - q.Val)) > 0 {
		if q.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// 235 二叉树的最近公共祖先
func lowestCommonBSTAncestor(root, p, q *leetcode.TreeNode) *leetcode.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonBSTAncestor(root.Left, p, q)
	right := lowestCommonBSTAncestor(root.Right, p, q)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}

// 450 删除二叉搜索树的节点
func minimum(node *leetcode.TreeNode) *leetcode.TreeNode {
	if node.Left == nil {
		return node
	}
	return minimum(node.Left)
}

func removeMin(node *leetcode.TreeNode) *leetcode.TreeNode {

	if node.Left == nil {
		rightNode := node.Right
		node.Right = nil
		return rightNode
	}
	node.Left = removeMin(node.Left)
	return node

}
func deleteNode(root *leetcode.TreeNode, key int) *leetcode.TreeNode {
	if root == nil {
		return root
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		// 找到对应节点
		if root.Left == nil {
			rightNode := root.Right
			root.Right = nil
			return rightNode
		} else if root.Right == nil {
			leftNode := root.Left
			root.Left = nil
			return leftNode
		} else {
			// 后继节点 即右子树的最小节点
			successor := minimum(root.Right)
			successor.Right = removeMin(root.Right)
			successor.Left = root.Left

			root.Left, root.Right = nil, nil
			return successor
		}
	}
}

// 513. 找树左下角的值
func findBottomLeftValue(root *leetcode.TreeNode) int {
	var res int
	queue := make([]*leetcode.TreeNode, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)

	for len(queue) != 0 {
		ql := len(queue)
		for i := 0; i < ql; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i == 0 {
				res = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return res
}

// 515. 在每个树行中找最大值
func largestValues(root *leetcode.TreeNode) []int {
	res := make([]int, 0)
	queue := make([]*leetcode.TreeNode, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)

	for len(queue) != 0 {
		ql := len(queue)
		var maxVal int
		for i := 0; i < ql; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i == 0 {
				maxVal = cur.Val
			} else if maxVal < cur.Val {
				maxVal = cur.Val
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, maxVal)
	}

	return res
}
