package question_871_880

import (
	"testing"
)

func Test_minEatingSpeed(t *testing.T) {
	tests := []struct {
		piles []int
		H     int
		want  int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minEatingSpeed(tt.piles, tt.H); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
