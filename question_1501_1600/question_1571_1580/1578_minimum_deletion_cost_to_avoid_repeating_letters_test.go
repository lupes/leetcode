package question_1571_1580

import (
	"testing"
)

func Test_minCost(t *testing.T) {
	tests := []struct {
		s    string
		cost []int
		want int
	}{
		{"abaac", []int{1, 2, 3, 4, 5}, 3},
		{"abc", []int{1, 2, 3}, 0},
		{"aabaa", []int{1, 2, 3, 4, 1}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minCost(tt.s, tt.cost); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
