package question_1471_1480

import (
	"reflect"
	"testing"
)

func Test_getStrongest(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{5, 1}},
		{[]int{1, 1, 3, 5, 5}, 2, []int{5, 5}},
		{[]int{6, 7, 11, 7, 6, 8}, 5, []int{11, 8, 6, 6, 7}},
		{[]int{6, -3, 7, 2, 11}, 3, []int{-3, 11, 2}},
		{[]int{-7, 22, 17, 3}, 2, []int{22, 17}},
		{[]int{-2, -4, -6, -8, -9, -7, -5, -3, -1}, 3, []int{-1, -9, -2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getStrongest(tt.arr, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStrongest() = %v, want %v", got, tt.want)
			}
		})
	}
}
