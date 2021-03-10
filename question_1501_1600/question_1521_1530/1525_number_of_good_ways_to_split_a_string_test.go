package question_1521_1530

import (
	"testing"
)

func Test_numSplits(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aacaba", 2},
		{"abcd", 1},
		{"aaaaa", 4},
		{"acbadbaada", 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numSplits(tt.s); got != tt.want {
				t.Errorf("numSplits() = %v, want %v", got, tt.want)
			}
		})
	}
}
