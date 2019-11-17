package question_41_50

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{"test#0", "aa", "a", false},
		{"test#1", "aa", "*", true},
		{"test#2", "cb", "?a", false},
		{"test#3", "adceb", "*a*b", true},
		{"test#4", "acdcb", "a*c?b", false},
		{"test#5", "aab", "c*a*b", false},
		{"test#6", "", "*", true},
		{"test#7", "a", "a*", true},
		{"test#8", "aa", "aa", true},
		{"test#9", "", "", true},
		{"test#10", "a", "", false},
		{"test#11", "ho", "**ho", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
