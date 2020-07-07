package question_401_410

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"aaa", 3},
		{"aaA", 3},
		{"aBA", 1},
		{"aAAa", 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
