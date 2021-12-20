package question_471_480

import (
	"testing"
)

func Test_findRadius(t *testing.T) {
	tests := []struct {
		houses  []int
		heaters []int
		want    int
	}{
		{[]int{1, 2, 3}, []int{2}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
		{[]int{1, 5}, []int{2}, 3},
		{[]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findRadius(tt.houses, tt.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
