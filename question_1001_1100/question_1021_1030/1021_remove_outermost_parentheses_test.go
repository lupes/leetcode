package question_1021_1030

import (
	"testing"
)

func Test_removeOuterParentheses(t *testing.T) {
	tests := []struct {
		S    string
		want string
	}{
		{"(()())(())", "()()()"},
		{"(()())(())(()(()))", "()()()()(())"},
		{"()()", ""},
		{"()", ""},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeOuterParentheses(tt.S); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
