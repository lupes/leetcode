package question_1661_1670

import (
	"testing"
)

func Test_waysToMakeFair(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 1, 6, 4}, 1},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := waysToMakeFair(tt.nums); got != tt.want {
				t.Errorf("waysToMakeFair() = %v, want %v", got, tt.want)
			}
		})
	}
}
