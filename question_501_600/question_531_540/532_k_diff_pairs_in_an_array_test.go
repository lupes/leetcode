package question_531_540

import (
	"testing"
)

func Test_findPairs(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{3, 1, 4, 5, 2}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
		{[]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3, 2},
		{[]int{-1, -2, -3}, 1, 2},
		{[]int{2, 8, 6, 9, 7, 4, 9, 0, 5, 4}, 1, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findPairs(tt.nums, tt.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
