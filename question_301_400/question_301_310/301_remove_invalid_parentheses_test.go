package question_301_310

import (
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"()())()", []string{"(())()", "()()()"}},
		{"(a)())()", []string{"(a())()", "(a)()()"}},
		{")(", []string{""}},
		{"n", []string{"n"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeInvalidParentheses(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %#v, want %v", got, tt.want)
			}
		})
	}
}
