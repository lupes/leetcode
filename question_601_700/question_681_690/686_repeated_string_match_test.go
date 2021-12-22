package question_681_690

import (
	"testing"
)

func Test_repeatedStringMatch(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want int
	}{
		{"abcd", "cdabcdab", 3},
		{"a", "aa", 2},
		{"a", "a", 1},
		{"abc", "wxyz", -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := repeatedStringMatch(tt.A, tt.B); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
