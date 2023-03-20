package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
  	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    current := dummy
    carry := 0

    for l1 != nil || l2 != nil {
        x, y := 0, 0
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }
        sum := x + y + carry
        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }

    if carry > 0 {
        current.Next = &ListNode{Val: carry}
    }

    return dummy.Next
}

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Expected Output: [7,0,8]
// Explanation: 342 + 465 = 807.

func main() {
    // Constructing the input linked lists
    l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
    l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

    // Computing the result using our function
    result := addTwoNumbers(l1, l2)

    // Printing the result
    for result != nil {
        fmt.Print(result.Val)
        if result.Next != nil {
            fmt.Print(" -> ")
        }
        result = result.Next
    }
}
