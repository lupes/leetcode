package question_641_650

import (
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abc", 3},
		{"aa", 3},
		{"aaa", 6},
		{"aaab", 7},
		{"baaab", 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countSubstrings(tt.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
