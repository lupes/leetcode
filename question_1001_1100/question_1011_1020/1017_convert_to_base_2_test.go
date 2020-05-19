package question_1011_1020

import (
	"testing"
)

func Test_baseNeg2(t *testing.T) {
	tests := []struct {
		N    int
		want string
	}{
		{2, "110"},
		{3, "111"},
		{4, "100"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := baseNeg2(tt.N); got != tt.want {
				t.Errorf("baseNeg2() = %v, want %v", got, tt.want)
			}
		})
	}
}
