package question_1661_1670

import (
	"testing"
)

func Test_maxRepeating(t *testing.T) {
	tests := []struct {
		sequence string
		word     string
		want     int
	}{
		{"ababc", "ab", 2},
		{"a", "a", 1},
		{"ababc", "ba", 1},
		{"ababc", "ac", 0},
		{"aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba", 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxRepeating(tt.sequence, tt.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
