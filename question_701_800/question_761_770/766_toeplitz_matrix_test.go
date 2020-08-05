package question_761_770

import (
	"testing"
)

func Test_isToeplitzMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   bool
	}{
		{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true},
		{[][]int{{1, 2}, {2, 2}}, false},
	}
	for _, tt := range tests {
		t.Run("teat", func(t *testing.T) {
			if got := isToeplitzMatrix(tt.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
