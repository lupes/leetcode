package question_1411_1420

import (
	"testing"
)

func Test_minStartValue(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-3, 2, -3, 4, 2}, 5},
		{[]int{1, 2}, 1},
		{[]int{1, -2, -3}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minStartValue(tt.nums); got != tt.want {
				t.Errorf("minStartValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
