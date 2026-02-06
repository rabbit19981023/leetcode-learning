package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleListNode(head *ListNode) *ListNode {
	reversed := reverseList(head)

	curr := reversed
	carry := 0

	for curr != nil {
		newVal := curr.Val*2 + carry
		curr.Val = newVal % 10
		carry = newVal / 10

		if curr.Next == nil {
			break
		}

		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return reverseList(reversed)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = prev
		prev = curr

		curr = next
	}

	return prev
}
