package question_81_90

import (
	"testing"
)

func Test_maximalRectangle(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{[][]byte{{'1'}}, 1},
		{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maximalRectangle(tt.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
