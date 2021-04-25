package question_1641_1650

import (
	"testing"
)

func Test_minDeletions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aab", 0},
		{"aaabbbcc", 2},
		{"ceabaacb", 2},
		{"abbcccddd", 3},
		{"bbcebab", 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDeletions(tt.s); got != tt.want {
				t.Errorf("minDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
