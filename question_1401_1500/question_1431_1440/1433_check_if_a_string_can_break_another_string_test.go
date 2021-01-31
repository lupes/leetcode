package question_1431_1440

import (
	"testing"
)

func Test_checkIfCanBreak(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"abc", "xya", true},
		{"abe", "acd", false},
		{"leetcodee", "interview", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkIfCanBreak(tt.s1, tt.s2); got != tt.want {
				t.Errorf("checkIfCanBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
