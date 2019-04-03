package question_0011_0020

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test#0", "", 0},
		{"test#1", ")(", 0},
		{"test#2", "()", 2},
		{"test#3", "(()", 2},
		{"test#4", ")()())", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
