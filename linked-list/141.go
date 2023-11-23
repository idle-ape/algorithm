package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

func init() {
	exec.Register(141, CycleListExec)
}

/*
	判断单链表中是否有环
	解法：通过快慢指针，慢指针从头节点开始每次走一步，快指针从头节点开始每次都两步，如果快慢指针有重合，说明链表中存在环，否则不存在
*/

func cycleList(head *List) bool {
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		if p2.Next != nil {
			p2 = p2.Next.Next
		} else {
			p2 = p2.Next
		}
		if p1 == p2 {
			return true
		}
	}
	return false
}

func CycleListExec() {
	// 构建带环的链表
	CycleList.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node3
	hasCycle := cycleList(CycleList)
	fmt.Printf("has cycle in list: %v\n", hasCycle)
}
