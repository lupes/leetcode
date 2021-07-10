package question_671_680

import (
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"ab", true},
		{"a", true},
		{"aba", true},
		{"abca", true},
		{"bddb", true},
		{"abbca", true},
		{"abbdca", false},
		{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", true},
		{"acxcybycxcxa", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := validPalindrome(tt.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
