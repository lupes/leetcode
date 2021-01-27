package question_1411_1420

import (
	"testing"
)

func Test_reformat(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"a0b1c2", "0a1b2c"},
		{"leetcode", ""},
		{"1229857369", ""},
		{"covid2019", "c2o0v1i9d"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reformat(tt.s); got != tt.want {
				t.Errorf("reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}
