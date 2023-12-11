package array

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。

	如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。

	以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

	你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

	你所设计的解决方案必须只使用常量级的额外空间。
*/

func init() {
	exec.Register(167, TwoSumExec)
}

// bourneTwoSum 知道用左右双指针，但是没有做出来，最后用map实现
func bourneTwoSum(arr []int, v int) []int {
	idxs := []int{}
	if len(arr) == 0 {
		return idxs
	}

	m := make(map[int]int, len(arr))
	for k, i := range arr {
		diff := v - i
		if idx, ok := m[diff]; ok {
			return append(idxs, idx+1, k+1)
		}
		m[i] = k
	}

	return idxs
}

// labuladongTwoSum labuladong 版解法，使用双指针解法
func labuladongTwoSum(arr []int, v int) []int {
	if len(arr) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(arr)-1
	for left < right {
		if sum := arr[left] + arr[right]; sum == v {
			return []int{left + 1, right + 1}
		} else if sum > v {
			right-- // 让 sum 小一些
		} else {
			left++ // 让 sum 大一些
		}
	}
	return []int{-1, -1}
}

func TwoSumExec() {
	arr := []int{2, 7, 11, 15, 17}
	idxs := bourneTwoSum(arr, 28)
	fmt.Printf("idxs: %v\n", idxs)
	idxs = labuladongTwoSum(arr, 28)
	fmt.Printf("idxs: %v\n", idxs)
}
