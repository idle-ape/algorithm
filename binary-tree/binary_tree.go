package binarytree

var maxDepth = &BinaryTree{
	Val: 3,
	Left: &BinaryTree{
		Val: 2,
		Left: &BinaryTree{
			Val: 6,
		},
		Right: &BinaryTree{
			Val: 1,
			Left: &BinaryTree{
				Val: 4,
			},
		},
	},
	Right: &BinaryTree{
		Val: 5,
	},
}

// BinaryTree define of binary tree
type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}
