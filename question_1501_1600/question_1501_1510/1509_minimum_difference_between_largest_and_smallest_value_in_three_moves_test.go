package question_1501_1510

import (
	"testing"
)

func Test_minDifference(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{5, 3, 2, 4}, 0},
		{[]int{1, 5, 0, 10, 14}, 1},
		{[]int{6, 6, 0, 1, 1, 4, 6}, 2},
		{[]int{1, 5, 6, 14, 15}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDifference(tt.nums); got != tt.want {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
