package question_651_660

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{[]int{1}, 1, 2, []int{1}},
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 4, 5, 6}, 4, 3, []int{1, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 6}, 4, 5, []int{2, 3, 4, 6}},
		{[]int{1, 1, 1, 10, 10, 10}, 1, 9, []int{10}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findClosestElements(tt.arr, tt.k, tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
