package question_1331_1340

import (
	"testing"
)

func Test_numberOfSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabc", 10},
		{"aaacb", 3},
		{"abc", 1},
		{"abca", 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberOfSubstrings(tt.s); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
