package binarytree

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	计算一棵二叉树的最长直径长度
*/

func init() {
	exec.Register(543, DiameterOfBinaryTreeExec)
}

func diameterOfBinaryTree(root *BinaryTree, d *int) int {
	if root == nil {
		return 0
	}

	dl := diameterOfBinaryTree(root.Left, d)
	dr := diameterOfBinaryTree(root.Right, d)

	// 每个节点的直径长度，等于左子树的最大深度加上右子树的最大深度
	// 而二叉树的后续位置，可以知道每个节点的左右子树的信息，所以这里通过二叉树的后续遍历，拿到左右子树的信息。
	*d = max(*d, dl+dr)

	return max(dl, dr) + 1
}

func DiameterOfBinaryTreeExec() {
	d1 := 0
	diameterOfBinaryTree(maxDepth, &d1)
	fmt.Printf("diameter of binary tree: %d\n", d1)

	d2 := 0
	diameterOfBinaryTree(diameter1, &d2)
	fmt.Printf("diameter of binary tree: %d\n", d2)
}
