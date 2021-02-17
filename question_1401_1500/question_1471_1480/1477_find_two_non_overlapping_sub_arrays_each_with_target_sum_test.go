package question_1471_1480

import (
	"testing"
)

func Test_minSumOfLengths(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		//{[]int{3, 2, 2, 4, 3}, 3, 2},
		//{[]int{7, 3, 4, 7}, 7, 2},
		//{[]int{4, 3, 2, 6, 2, 3, 4}, 6, -1},
		//{[]int{5, 5, 4, 4, 5}, 3, -1},
		//{[]int{3, 1, 1, 1, 5, 1, 2, 1}, 3, 3},
		{[]int{1, 1, 1, 2, 2, 2, 4, 4}, 6, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minSumOfLengths(tt.arr, tt.target); got != tt.want {
				t.Errorf("minSumOfLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}
