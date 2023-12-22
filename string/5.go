package string

import (
	"fmt"

	"github.com/idle-ape/algorithm/exec"
)

/*
	最长回文子串

	给你一个字符串 s，找到 s 中最长的回文子串。

	如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

	abcgcdfoq
*/

func init() {
	exec.Register(5, LongestPalindromeExec)
}

// longestPalindrome
func longestPalindrome(s string) string {
	ps := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(s2) && len(s1) > len(ps) {
			ps = s1
		}

		if len(s1) < len(s2) && len(s2) > len(ps) {
			ps = s2
		}
	}
	return ps
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func LongestPalindromeExec() {
	s := "abcgcdfoq"
	ps := longestPalindrome(s)
	fmt.Println(ps)
}
