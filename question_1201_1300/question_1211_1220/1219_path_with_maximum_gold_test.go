package question_1211_1220

import (
	"testing"
)

func Test_getMaximumGold(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}, 24},
		{[][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}, 28},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getMaximumGold(tt.grid); got != tt.want {
				t.Errorf("getMaximumGold() = %v, want %v", got, tt.want)
			}
		})
	}
}
