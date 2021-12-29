package question_1991_2000

import (
	"testing"
)

func Test_countQuadruplets(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 6}, 1},
		{[]int{3, 3, 6, 4, 5}, 0},
		{[]int{1, 1, 1, 3, 5}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countQuadruplets(tt.nums); got != tt.want {
				t.Errorf("countQuadruplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
