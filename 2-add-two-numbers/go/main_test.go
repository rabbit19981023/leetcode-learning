package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1 - Basic Case",
			l1:       &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:       &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name:     "Example 2 - Both Zero",
			l1:       &ListNode{Val: 0},
			l2:       &ListNode{Val: 0},
			expected: &ListNode{Val: 0},
		},
		{
			name:     "Example 3 - Single Digit",
			l1:       &ListNode{Val: 2},
			l2:       &ListNode{Val: 4},
			expected: &ListNode{Val: 6},
		},
		{
			name:     "Example 4 - Carry Over",
			l1:       &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			l2:       &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			expected: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1}}}},
		},
		{
			name:     "Example 5 - Different Lengths",
			l1:       &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			l2:       &ListNode{Val: 9},
			expected: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}},
		},
		{
			name:     "Example 6 - Large Numbers",
			l1:       &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}},
			l2:       &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}},
			expected: &ListNode{Val: 6, Next: &ListNode{Val: 5, Next: &ListNode{Val: 8, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := addTwoNumbers(test.l1, test.l2)

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
