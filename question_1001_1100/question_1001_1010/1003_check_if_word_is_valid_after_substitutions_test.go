package question_1001_1010

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		S    string
		want bool
	}{
		{"aabcbc", true},
		{"abcabcababcc", true},
		{"abccba", false},
		{"cababc", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isValid(tt.S); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
