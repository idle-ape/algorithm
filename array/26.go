package array

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	原地删除有序数组中的重复的元素，如：
	[0 0 1 2 3 3 5] -> [0 1 2 3 5 _ _] 因为有两个元素是重复的，所以原数组中的最后两位是删除后的元素

	https://labuladong.gitee.io/algo/di-ling-zh-bfe1b/shuang-zhi-fa4bd/#div_remove-duplicates-from-sorted-array
*/

func init() {
	exec.Register(26, DelDescArraySameEleExec)
}

func delDescArraySameEle(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	slow, fast := 0, 0
	for fast < len(arr) {
		if arr[slow] != arr[fast] {
			slow++
			arr[slow] = arr[fast]
		}
		fast++
	}
	return slow
}

func DelDescArraySameEleExec() {
	arr := []int{0, 0, 1, 2, 2, 3, 3}
	n := delDescArraySameEle(arr)
	fmt.Printf("after del, len: %d, eles: %v\n", n, arr)
}
