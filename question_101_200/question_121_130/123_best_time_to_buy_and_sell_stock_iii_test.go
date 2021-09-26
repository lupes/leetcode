package question_121_130

import (
	"testing"
)

func Test_maxProfit3(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{[]int{1, 2, 3, 4, 5}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxProfit3(tt.prices); got != tt.want {
				t.Errorf("maxProfit3() = %v, want %v", got, tt.want)
			}
		})
	}
}
