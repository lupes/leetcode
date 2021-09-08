package question_951_960

import (
	"testing"
)

func Test_minDeletionSize(t *testing.T) {
	tests := []struct {
		A    []string
		want int
	}{
		{[]string{"ca", "bb", "ac"}, 1},
		{[]string{"xc", "yb", "za"}, 0},
		{[]string{"zyx", "wvu", "tsr"}, 3},
		{[]string{"aa", "ab", "ab"}, 0},
		{[]string{"aac", "abb", "aba"}, 1},
		{[]string{"xga", "xfb", "yfa"}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDeletionSize(tt.A); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
