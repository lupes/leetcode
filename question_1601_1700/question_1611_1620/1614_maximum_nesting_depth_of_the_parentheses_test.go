package question_1611_1620

import (
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"1+(2*3)/(2-1)", 1},
		{"1", 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxDepth(tt.s); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
