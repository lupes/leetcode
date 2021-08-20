package question_81_90

import (
	"testing"
)

func Test_isScramble(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		//{"great", "rgeat", true},
		//{"great", "eatgr", true},
		{"rstuvwxyzabcdefghigklmnopq", "abcdefghigklmnopqrstuvwxyz", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"aa", "cc", false},
		{"abc", "cba", true},
		{"abcd", "cbda", true},
		{"abcd", "cdba", true},
		{"ccabcbabcbabbbbcbb", "bbbbabccccbbbabcba", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isScramble(tt.s1, tt.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
