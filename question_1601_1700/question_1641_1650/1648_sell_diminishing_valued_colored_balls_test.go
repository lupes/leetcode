package question_1641_1650

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		inventory []int
		orders    int
		want      int
	}{
		{[]int{2, 5}, 4, 14},
		{[]int{3, 5}, 6, 19},
		{[]int{2, 8, 4, 10, 6}, 20, 110},
		{[]int{1000000000}, 1000000000, 21},
		{[]int{73, 98, 35, 100}, 178, 11120},
		{[]int{497978859, 167261111, 483575207, 591815159}, 836556809, 373219333},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxProfit(tt.inventory, tt.orders); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
