package question_1051_1060

import (
	"testing"
)

func Test_maxSatisfied(t *testing.T) {
	tests := []struct {
		customers []int
		grumpy    []int
		X         int
		want      int
	}{
		{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3, 16},
		{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 5, 18},
		{[]int{7, 0, 1, 2, 1, 1, 7, 5}, []int{1, 1, 0, 1, 0, 1, 0, 1}, 1, 16},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxSatisfied(tt.customers, tt.grumpy, tt.X); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
