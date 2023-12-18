package array

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	反转数组，原地反转一个 char[] 类型的字符数组
*/

func init() {
	exec.Register(344, ReverseStringExec)
}

// reverseString
func reverseString(arr []byte) []byte {
	if len(arr) == 0 {
		return []byte{}
	}

	left, right := 0, len(arr)-1
	for left < right {
		// tmp := arr[left]
		// arr[left] = arr[right]
		// arr[right] = tmp
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func ReverseStringExec() {
	bytes := []byte("abcdef")
	fmt.Printf("origin str: %s\n", bytes)
	reversedStr := reverseString(bytes)
	fmt.Printf("reversed str: %s\n", reversedStr)
}
