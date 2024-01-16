package dynamicprogramming

import (
	"fmt"
	"math"

	"github.com/idle-ape/algorithm/exec"
	"golang.org/x/exp/constraints"
)

/*
	凑零钱问题：给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，
	问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1。
*/

func init() {
	exec.Register(322, CoinExchangeExec)
}

func coinExchange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := math.MaxInt32

	// 假设你有面值为 1, 2, 5 的硬币，你想求 amount = 11 时的最少硬币数，
	// 如果你知道凑出 amount = 10, 9, 6 的最少硬币数（子问题），你只需要把子问题的答案加一（再选一枚面值为 1, 2, 5 的硬币），
	// 求个最小值，就是原问题的答案

	for _, coin := range coins {
		sub := coinExchange(coins, amount-coin)
		if sub == -1 {
			continue
		}
		res = min(res, sub+1)
	}

	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func CoinExchangeExec() {
	coins, amount := []int{1, 2, 5}, 11
	fmt.Printf("res: %d\n", coinExchange(coins, amount))
}

func min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	}
	return a
}
