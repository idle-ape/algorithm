package binarytree

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	二叉树的层序遍历
*/

var levelOrderBST = &BinaryTree{
	Val: 3,
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

func init() {
	exec.Register(102, LevelOrderExec)
}

func levelOrder(bst *BinaryTree) [][]int {
	if bst == nil {
		return [][]int{}
	}

	// 保存层序遍历的结果
	elems := make([][]int, 0)

	nodes := []*BinaryTree{bst}
	for len(nodes) > 0 {
		e := []int{}
		nextLevelNodes := []*BinaryTree{}
		for i := 0; i < len(nodes); i++ {
			e = append(e, nodes[i].Val)
			if nodes[i].Left != nil {
				nextLevelNodes = append(nextLevelNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nextLevelNodes = append(nextLevelNodes, nodes[i].Right)
			}
		}
		elems = append(elems, e)
		nodes = nextLevelNodes
	}

	return elems
}

func LevelOrderExec() {
	elems := levelOrder(levelOrderBST)
	fmt.Printf("level order elements: %v\n", elems)
}
