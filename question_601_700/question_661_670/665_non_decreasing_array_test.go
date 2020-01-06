package question_661_670

import (
	"testing"
)

func Test_checkPossibility(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{4}, true},
		{[]int{4, 2}, true},
		{[]int{4, 2, 3}, true},
		{[]int{-1, 4, 2, 3}, true},
		{[]int{1, 2, 3}, true},
		{[]int{2, 1, 3}, true},
		{[]int{4, 2, 1}, false},
		{[]int{3, 4, 2, 3}, false},
		{[]int{2, 3, 3, 2, 4}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkPossibility(tt.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
