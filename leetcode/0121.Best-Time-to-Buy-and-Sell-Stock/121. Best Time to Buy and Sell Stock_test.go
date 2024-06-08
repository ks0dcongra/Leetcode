package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question121 struct {
	name121 string
	para121
	ans121
}

// para 是参数
// one 代表第一个参数
type para121 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans121 struct {
	one int
}

func Test_Problem121(t *testing.T) {

	qs := []question121{

		{
			"1",
			para121{[]int{}},
			ans121{0},
		},

		{
			"2",
			para121{[]int{7, 1, 5, 3, 6, 4}},
			ans121{5},
		},

		{
			"3",
			para121{[]int{7, 6, 4, 3, 1}},
			ans121{0},
		},

		{
			"4",
			para121{[]int{1, 3, 2, 8, 4, 9}},
			ans121{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 121------------------------\n")

	// for _, q := range qs {
	// 	_, p := q.ans121, q.para121
	// 	fmt.Printf("【input】:%v       【output】:%v\n", p, maxProfit(p.one))
	// }
	for _, q := range qs {
		t.Run(q.name121, func(t *testing.T) {
			a, p := q.ans121, q.para121
			assert.Equal(t, a.one, maxProfit(p.one))
		})
	}
	fmt.Printf("\n\n\n")
}
