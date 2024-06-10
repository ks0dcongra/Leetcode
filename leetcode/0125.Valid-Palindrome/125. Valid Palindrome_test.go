package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem125(t *testing.T) {

	tcs := []struct {
		name string
		s   string
		ans bool
	}{

		{
			"1",
			"0p",
			false,
		},

		{
			"2",
			"0",
			true,
		},

		{
			"3",
			"race a car",
			false,
		},

		{
			"4",
			"A man, a plan, a canal: Panama",
			true,
		},

		{
			"5",
			".,",
			true,
		},
	}
	fmt.Printf("------------------------Leetcode Problem 125------------------------\n")

	// for _, tc := range tcs {
	// 	fmt.Printf("【input】:%v       【output】:%v\n", tc, isPalindrome(tc.s))
	// }
	for _, q := range tcs {
		t.Run(q.name, func(t *testing.T) {
			expect := q.ans
			assert.Equal(t, expect, isPalindrome(q.s))
		})
	}

	fmt.Printf("\n\n\n")
}
