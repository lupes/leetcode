package question_671_680

import (
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{1, 3, 5, 7}, 4},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}, 4},
		{[]int{1, 1, 2, 1, 2, 3, 1, 2, 3, 4, 1}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findLengthOfLCIS(tt.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
