package question_1591_1600

import (
	"testing"
)

func Test_minOperationsMaxProfit(t *testing.T) {
	tests := []struct {
		customers    []int
		boardingCost int
		runningCost  int
		want         int
	}{
		{[]int{8, 3}, 5, 6, 3},
		{[]int{10, 9, 6}, 6, 4, 7},
		{[]int{3, 4, 0, 5, 1}, 1, 92, -1},
		{[]int{10, 10, 6, 4, 7}, 3, 8, 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minOperationsMaxProfit(tt.customers, tt.boardingCost, tt.runningCost); got != tt.want {
				t.Errorf("minOperationsMaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
