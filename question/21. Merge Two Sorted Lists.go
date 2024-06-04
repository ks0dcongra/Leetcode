package main

import "fmt"

func main() {
    nums := []int{2,7,11,15}
    target := 9
    fmt.Println(twoSum(nums, target))

    nums = []int{3,2,4}
    target = 6
    fmt.Println(twoSum(nums, target))

    nums = []int{3,3}
    target = 6
    fmt.Println(twoSum(nums, target))
}

    /**
    * Definition for singly-linked list.
    * type ListNode struct {
    *     Val int
    *     Next *ListNode
    * }
    */
	func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}
		if l1.Val < l2.Val {
			l1.Next = mergeTwoLists(l1.Next, l2)
			return l1
		}
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}