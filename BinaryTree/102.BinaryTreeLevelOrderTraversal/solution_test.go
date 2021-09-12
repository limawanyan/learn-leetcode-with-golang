package _02_BinaryTreeLevelOrderTraversal

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	tree :=	&TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 3}
	tree.Left.Left = &TreeNode{Val: 4}
	tree.Left.Right = &TreeNode{Val: 5}
	tree.Right.Left = &TreeNode{Val: 6}
	tree.Right.Right = &TreeNode{Val: 7}
	fmt.Println(zigzagLevelOrder(tree))
}
