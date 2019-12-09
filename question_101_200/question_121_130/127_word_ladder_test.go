package question_121_130

import (
	"testing"
)

func Test_ladderLength(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := ladderLength(tt.beginWord, tt.endWord, tt.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
