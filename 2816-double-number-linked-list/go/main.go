package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleListNode(head *ListNode) *ListNode {
	carry := doubleFromTail(head)

	if carry > 0 {
		newHead := &ListNode{Val: carry}
		newHead.Next = head

		return newHead
	}

	return head
}

func doubleFromTail(node *ListNode) int {
	if node == nil {
		return 0
	}

	carry := doubleFromTail(node.Next)

	newVal := node.Val*2 + carry
	node.Val = newVal % 10

	return newVal / 10
}
