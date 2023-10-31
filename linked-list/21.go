package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
	l1: 1 - 2 - 4 - 6
	l2: 1 - 3 - 5 - 9
	l: 1 - 1 - 2 - 3 - 4 - 5 - 6 - 9

	技巧：
		使用虚拟头节点，即dummy节点，可以避免处理空指针的情况。当需要创造一条新链表的时候，可以使用虚拟头结点简化边界情况的处理。
		比如说，让你把两条有序链表合并成一条新的有序链表，是不是要创造一条新链表？再比你想把一条链表分解成两条链表，是不是也在创造新链表？这些情况都可以使用虚拟头结点简化边界情况的处理。
*/

func init() {
	exec.Register(21, MergeTwoListExec)
}

// MergeTwoList 合并两个有序链表
func MergeTwoList(l1, l2 *List) *List {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &List{}
	p1, p2, p := l1, l2, dummy
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

func MergeTwoListExec() {
	fmt.Printf("list1: %s\n", List1)
	fmt.Printf("list2: %s\n", List2)
	fmt.Printf("merged: %s\n", MergeTwoList(List1, List2))
}
