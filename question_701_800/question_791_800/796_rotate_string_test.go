package question_791_800

import (
	"testing"
)

func Test_rotateString(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want bool
	}{
		{"abcde", "cdeab", true},
		{"ab", "ba", true},
		{"abcde", "abcde", true},
		{"abcde", "abced", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rotateString(tt.A, tt.B); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
