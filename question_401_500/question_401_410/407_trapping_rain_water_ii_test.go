package question_401_410

import (
	"testing"
)

func Test_trapRainWater(t *testing.T) {
	tests := []struct {
		heightMap [][]int
		want      int
	}{
		{[][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := trapRainWater(tt.heightMap); got != tt.want {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
