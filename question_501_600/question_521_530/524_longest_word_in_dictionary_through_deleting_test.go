package question_521_530

import (
	"testing"
)

func Test_findLongestWord(t *testing.T) {
	tests := []struct {
		s    string
		d    []string
		want string
	}{
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
		{"abce", []string{"abe", "abc"}, "abc"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findLongestWord(tt.s, tt.d); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
