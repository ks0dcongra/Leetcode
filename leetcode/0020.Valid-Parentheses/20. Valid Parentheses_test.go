package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question20 struct {
	name20 string
	para20
	ans20
}

// para 是参数
// one 代表第一个参数
type para20 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans20 struct {
	one bool
}

func Test_Problem20(t *testing.T) {

	qs := []question20{

		{
			"1",
			para20{"()[]{}"},
			ans20{true},
		},
		{
			"2",
			para20{"(]"},
			ans20{false},
		},
		{
			"3",
			para20{"({[]})"},
			ans20{true},
		},
		{
			"4",
			para20{"(){[({[]})]}"},
			ans20{true},
		},
		{
			"5",
			para20{"((([[[{{{"},
			ans20{false},
		},
		{
			"6",
			para20{"(())]]"},
			ans20{false},
		},
		{
			"7",
			para20{""},
			ans20{true},
		},
		{
			"8",
			para20{"["},
			ans20{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")

	for _, q := range qs {
		t.Run(q.name20, func(t *testing.T) {
			a, p := q.ans20, q.para20
			assert.Equal(t, a.one, isValid(p.one))
		})
	}

}
