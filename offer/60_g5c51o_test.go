package offer

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 2}, 1, []int{2}},
		{[]int{1, 2}, 2, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := topKFrequent(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
