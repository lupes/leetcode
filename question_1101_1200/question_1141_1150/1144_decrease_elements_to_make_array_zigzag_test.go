package question_1141_1150

import (
	"testing"
)

func Test_movesToMakeZigzag(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{1, 2, 3}, 2},
		{[]int{9, 6, 1, 6, 2}, 4},
		{[]int{2, 7, 10, 9, 8, 9}, 4},
		{[]int{10, 4, 4, 10, 10, 6, 2, 3}, 13},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := movesToMakeZigzag(tt.nums); got != tt.want {
				t.Errorf("movesToMakeZigzag() = %v, want %v", got, tt.want)
			}
		})
	}
}
