package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question242 struct {
	name242 string
	para242
	ans242
}

// para 是参数
// one 代表第一个参数
type para242 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans242 struct {
	one bool
}

func Test_Problem242(t *testing.T) {

	qs := []question242{

		{
			"1",
			para242{"", ""},
			ans242{true},
		},
		{
			"2",
			para242{"", "1"},
			ans242{false},
		},

		{
			"3",
			para242{"anagram", "nagaram"},
			ans242{true},
		},

		{
			"4",
			para242{"rat", "car"},
			ans242{false},
		},

		{
			"5",
			para242{"a", "ab"},
			ans242{false},
		},

		{
			"6",
			para242{"ab", "a"},
			ans242{false},
		},

		{
			"7",
			para242{"aa", "bb"},
			ans242{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 242------------------------\n")

	// for _, q := range qs {
	// 	_, p := q.ans242, q.para242
	// 	fmt.Printf("【input】:%v       【output】:%v\n", p, isAnagram(p.one, p.two))
	// }
	for _, q := range qs {
		t.Run(q.name242, func(t *testing.T) {
			a, p := q.ans242, q.para242
			assert.Equal(t, a.one, isAnagram1(p.one, p.two))
		})
	}
	fmt.Printf("\n\n\n")
}
