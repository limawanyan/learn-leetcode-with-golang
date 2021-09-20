package Offer54_TheKthLargestNode

import (
	"fmt"
	"testing"
	. "learn-leetcode-with-golang/BinaryTree"
)

func TestKthNode(t *testing.T) {
	tree := TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 3}
	tree.Left.Left = &TreeNode{Val: 4}
	tree.Right.Right = &TreeNode{Val: 5}
	fmt.Println(kthLargest2(&tree,2))
}