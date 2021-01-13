package question_1331_1340

import (
	"testing"
)

func Test_minSteps(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want int
	}{
		{"aba", "bab", 1},
		{"leetcode", "practice", 5},
		{"anagram", "mangaar", 0},
		{"xxyyzz", "xxyyzz", 0},
		{"friend", "family", 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minSteps(tt.s, tt.t); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
