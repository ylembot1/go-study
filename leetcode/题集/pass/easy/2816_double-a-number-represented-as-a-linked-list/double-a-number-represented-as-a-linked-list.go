package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func dep(p *ListNode) (*ListNode, int) {
	if p == nil {
		return nil, 0
	}

	q, flag := dep(p.Next)

	value := p.Val*2 + flag
	p.Val = value % 10
	flag = value / 10
	p.Next = q

	return p, flag
}

func doubleIt(head *ListNode) *ListNode {
	p, flag := dep(head)

	if flag == 0 {
		return p
	}
	return &ListNode{Val: flag, Next: p}
}
