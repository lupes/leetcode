package question_1131_1140

import (
	"testing"
)

func Test_alphabetBoardPath(t *testing.T) {
	tests := []struct {
		target string
		want   string
	}{
		{"leet", "DDR!UURRR!!DDD!"},
		{"code", "RR!DDRR!UUL!R!"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := alphabetBoardPath(tt.target); got != tt.want {
				t.Errorf("alphabetBoardPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
