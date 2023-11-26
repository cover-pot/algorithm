package senior_datastruct

import (
	"errors"
	"fmt"
)

// TreeNode BST的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BinarySearchTree BST
type BinarySearchTree struct {
	Root *TreeNode
	Size int
}

// Add 添加元素
func (b *BinarySearchTree) Add(val int) {
	b.Root = b.add(b.Root, val)
	return
}

// 返回插入节点的二叉树的根
func (b *BinarySearchTree) add(node *TreeNode, val int) *TreeNode {
	if node == nil {
		b.Size++
		return &TreeNode{Val: val}
	}
	if node.Val > val {
		node.Left = b.add(node.Left, val)
	} else if node.Val < val {
		node.Right = b.add(node.Right, val)
	}
	return node
}

// Contains 查询元素
func (b *BinarySearchTree) Contains(val int) bool {
	return b.contains(b.Root, val)
}

func (b *BinarySearchTree) contains(node *TreeNode, val int) bool {
	if node == nil {
		return false
	}

	if node.Val == val {
		return true
	} else if node.Val < val {
		return b.contains(node.Right, val)
	} else {
		return b.contains(node.Left, val)
	}
}

// PreOrder 前序遍历：根节点 -> 左子树 -> 右子树
func (b *BinarySearchTree) PreOrder() {
	b.preOrder(b.Root)
}

func (b *BinarySearchTree) preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d\t", node.Val)
	b.preOrder(node.Left)
	b.preOrder(node.Right)
}

// InOrder 中序遍历： 左子树 -> 根节点-> 右子树。 刚好是排序后的结果
func (b *BinarySearchTree) InOrder() {
	b.inOrder(b.Root)
}

func (b *BinarySearchTree) inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	b.preOrder(node.Left)
	fmt.Printf("%d\t", node.Val)
	b.preOrder(node.Right)
}

// Minimum BST的最小值
func (b *BinarySearchTree) Minimum() int {
	if b.Root == nil {
		return 0
	}
	return b.minimum(b.Root).Val
}

func (b *BinarySearchTree) minimum(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node
	}
	return b.minimum(node.Left)
}

// Maximum BST的最大值
func (b *BinarySearchTree) Maximum() int {
	if b.Root == nil {
		return 0
	}
	return b.minimum(b.Root).Val
}

func (b *BinarySearchTree) maximum(node *TreeNode) *TreeNode {
	if node.Right == nil {
		return node
	}
	return b.maximum(node.Right)
}

// RemoveMin 删除BST的最小值
func (b *BinarySearchTree) RemoveMin() int {
	ret := b.Minimum()

	b.Root = b.removeMin(b.Root)

	return ret
}

func (b *BinarySearchTree) removeMin(node *TreeNode) *TreeNode {
	if node.Left == nil {
		rightNode := node.Right
		node.Right = nil
		return rightNode
	}
	node.Left = b.removeMin(node.Left)
	return node
}

// RemoveMax 删除BST的最大值
func (b *BinarySearchTree) RemoveMax() int {
	ret := b.Maximum()

	b.Root = b.removeMax(b.Root)

	return ret
}

func (b *BinarySearchTree) removeMax(node *TreeNode) *TreeNode {
	if node.Right == nil {
		leftNode := node.Left
		node.Left = nil
		return leftNode
	}
	node.Right = b.removeMax(node.Right)
	return node
}

func (b *BinarySearchTree) Remove(val int) {
	b.Root = b.remove(b.Root, val)
}

func (b *BinarySearchTree) remove(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if val < node.Val {
		node.Left = b.remove(node.Left, val)
		return node
	} else if val > node.Val {
		node.Right = b.remove(node.Right, val)
		return node
	} else { // val = node.Val
		if node.Left == nil {
			rightNode := node.Right
			node.Right = nil
			return rightNode
		}
		if node.Right == nil {
			leftNode := node.Left
			node.Left = nil
			return leftNode
		}

		// 找出node的后继节点
		successor := b.minimum(node.Right)
		// 移除node.Right的最小值节点
		successor.Right = b.removeMin(node.Right)
		successor.Left = node.Left

		node.Left, node.Right = nil, nil
		return successor
	}
}

type MultiNode struct {
	Id  int
	Pid int
}

type MultiTreeNode struct {
	Node     *MultiNode
	Children []*MultiTreeNode
}

// 构造多叉树
func buildMultiTree(nodes []*MultiNode) (*MultiTreeNode, error) {
	if len(nodes) == 0 {
		return nil, nil
	}

	m := make(map[int]*MultiTreeNode)

	for _, node := range nodes {
		m[node.Id] = &MultiTreeNode{Node: node}
	}

	var root *MultiTreeNode

	for _, node := range nodes {
		curNode := m[node.Id]
		if node.Pid == 0 {
			root = curNode
		} else {
			parentNode, exist := m[node.Pid]
			if !exist {
				return nil, errors.New("invalid nodes")
			}
			parentNode.Children = append(parentNode.Children, curNode)
		}
	}

	return root, nil
}
