package question_1201_1210

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"abcd", 2, "abcd"},
		{"deeedbbcccbdaa", 3, "aa"},
		{"pbbcggttciiippooaais", 2, "ps"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeDuplicates(tt.s, tt.k); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
