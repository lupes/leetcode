package question_1151_1160

import (
	"testing"
)

func Test_numRollsToTargetHelper(t *testing.T) {
	tests := []struct {
		d      int
		f      int
		target int
		want   int
	}{
		//{1, 6, 3, 1},
		//{2, 6, 7, 6},
		//{1, 2, 3, 0},
		//{10, 10, 30, 0},
		{30, 30, 500, 0},
		//{5, 30, 100, 263121},
		//{7, 30, 200, 263121},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numRollsToTarget(tt.d, tt.f, tt.target); got != tt.want {
				t.Errorf("numRollsToTargetHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
