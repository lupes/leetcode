package question_521_530

import (
	"testing"
)

func Test_findMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
		{[]int{0, 1, 0, 1, 0}, 4},
		{[]int{0, 1, 0, 1, 1, 1, 0, 0, 0}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMaxLength(tt.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
