package question_361_370

import (
	"testing"
)

func Test_canMeasureWater(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		z    int
		want bool
	}{
		{3, 5, 4, true},
		{2, 6, 5, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canMeasureWater(tt.x, tt.y, tt.z); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
