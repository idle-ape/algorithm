package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

func init() {
	exec.Register(86, PartitionListExec)
}

/*
	给你一个链表的头节点 head 和一个特定值 x ，对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前
	你应当 保留 两个分区中每个节点的初始相对位置。
	比如：
		1 - 4 - 3 - 2 - 5 - 2	=>  1 - 2 - 2 - 4 - 3 - 5
*/

func partitionList(head *List, x int) *List {
	dummy1, dummy2 := &List{Val: -1}, &List{Val: -1}
	p1, p2 := dummy1, dummy2

	p := head
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 这里不能直接让链表的指针前进，因为这样直接前进的话最后会形成一个环，遍历的时候就死循环了
		// p = p.Next
		// 断开原链表中的每个节点的Next指针
		t := p.Next
		p.Next = nil
		p = t
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

func PartitionListExec() {
	fmt.Printf("list: %s\n", List3)
	l := partitionList(List3, 3)
	fmt.Printf("partion done: %s\n", l)
}
