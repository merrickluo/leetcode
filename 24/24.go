/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	nowHead := head.Next
	nextHead := nowHead.Next
	nowHead.Next = head
	head.Next = swapPairs(nextHead)
	return nowHead
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	printList(head)
	fmt.Println()
	printList(swapPairs(head))
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print(" ")
		node = node.Next
	}
}
