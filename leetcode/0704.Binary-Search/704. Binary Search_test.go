package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question704 struct {
	name704 string
	para704
	ans704
}

// para 是参数
// one 代表第一个参数
type para704 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans704 struct {
	one int
}

func Test_Problem704(t *testing.T) {

	qs := []question704{

		{
			"1",
			para704{[]int{-1, 0, 3, 5, 9, 12}, 9},
			ans704{4},
		},

		{
			"2",
			para704{[]int{-1, 0, 3, 5, 9, 12}, 2},
			ans704{-1},
		},
		{
			"3",
			para704{[]int{1, 2, 3}, 6},
			ans704{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 704------------------------\n")

	// for _, q := range qs {
	// 	_, p := q.ans704, q.para704
	// 	fmt.Printf("【input】:%v       【output】:%v\n", p, search704(p.nums, p.target))
	// }
	for _, q := range qs {
		t.Run(q.name704, func(t *testing.T) {
			a, p := q.ans704, q.para704
			assert.Equal(t, a.one, search704(p.nums, p.target))
		})
	}
	fmt.Printf("\n\n\n")
}
