package question_451_460

import (
	"testing"
)

func Test_minMoves(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 1}, 3},
		{[]int{1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minMoves(tt.nums); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
