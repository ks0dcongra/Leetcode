package main

import "fmt"

func main() {
	s := "()"
    
	fmt.Println(isValid(s))

    s = "()[]{}"
    
	fmt.Println(isValid(s))

    s = "[]"
    
	fmt.Println(isValid(s))

    s = "["
    
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s) == 0 {
         return false 
    }

    stack := make([]rune,0)
    for _, v := range s {
        if v == '[' || v == '(' || v == '{' {
            stack = append(stack, v)
        } else if (v == ')' && len(stack) > 0 && stack[len(stack)-1] =='(') || 
        (v == '}' && len(stack) > 0 && stack[len(stack)-1] =='{')  || 
        (v == ']' && len(stack) > 0 && stack[len(stack)-1] =='[') { 
            stack = stack[:len(stack)-1]
        } else {
            return false
        }
    } 
 
    if len(stack) != 0 {
        return false
    }

    return true
}