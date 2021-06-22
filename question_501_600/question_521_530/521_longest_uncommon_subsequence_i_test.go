package question_521_530

import (
	"testing"
)

func Test_findLUSlength(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"aba", "cdc", 3},
		{"aaa", "bbb", 3},
		{"aaa", "aaa", -1},
		{"abcd", "abc", 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findLUSlength(tt.a, tt.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
