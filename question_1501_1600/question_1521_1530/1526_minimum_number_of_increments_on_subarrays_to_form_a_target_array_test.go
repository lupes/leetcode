package question_1521_1530

import (
	"testing"
)

func Test_minNumberOperations(t *testing.T) {
	tests := []struct {
		target []int
		want   int
	}{
		{[]int{1, 2, 3, 2, 1}, 3},
		{[]int{3, 1, 1, 2}, 4},
		{[]int{3, 1, 5, 4, 2}, 7},
		{[]int{1, 1, 1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minNumberOperations(tt.target); got != tt.want {
				t.Errorf("minNumberOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
