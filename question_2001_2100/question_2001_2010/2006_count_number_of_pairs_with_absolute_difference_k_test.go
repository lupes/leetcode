package question_2001_2010

import (
	"testing"
)

func Test_countKDifference(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 2, 1}, 1, 4},
		{[]int{1, 3}, 3, 0},
		{[]int{3, 2, 1, 5, 4}, 2, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countKDifference(tt.nums, tt.k); got != tt.want {
				t.Errorf("countKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
