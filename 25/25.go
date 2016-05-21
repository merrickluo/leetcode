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

func reverseKGroup(head *ListNode, k int) *ListNode {
	last := nodeAtPosition(head, k-1)

	if last == nil {
		return head
	}

	next := last.Next

	last.Next = nil
	result := reverseList(head)

	if next != nil {
		head.Next = reverseKGroup(next, k)
	}

	return result
}

func reverseList(head *ListNode) *ListNode {
	node := head
	var prev *ListNode = nil
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

func nodeAtPosition(head *ListNode, k int) *ListNode {
	node := head
	var i = 0
	for ; i < k && node != nil; i++ {
		node = node.Next
	}
	return node
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}
	printList(head)
	fmt.Println()
	printList(reverseKGroup(head, 3))
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print(" ")
		node = node.Next
	}
}
