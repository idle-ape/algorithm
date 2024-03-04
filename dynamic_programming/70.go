package dynamicprogramming

import "github.com/idle-ape/algorithm/exec"

/*
	爬楼梯问题
		假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
		每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

func init() {
	exec.Register(70, ClimbingStairsExec)
}

func climbingStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func ClimbingStairsExec() {

}
