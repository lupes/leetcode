package question_1671_1680

import (
	"reflect"
	"testing"
)

func Test_mostCompetitive(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{3, 5, 2, 6}, 2, []int{2, 6}},
		{[]int{3, 5, 2, 6}, 4, []int{3, 5, 2, 6}},
		{[]int{3, 5, 2, 6}, 3, []int{3, 2, 6}},
		{[]int{1, 3, 4, 1, 3, 5}, 3, []int{1, 1, 3}},
		{[]int{2, 4, 3, 3, 5, 4, 9, 6}, 4, []int{2, 3, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mostCompetitive(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostCompetitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
