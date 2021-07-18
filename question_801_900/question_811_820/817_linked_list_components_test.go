package question_811_820

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_numComponents(t *testing.T) {
	tests := []struct {
		head *ListNode
		G    []int
		want int
	}{
		{NewList(0), []int{0}, 1},
		{NewList(0, 1, 2), []int{0, 2}, 2},
		{NewList(0, 1, 2, 3), []int{0, 1, 3}, 2},
		{NewList(0, 1, 2, 3, 4), []int{0, 3, 1, 4}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numComponents(tt.head, tt.G); got != tt.want {
				t.Errorf("numComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
