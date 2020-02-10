package question_801_810

import (
	"testing"
)

func Test_expressiveWords(t *testing.T) {
	tests := []struct {
		S     string
		words []string
		want  int
	}{
		{"abcd", []string{"abc"}, 0},
		{"heeellooo", []string{"hello", "hi", "helo"}, 1},
		{"dddiiiinnssssssoooo",
			[]string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := expressiveWords(tt.S, tt.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
