package question_1691_1700

import (
	"testing"
)

func Test_maxResult(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, -1, -2, 4, -7, 3}, 2, 7},
		{[]int{10, -5, -2, 4, 0, 3}, 3, 17},
		{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2, 0},
		{[]int{1, -5, -20, 4, -1, 3, -6, -3}, 2, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxResult(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
