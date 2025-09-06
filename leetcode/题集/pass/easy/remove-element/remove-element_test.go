package leetcode

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
