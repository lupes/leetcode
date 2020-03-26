package question_901_910

import (
	"testing"
)

func Test_totalFruit(t *testing.T) {
	tests := []struct {
		tree []int
		want int
	}{
		{[]int{1, 2, 1, 2, 3}, 4},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{0, 1, 2, 2, 3, 2}, 4},
		{[]int{1, 2, 3, 2, 2}, 4},
		{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := totalFruit(tt.tree); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
