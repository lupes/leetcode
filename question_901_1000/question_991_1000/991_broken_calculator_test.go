package question_991_1000

import (
	"testing"
)

func Test_brokenCalc(t *testing.T) {
	tests := []struct {
		X int
		Y int
		W int
	}{
		{2, 3, 2},
		{5, 8, 2},
		{3, 10, 3},
		{1024, 1, 1023},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := brokenCalc(tt.X, tt.Y); got != tt.W {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.W)
			}
		})
	}
}
