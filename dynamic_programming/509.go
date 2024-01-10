package dynamicprogramming

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
 斐波那契数列：1 1 2 3 5 8 13 21 ...

 采用自底向上的思路，又最基础的case，即dp[0]和dp[1]往上推导出对应的值
*/

func init() {
	exec.Register(509, FibExec)
}

func fib(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func FibExec() {
	fmt.Printf("fib(%d): %d\n", 5, fib(5))
	fmt.Printf("fib(%d): %d\n", 8, fib(8))
	fmt.Printf("fib(%d): %d\n", 9, fib(9))
}
