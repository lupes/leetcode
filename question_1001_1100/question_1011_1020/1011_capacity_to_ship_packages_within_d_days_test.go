package question_1011_1020

import (
	"testing"
)

func Test_shipWithinDaysHelper(t *testing.T) {
	tests := []struct {
		weights []int
		wight   int
		want    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15, 5},
		{[]int{1, 2, 3, 1, 1}, 3, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shipWithinDaysHelper(tt.weights, tt.wight); got != tt.want {
				t.Errorf("shipWithinDaysHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shipWithinDays(t *testing.T) {
	tests := []struct {
		weights []int
		D       int
		want    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
		{[]int{1, 2, 3, 1, 1}, 4, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shipWithinDays(tt.weights, tt.D); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
