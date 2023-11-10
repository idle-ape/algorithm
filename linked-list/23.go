package linkedlist

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	合并多个有序链表
*/

func init() {
	exec.Register(23, MergeMultiListExec)
}

func mergeMultiList(list []*List) *List {
	return nil
}

func MergeMultiListExec() {
	fmt.Printf("list: %s\n", List3)
	l := mergeMultiList([]*List{List3, List1})
	fmt.Printf("merge done: %s\n", l)
}
