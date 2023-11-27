package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

func init() {
	exec.Register(83, RemoveDumplicatNodeExec)
}

/*
	删除排序链表中的重复元素，如：
	1 - 2 - 2 - 3 - 3 - 5 - 5
	1 - 2 - 3 - 5

	思路和力扣的第26题【原地删除排序数组中的重复元素】一致，通过快慢两个指针，开始时快慢指针都指向起始位置，然后比较快慢指针对应的节点的值
	如果相等，则让慢指针节点指向快指针所在的节点，同时慢指针向后移动
*/

func removeDumplicatNode(head *List) *List {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}

	slow.Next = nil
	return head
}

func RemoveDumplicatNodeExec() {
	list := &List{
		Val: 1,
		Next: &List{
			Val: 2,
			Next: &List{
				Val: 2,
				Next: &List{
					Val: 3,
					Next: &List{
						Val: 4,
						Next: &List{
							Val: 5,
						},
					},
				},
			},
		},
	}
	fmt.Printf("delete done: %s\n", removeDumplicatNode(list))
}
