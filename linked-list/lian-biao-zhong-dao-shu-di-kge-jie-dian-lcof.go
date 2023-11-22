package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	返回链表的倒数第 k 个节点
	https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

	通过两个指针，第一个指针先走K步，然后第二个指针指向链表头，
	两个指针同时走，当第一个指针走到头时，第二个指针的位置即为第K个节点
*/

func init() {
	exec.Register(0, FindFromEndExec)
}

func findFromEnd(head *List, k int) *List {
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}

func FindFromEndExec() {
	fmt.Printf("list: %s\n", List3)
	k := 3
	l := findFromEnd(List3, k)
	fmt.Printf("reverse %d node: %d\n", k, l.Val)
}
