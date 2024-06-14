package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

type question110 struct {
	name string
	para110
	ans110
}

// para 是参数
// one 代表第一个参数
type para110 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans110 struct {
	one bool
}

func Test_Problem110(t *testing.T) {

	qs := []question110{

		{
			"1",
			para110{[]int{3, 4, 4, 5, structures.NULL, structures.NULL, 5, 6, structures.NULL, structures.NULL, 6}},
			ans110{false},
		},

		{
			"2",
			para110{[]int{1, 2, 2, structures.NULL, 3, 3}},
			ans110{true},
		},

		{
			"3",
			para110{[]int{}},
			ans110{true},
		},

		{
			"4",
			para110{[]int{1}},
			ans110{true},
		},

		{
			"5",
			para110{[]int{1, 2, 3}},
			ans110{true},
		},

		{
			"6",
			para110{[]int{1, 2, 2, 3, 4, 4, 3}},
			ans110{true},
		},

		{
			"7",
			para110{[]int{1, 2, 2, structures.NULL, 3, structures.NULL, 3}},
			ans110{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 110------------------------\n")

	// for _, q := range qs {
	// 	_, p := q.ans110, q.para110
	// 	fmt.Printf("【input】:%v      ", p)
	// 	rootOne := structures.Ints2TreeNode(p.one)
	// 	fmt.Printf("【output】:%v      \n", isBalanced(rootOne))
	// }
	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans110, q.para110
			rootOne := structures.Ints2TreeNode(p.one)
			assert.Equal(t, a.one, isBalanced(rootOne))
		})
	}
	fmt.Printf("\n\n\n")
}
