package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

func init() {
	exec.Register(142, StartOfCycleListExec)
}

/*
	判断单链表中是否有环，如果有，环的起点是哪个节点
	解法：
		通过快慢指针，慢指针从头节点开始每次走一步，快指针从头节点开始每次都两步，如果快慢指针有重合，说明链表中存在环，否则不存在
		上面的步骤和第 141 题的判断链表中是否有环一样，难点在于如何获取环的起点。
		1、假设快慢指针相遇时，慢指针走了K步，则快指针一定是走了2K步，快指针比慢指针多走了K步
		2、假设环的起点距离相遇点为M，那么环起点距离链表起点的距离即为K-M，在相遇点再走K-M步也会到环的起点，所以，在快慢指针相遇的时候，
			让其中一个指针从链表头开始和另一个指针同速再走K-M，两个指针相遇时即为环的起点

		https://labuladong.gitee.io/algo/di-yi-zhan-da78c/shou-ba-sh-8f30d/shuang-zhi-0f7cc/
*/

func startOfcycleList(head *List) (bool, *List) {
	slow, fast := head, head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if hasCycle {
		// 同速前进
		slow = head
		for slow != nil && fast != nil {
			slow = slow.Next
			fast = fast.Next
			if slow == fast {
				return true, slow
			}
		}
	}
	return false, nil
}

func StartOfCycleListExec() {
	// 构建带环的链表
	//  1 - 2 - 5 - 2 - 4 - 3 - -
	// 				|			｜
	// 				| 			｜
	// 				| _	_ _ _ _	｜
	CycleList.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node3
	hasCycle, startNode := startOfcycleList(CycleList)
	fmt.Printf("has cycle in list: %v, start node of cycle: %d\n", hasCycle, startNode.Val)
}
