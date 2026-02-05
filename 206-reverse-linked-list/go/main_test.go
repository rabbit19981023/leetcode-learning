package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1 - Normal Case",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			expected: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
		},
		{
			name:     "Example 2 - Two Nodes",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
		},
		{
			name:     "Example 3 - Empty List",
			head:     nil,
			expected: nil,
		},
		{
			name:     "Example 4 - Single Node",
			head:     &ListNode{Val: 1},
			expected: &ListNode{Val: 1},
		},
		{
			name:     "Example 5 - Three Nodes",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			expected: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
		},
		{
			name:     "Example 6 - Longer List",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}}}}},
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 6, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := reverseList(test.head)

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
	if l == nil {
		return "nil"
	}

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
