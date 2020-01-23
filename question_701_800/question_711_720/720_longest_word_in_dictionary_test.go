package question_711_720

import (
	"testing"
)

func Test_longestWord(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{[]string{"w", "wo", "wor", "worl", "world"}, "world"},
		{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"},
		{[]string{"banana", "app", "appl", "ap", "apply", "apple"}, ""},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := longestWord(tt.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
