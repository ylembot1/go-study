package leetcode

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 2, 4, 5}, 3))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
