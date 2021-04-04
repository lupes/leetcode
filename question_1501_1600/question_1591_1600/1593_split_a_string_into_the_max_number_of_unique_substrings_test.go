package question_1591_1600

import (
	"testing"
)

func Test_maxUniqueSplit(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"ababccc", 5},
		{"abcdefghijklmnopq", 17},
		{"abcdefghijklmnopqrstuvw", 23},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxUniqueSplit(tt.s); got != tt.want {
				t.Errorf("maxUniqueSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
