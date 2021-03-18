package question_1531_1540

import (
	"testing"
)

func Test_canConvertString(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		k    int
		want bool
	}{
		{"input", "ouput", 9, true},
		{"abc", "bcd", 10, false},
		{"aab", "bbb", 27, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canConvertString(tt.s, tt.t, tt.k); got != tt.want {
				t.Errorf("canConvertString() = %v, want %v", got, tt.want)
			}
		})
	}
}
