package question_1441_1450

import (
	"testing"
)

func Test_maxPower(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 2},
		{"abbcccddddeeeeedcba", 5},
		{"triplepillooooow", 5},
		{"hooraaaaaaaaaaay", 11},
		{"tourist", 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxPower(tt.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
