package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, q := head, head.Next
	head.Next = nil
	for q != nil {
		if q.Val != p.Val {
			p.Next = q
			p = q
		}
		q = q.Next
	}
	p.Next = nil

	return head
}
