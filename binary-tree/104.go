package binarytree

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
	"golang.org/x/exp/constraints"
)

/*
	计算二叉树的最大深度
*/

func init() {
	exec.Register(104, MaxDepthExec)
}

func maxDepthOfBinaryTree(t *BinaryTree) int {
	if t == nil {
		return 0
	}
	depthOfLeft := maxDepthOfBinaryTree(t.Left)
	depthOfRight := maxDepthOfBinaryTree(t.Right)

	// 整棵树的最大深度等于左右子树的最大深度取最大值，
	// 然后再加上根节点自己
	return max(depthOfLeft, depthOfRight) + 1
}

func MaxDepthExec() {
	d := maxDepthOfBinaryTree(maxDepth)
	fmt.Printf("max depth: %d\n", d)
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
