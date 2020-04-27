package question_1001_1010

import (
	"testing"
)

func Test_minDominoRotations(t *testing.T) {
	tests := []struct {
		A    []int
		B    []int
		want int
	}{
		{[]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}, 2},
		{[]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}, -1},
		{[]int{3, 5, 1, 2, 3, 4, 3, 6, 3, 2}, []int{1, 3, 3, 3, 4, 3, 3, 3, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDominoRotations(tt.A, tt.B); got != tt.want {
				t.Errorf("minDominoRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
