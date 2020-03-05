package question_861_870

import (
	"testing"
)

func Test_binaryGap(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 0},
		{5, 2},
		{6, 1},
		{7, 1},
		{8, 0},
		{9, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := binaryGap(tt.N); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
