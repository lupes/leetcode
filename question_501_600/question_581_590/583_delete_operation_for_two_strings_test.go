package question_581_590

import (
	"testing"
)

func Test_minDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		{"sea", "eat", 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
