package question_1491_1500

import (
	"testing"
)

func Test_longestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
		{[]int{1, 1, 0, 0, 1, 1, 1, 0, 1}, 4},
		{[]int{0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := longestSubarray(tt.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
