package main

// import "log"
	
// func main() {
// 	l1 := &ListNode{
// 		Val: 2,
// 		Next: &ListNode{
// 			Val: 4,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: nil,
// 			},
// 		},
// 	}

// 	l2 := &ListNode{
// 		Val: 5,
// 		Next: &ListNode{
// 			Val: 6,
// 			Next: &ListNode{
// 				Val: 4,
// 				Next: nil,
// 			},
// 		},
// 	}

// 	result := addTwoNumbers(l1, l2)
// 	PrintLinkedList(result)
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	head := &ListNode{Val: 0}
// 	n1, n2, carry, current := 0, 0, 0, head
// 	for l1 != nil || l2 != nil || carry != 0 {
// 		if l1 == nil {
// 			n1 = 0
// 		} else {
// 			n1 = l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 == nil {
// 			n2 = 0
// 		} else {
// 			n2 = l2.Val
// 			l2 = l2.Next
// 		}
// 		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
// 		current = current.Next
// 		carry = (n1 + n2 + carry) / 10
// 	}
// 	log.Println(head.Next)
// 	return head.Next
// }

// func PrintLinkedList(head *ListNode) {
// 	if head == nil {
// 		log.Println("[]")
// 		return
// 	}
// 	result := []int{}
// 	current := head
// 	for current != nil {
// 		result = append(result, current.Val)
// 		current = current.Next
// 	}
// 	log.Println(result)
// }