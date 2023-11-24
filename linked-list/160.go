package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

func init() {
	exec.Register(160, GetIntersectionNodeExec)
}

/*
	判断两条链表是相交，如果相交则返回相交节点

		1 - 2 - 3
				\
					4 - 5
				/
	8 -	2 - 0 - 6

	1、简单的解法，遍历其中一个链表，将每个节点放到hash表中，然后遍历第二个链表时，判断其节点有没有在hash表中，有则说明存在相交的节点

	2、取巧一点的解法，将两个链表拼接成出两个长链表，通过双指针在两个长链表上移动，双指针则能同时到达相交的节点

	p1
	↓
	1 - 2 - 3 - 4 - 5 - 8 - 2 - 0 - 6 - [4] - 5
	8 - 2 - 0 - 6 - 4 - 5 - 1 - 2 - 3 - [4] - 5
	↑									 ↑
	p2								  相交节点

*/

func getIntersectionNode(head1, head2 *List) *List {
	p1 := head1
	p2 := head2

	// 这里不能通过遍历的方式拿到每个链表的末尾节点，然后真实的将两个链表拼接在一起，这样会形成一个环！

	for p1 != p2 {
		if p1 == nil {
			p1 = head2
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = head1
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

func GetIntersectionNodeExec() {
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	n2 := &List{
		Val: 1,
		Next: &List{
			Val: 8,
			Next: &List{
				Val:  6,
				Next: node3,
			},
		},
	}

	fmt.Printf("n1: %s\n", node1)
	fmt.Printf("n2: %s\n", n2)
	node := getIntersectionNode(node1, n2)
	fmt.Printf("intersection node: %v\n", node.Val)
}
