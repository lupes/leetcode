package offer

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		//{[]int{1, 2, 3}, [][]int{{}}},
		{[]int{0}, [][]int{[]int{}, []int{0}}},
		{[]int{0, 1}, [][]int{[]int{}, []int{0}, []int{1}, []int{0, 1}}},
		{[]int{0, 1, 2}, [][]int{[]int{}, []int{0}, []int{1}, []int{0, 1}, []int{2}, []int{0, 2}, []int{1, 2}, []int{0, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := subsets(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
