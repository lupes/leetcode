package question_1451_1460

import (
	"testing"
)

func Test_maxVowels(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		{"rhythms", 4, 0},
		{"tryhard", 4, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxVowels(tt.s, tt.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
