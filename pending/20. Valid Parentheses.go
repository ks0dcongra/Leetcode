package main

// import (
// 	"fmt"
// 	"log"
// )

// type ListNode struct {
// 	Val  *int
// 	Next *ListNode
// }

// func main() {
// 	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
// 	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

// 	l1 = &ListNode{Val: nil, Next: nil}
// 	l2 = &ListNode{Val: 0, Next: nil}

// 	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
// 	l2 = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
// }

// func main() {
// 	s := "()"
	
// 	fmt.Println(isValid(s))

// 	s = "()[]{}"

// 	fmt.Println(isValid(s))

// 	s = "(]"

// 	fmt.Println(isValid(s))
// }

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	stack := make([]rune, 0)
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
		} else if (v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			(v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return true
}
