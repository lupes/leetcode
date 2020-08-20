package question_951_960

import (
	"testing"
)

func Test_isAlienSorted(t *testing.T) {
	tests := []struct {
		words []string
		order string
		want  bool
	}{
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
		{[]string{"app", "apple"}, "abcdefghijklmnopqrstuvwxyz", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isAlienSorted(tt.words, tt.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
