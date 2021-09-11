package question_591_600

import (
	"testing"
)

func Test_findIntegers(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findIntegers(tt.n); got != tt.want {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
