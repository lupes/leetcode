package question_1201_1210

import (
	"testing"
)

func Test_equalSubstring(t *testing.T) {
	tests := []struct {
		s       string
		t       string
		maxCost int
		want    int
	}{
		{"abcd", "bcdf", 4, 3},
		{"abc", "ade", 1, 1},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := equalSubstring(tt.s, tt.t, tt.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
