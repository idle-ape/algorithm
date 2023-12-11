package array

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。如：
	[0 3 1 2 3 3 5], 3 -> [0 5 1 2 _ _]

	https://labuladong.gitee.io/algo/di-ling-zh-bfe1b/shuang-zhi-fa4bd/#div_remove-duplicates-from-sorted-array
*/

func init() {
	exec.Register(27, DelArraySameEleExec)
}

func delArraySameEle(arr []int, v int) int {
	if len(arr) == 0 {
		return 0
	}

	slow, fast := 0, 0
	for fast < len(arr) {
		if arr[fast] != v {
			arr[slow] = arr[fast]
			slow++
		}
		fast++
	}
	return slow
}

func DelArraySameEleExec() {
	arr := []int{1, 3, 2, 2, 3}
	n := delArraySameEle(arr, 3)
	fmt.Printf("after del, len: %d, eles: %v\n", n, arr)
}
