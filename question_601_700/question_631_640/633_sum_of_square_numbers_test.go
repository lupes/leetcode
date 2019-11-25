package question_631_640

import (
	"testing"
)

func Test_judgeSquareSum(t *testing.T) {
	tests := []struct {
		c    int
		want bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := judgeSquareSum(tt.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
