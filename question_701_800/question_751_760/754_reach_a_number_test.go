package question_751_760

import (
	"testing"
)

func Test_reachNumber(t *testing.T) {
	tests := []struct {
		target int
		want   int
	}{
		{1, 1},
		{2, 3},
		{3, 2},
		{100, 15},
		{101, 14},
		{102, 15},
		{1002, 47},
		{10012302, 4475},
		{999999997, 44721},
		{999999998, 44723},
		{999999999, 44721},
		{99999999999999, 14142137},
		{9999999999999999, 141421357},
		{999999999999999999, 1414213562},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reachNumber(tt.target); got != tt.want {
				t.Errorf("reachNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
