package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDoubleListNode(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1 Basic Case",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}},
			expected: &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8}}},
		},
		{
			name:     "Example 2 Multiple 9s",
			head:     &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 8}}}},
		},
		{
			name:     "Example 3 Single Node",
			head:     &ListNode{Val: 5},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 0}},
		},
		{
			name:     "Example 4 Zero Node",
			head:     &ListNode{Val: 0},
			expected: &ListNode{Val: 0},
		},
		{
			name:     "Example 5 No Carry",
			head:     &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			expected: &ListNode{Val: 8, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2}}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := doubleListNode(test.head)

			if !listEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", listToString(test.expected), listToString(result))
			}
		})
	}
}

func listEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}

func listToString(l *ListNode) string {
	var sb strings.Builder

	for l != nil {
		fmt.Fprintf(&sb, "%d", l.Val)

		if l.Next != nil {
			sb.WriteString(" -> ")
		}

		l = l.Next
	}

	return sb.String()
}
