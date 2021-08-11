package question_221_230

import (
	"testing"
)

func Test_computeArea(t *testing.T) {
	tests := []struct {
		ax1  int
		ay1  int
		ax2  int
		ay2  int
		bx1  int
		by1  int
		bx2  int
		by2  int
		want int
	}{
		{-3, 0, 3, 4, 0, -1, 9, 2, 3},
		{-2, -2, 2, 2, -2, -2, 2, 2, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := computeArea(tt.ax1, tt.ay1, tt.ax2, tt.ay2, tt.bx1, tt.by1, tt.bx2, tt.by2); got != tt.want {
				t.Errorf("computeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
