package twoSum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(num int) *ListNode {
	listNode := &ListNode{}
	child := listNode
	for num > 0 {
		child.Val = num % 10
		num = num / 10
		if num != 0 {
			child.Next = &ListNode{}
		}
		child = child.Next
	}
	return listNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := l1
	for l1 != nil || l2 != nil {
		if l2 != nil {
			l1.Val += l2.Val
			l2 = l2.Next
		}
		if l1.Val >= 10 {
			l1.Val %= 10
			if l1.Next == nil {
				l1.Next = &ListNode{Val: 1}
			} else {
				l1.Next.Val += 1
			}
		}
		if l1.Next == nil {
			l1.Next = l2
			return res
		}
		l1 = l1.Next
	}
	return res
}

//	Input: l1 = [2,4,3], l2 = [5,6,4]
//	Output: [7,0,8]
//	Explanation: 342 + 465 = 807.
func TestTwoSum(t *testing.T) {
	expectedNumber := 9999000000
	l1 := New(9)
	l2 := New(9998999991)
	expected := New(expectedNumber)

	actual := addTwoNumbers(l1, l2)
	for expectedNumber > 0 {
		require.Equal(t, expected.Val, actual.Val)
		expected = expected.Next
		actual = actual.Next
		expectedNumber /= 10
	}
}
