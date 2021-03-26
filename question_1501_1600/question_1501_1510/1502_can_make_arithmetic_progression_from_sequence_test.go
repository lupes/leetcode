package question_1501_1510

import (
	"testing"
)

func Test_canMakeArithmeticProgression(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{3, 5, 1}, true},
		{[]int{1, 2, 4}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canMakeArithmeticProgression(tt.arr); got != tt.want {
				t.Errorf("canMakeArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}
