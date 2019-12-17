package question_271_280

import (
	"testing"
)

func Test_numSquares(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 1},
		{5, 2},
		{6, 3},
		{7, 4},
		{8, 2},
		{9, 1},
		{10, 2},
		{11, 3},
		{12, 3},
		{13, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numSquares(tt.n); got != tt.want {
				t.Errorf("numSquares(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
