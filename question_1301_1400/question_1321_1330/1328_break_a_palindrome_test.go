package question_1321_1330

import (
	"testing"
)

func Test_breakPalindrome(t *testing.T) {
	tests := []struct {
		palindrome string
		want       string
	}{
		{"abccba", "aaccba"},
		{"abcdcba", "aacdcba"},
		{"a", ""},
		{"aa", "ab"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := breakPalindrome(tt.palindrome); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
