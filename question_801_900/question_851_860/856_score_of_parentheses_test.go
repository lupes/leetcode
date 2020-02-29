package question_851_860

import (
	"testing"
)

func Test_scoreOfParentheses(t *testing.T) {
	tests := []struct {
		S    string
		want int
	}{
		//{"()", 1},
		//{"()()", 2},
		//{"()()()", 3},
		{"(())", 2},
		{"(()(()))", 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := scoreOfParentheses(tt.S); got != tt.want {
				t.Errorf("scoreOfParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
