package question_71_80

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "aa", ""},
		{"ADOBECODEBANCA", "AABC", "BANCA"},
		{"abcabdebac", "cda", "cabd"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
