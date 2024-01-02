package binarytree

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	二叉树的最大路径和
*/

var bst = &BinaryTree{
	Val: 9,
	Left: &BinaryTree{
		Val: 6,
	},
	Right: &BinaryTree{
		Val: -3,
		Left: &BinaryTree{
			Val: -6,
		},
		Right: &BinaryTree{
			Val: 2,
			Left: &BinaryTree{
				Val: 2,
			},
		},
	},
}

var bst2 = &BinaryTree{
	Val: -10,
	Left: &BinaryTree{
		Val: 9,
	},
	Right: &BinaryTree{
		Val: 20,
		Left: &BinaryTree{
			Val: 15,
		},
		Right: &BinaryTree{
			Val: 7,
		},
	},
}

var mps = 0

func init() {
	exec.Register(124, MaxPathSumExec)
}

func maxPathSum(root *BinaryTree) int {
	if root == nil {
		return 0
	}

	// 计算左边分支最大值，左边分支如果为负数还不如不选择
	sumLeft := max(maxPathSum(root.Left), 0)
	// 计算右边分支最大值，右边分支如果为负数还不如不选择
	sumRight := max(maxPathSum(root.Right), 0)
	// left->root->right 作为路径与已经计算过历史最大值做比较
	mps = max(mps, root.Val+sumLeft+sumRight)
	// 返回经过root的单边最大分支给当前root的父节点计算使用
	return max(sumLeft, sumRight) + root.Val
}

func MaxPathSumExec() {
	maxPathSum(bst)
	fmt.Printf("max path sum: %d\n", mps)
}
