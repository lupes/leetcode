package question_141_150

import (
	"testing"
)

func Test_evalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{"test#0", []string{"2", "1", "+", "3", "*"}, 9},
		{"test#1", []string{"4", "13", "5", "/", "+"}, 6},
		{"test#2", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
