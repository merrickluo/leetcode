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

func rotateRight(head *ListNode, k int) *ListNode {
	count := countList(head)

	if count == 0 {
		return head
	}

	k = k % count
	last := nodeAtPosition(head, count-k-1)

	if last == nil || last.Next == nil {
		return head
	}

	next := last.Next
	last.Next = nil

	node := next
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head

	return next
}

func countList(head *ListNode) int {
	count := 0
	node := head
	for node != nil {
		count++
		node = node.Next
	}
	return count
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
	var head *ListNode = nil
	// head := &ListNode{1, nil}
	// head.Next = &ListNode{2, nil}
	// head.Next.Next = &ListNode{3, nil}
	// head.Next.Next.Next = &ListNode{4, nil}
	// head.Next.Next.Next.Next = &ListNode{5, nil}
	printList(head)
	fmt.Println()
	printList(rotateRight(head, 3))
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print(" ")
		node = node.Next
	}
}
