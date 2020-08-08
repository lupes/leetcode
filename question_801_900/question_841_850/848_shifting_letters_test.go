package question_841_850

import (
	"testing"
)

func Test_shiftingLetters(t *testing.T) {
	tests := []struct {
		S      string
		shifts []int
		want   string
	}{
		{"abc", []int{3, 5, 9}, "rpl"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shiftingLetters(tt.S, tt.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
