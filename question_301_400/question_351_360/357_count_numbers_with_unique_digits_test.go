package question_351_360

import (
	"testing"
)

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 10},
		{2, 91},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countNumbersWithUniqueDigits(tt.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
