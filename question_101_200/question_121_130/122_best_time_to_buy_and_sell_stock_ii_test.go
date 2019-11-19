package question_121_130

import (
	"testing"
)

func Test_maxProfit2(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxProfit2(tt.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
