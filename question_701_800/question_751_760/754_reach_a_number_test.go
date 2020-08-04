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
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reachNumber(tt.target); got != tt.want {
				t.Errorf("reachNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
