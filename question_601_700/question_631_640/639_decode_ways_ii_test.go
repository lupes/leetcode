package question_631_640

import (
	"testing"
)

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"*", 9},
		{"0", 0},
		{"**", 96},
		{"1*", 18},
		{"2*", 15},
		{"*1*1*0", 404},
		{"*10*1", 99},
		{"104", 1},
		{"6*66*6", 121},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
