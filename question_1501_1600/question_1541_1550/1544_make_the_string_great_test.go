package question_1541_1550

import (
	"testing"
)

func Test_makeGood(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"leEeetcode", "leetcode"},
		{"abBAcC", ""},
		{"s", "s"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := makeGood(tt.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
