package question_141_150

import (
	"fmt"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_reorderList(t *testing.T) {
	tests := []struct {
		head *ListNode
	}{
		{nil},
		{NewList(1)},
		{NewList(1, 2)},
		{NewList(1, 2, 3)},
		{NewList(1, 2, 3, 4)},
		{NewList(1, 2, 3, 4, 5)},
		{NewList(1, 2, 3, 4, 5, 6)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			fmt.Printf("%+v\n", tt.head)
			reorderList(tt.head)
			fmt.Printf("%+v\n", tt.head)
		})
	}
}
