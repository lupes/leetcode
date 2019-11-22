package question_571_580

import (
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	tests := []struct {
		candies []int
		want    int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{1, 1, 2, 3, 4, 5}, 3},
		{[]int{1, 1, 1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := distributeCandies(tt.candies); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
