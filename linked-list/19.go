package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

func init() {
	exec.Register(19, RemoveKthNodeExec)
}

/*
	删除链表的倒数第K个节点
*/

func removeKthNode(head *List, k int) *List {
	// 保存倒数第K个节点的上一个节点
	var kThPrev *List
	// 找到倒数第K个节点
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	for p1 != nil && p2 != nil {
		kThPrev = p2
		p1 = p1.Next
		p2 = p2.Next
	}

	// 将倒数第K个节点的上一个节点的Next指针指向倒数第K个节点的下一个节点
	kThPrev.Next = p2.Next

	return head
}

func RemoveKthNodeExec() {
	k := 3
	fmt.Printf("list: %s\n", List3)

	// 没有看labuladong之前的实现
	// l := removeKthNode(List3, k)
	// fmt.Printf("delete done: %s\n", l)

	/*
		labuladong版本的实现
		通过找到倒数第K个链表节点的方式，找到倒数第K+1的节点，然后改变倒数第K+1节点的指针的指向
	*/
	reverseOrderKThNode := findFromEnd(List3, k+1)
	reverseOrderKThNode.Next = reverseOrderKThNode.Next.Next
	fmt.Printf("delete done: %s\n", List3)
}
