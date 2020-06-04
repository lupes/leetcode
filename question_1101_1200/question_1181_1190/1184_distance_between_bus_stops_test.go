package question_1181_1190

import (
	"testing"
)

func Test_distanceBetweenBusStops(t *testing.T) {
	tests := []struct {
		distance    []int
		start       int
		destination int
		want        int
	}{
		{[]int{1, 2, 3, 4}, 0, 1, 1},
		{[]int{1, 2, 3, 4}, 0, 2, 3},
		{[]int{1, 2, 3, 4}, 0, 3, 4},
		{[]int{1, 2, 3, 4}, 3, 2, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := distanceBetweenBusStops(tt.distance, tt.start, tt.destination); got != tt.want {
				t.Errorf("distanceBetweenBusStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
