package question_601_610

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0}, 1, true},
		{[]int{0, 1, 0, 0}, 2, false},
		{[]int{0, 0, 1, 0, 0}, 2, true},
		{[]int{0, 0, 1, 0, 0, 0, 0}, 3, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canPlaceFlowers(tt.flowerbed, tt.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
