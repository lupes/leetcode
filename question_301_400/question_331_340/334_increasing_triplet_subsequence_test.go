package question_331_340

import (
	"testing"
)

func Test_increasingTriplet(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := increasingTriplet(tt.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
