package my_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_string(t *testing.T) {
	fmt.Println(strings.Join([]string{"a", "b"}, " "))

	fmt.Println(strings.Replace("leetcode, liutao", "l", "L", -1))
}
