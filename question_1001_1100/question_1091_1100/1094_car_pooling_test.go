package question_1091_1100

import (
	"testing"
)

func Test_carPooling(t *testing.T) {
	tests := []struct {
		trips    [][]int
		capacity int
		want     bool
	}{
		{[][]int{{2, 1, 5}, {3, 3, 7}}, 4, false},
		{[][]int{{2, 1, 5}, {3, 3, 7}}, 5, true},
		{[][]int{{2, 1, 5}, {3, 5, 7}}, 3, true},
		{[][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := carPooling(tt.trips, tt.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
