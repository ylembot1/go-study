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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}

	p := res
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		t := p2.Val
		if p1.Val < p2.Val {
			t = p1.Val
			p1 = p1.Next
		} else {
			p2 = p2.Next
		}
		p.Next = &ListNode{
			Val:  t,
			Next: nil,
		}
		p = p.Next
	}

	pp := p1
	if pp == nil {
		pp = p2
	}
	for pp != nil {
		p.Next = &ListNode{
			Val:  pp.Val,
			Next: nil,
		}
		p = p.Next
		pp = pp.Next
	}

	return res.Next
}
