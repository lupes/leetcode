package question_1181_1190

import (
	"testing"
)

func Test_reverseParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"(abcd)", "dcba"},
		{"(u(love)i)", "iloveu"},
		{"(ed(et(oc))el)", "leetcode"},
		{"a(bcdefghijkl(mno)p)q", "apmnolkjihgfedcbq"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reverseParentheses(tt.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
