package leetcode

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}
	p := deleteDuplicates(a)
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
