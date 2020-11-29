package question_211_220

import (
	"testing"
)

func Test_shortestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"ab", "bab"},
		{"aacecaaa", "aaacecaaa"},
		{"aa", "aa"},
		{"abc", "cbabc"},
		{"abcd", "dcbabcd"},
		{"aaa", "aaa"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shortestPalindrome(tt.s); got != tt.want {
				t.Errorf("shortestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
