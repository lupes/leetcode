package question_161_170

import (
	"testing"
)

func Test_fractionToDecimal(t *testing.T) {
	tests := []struct {
		numerator   int
		denominator int
		want        string
	}{
		{1, 0, ""},
		{0, 3, "0"},
		{2, 1, "2"},
		{-50, 8, "-6.25"},
		{200, 5, "40"},
		{200, 500, "0.4"},
		{2, -1, "-2"},
		{-2, -1, "2"},
		{1, 2, "0.5"},
		{1, 3, "0.(3)"},
		{2, 3, "0.(6)"},
		{2, 7, "0.(285714)"},
		{7, -12, "-0.58(3)"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := fractionToDecimal(tt.numerator, tt.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
