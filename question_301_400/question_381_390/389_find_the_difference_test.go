package question_381_390

import (
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want byte
	}{
		{"abcd", "abcde", 'e'},
		{"abcd", "abcdd", 'd'},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findTheDifference(tt.s, tt.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
