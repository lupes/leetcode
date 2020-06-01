package question_1151_1160

import (
	"testing"
)

func Test_countCharacters(t *testing.T) {
	tests := []struct {
		words []string
		chars string
		want  int
	}{
		{[]string{"cat", "bt", "hat", "tree"}, "atach", 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countCharacters(tt.words, tt.chars); got != tt.want {
				t.Errorf("countCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
