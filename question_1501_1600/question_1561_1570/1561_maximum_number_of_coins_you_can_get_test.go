package question_1561_1570

import (
	"testing"
)

func Test_maxCoins(t *testing.T) {
	tests := []struct {
		piles []int
		want  int
	}{
		{[]int{2, 4, 1, 2, 7, 8}, 9},
		{[]int{1, 2, 3}, 2},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 18},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxCoins(tt.piles); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
