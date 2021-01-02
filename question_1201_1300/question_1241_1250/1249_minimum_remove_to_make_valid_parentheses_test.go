package question_1241_1250

import (
	"testing"
)

func Test_minRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
		{"(a(b(c)d)", "a(b(c)d)"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
