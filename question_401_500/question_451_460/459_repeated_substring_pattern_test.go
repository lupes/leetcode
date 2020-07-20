package question_451_460

import (
	"testing"
)

func Test_repeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"abab", true},
		{"abcab", false},
		{"aba", false},
		{"abcabcabcabc", true},
		{"aa", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
