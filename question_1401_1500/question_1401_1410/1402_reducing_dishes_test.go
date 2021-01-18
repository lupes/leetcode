package question_1401_1410

import (
	"testing"
)

func Test_maxSatisfaction(t *testing.T) {
	tests := []struct {
		satisfaction []int
		want         int
	}{
		{[]int{-1, -8, 0, 5, -9}, 14},
		{[]int{4, 3, 2}, 20},
		{[]int{-1, -4, -5}, 0},
		{[]int{-2, 5, -1, 0, 3, -3}, 35},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxSatisfaction(tt.satisfaction); got != tt.want {
				t.Errorf("maxSatisfaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
