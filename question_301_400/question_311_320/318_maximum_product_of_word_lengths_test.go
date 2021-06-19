package question_311_320

import (
	"testing"
)

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}, 16},
		{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{[]string{"a", "aa", "aaa", "aaaa"}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxProduct(tt.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
