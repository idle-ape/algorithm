package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

func init() {
	exec.Register(876, FindMiddleNodeExec)
}

/*
	找出单链表的中间节点
	解法：通过快慢指针，慢指针从头节点开始每次走一步，快指针从头节点开始每次都两步，当快指针到末尾时，慢指针此时正好在中间的节点
*/

func findMiddleNode(head *List) *List {
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		if p2.Next != nil {
			p2 = p2.Next.Next
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func FindMiddleNodeExec() {
	fmt.Printf("list: %s\n", List1)
	node := findMiddleNode(List1)
	fmt.Printf("middle node: %d\n", node.Val)

	fmt.Printf("list: %s\n", List3)
	node = findMiddleNode(List3)
	fmt.Printf("middle node: %d\n", node.Val)
}
