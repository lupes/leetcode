package question_671_680

import (
	"testing"
)

func Test_findNumberOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 3, 5, 4, 7}, 4},
		{[]int{2, 2, 2, 2, 2}, 5},
		{[]int{1, 2, 2, 2, 2, 2}, 5},
		{[]int{1, 2, 4, 3, 5, 4, 7, 2}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findNumberOfLIS(tt.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
