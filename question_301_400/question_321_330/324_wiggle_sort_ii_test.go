package question_321_330

import (
	"reflect"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test", []int{1, 1, 2}, []int{1, 2, 1}},
		{"test", []int{1, 2, 3, 3}, []int{2, 3, 1, 3}},
		{"test", []int{4, 5, 5, 6}, []int{5, 6, 4, 5}},
		{"test", []int{1, 5, 1, 1, 6, 4}, []int{1, 6, 1, 5, 1, 4}},
		{"test", []int{1, 3, 2, 2, 3, 1}, []int{2, 3, 1, 3, 1, 2}},
		{"test", []int{5, 3, 1, 2, 6, 7, 8, 5, 5}, []int{5, 7, 3, 6, 2, 5, 1, 8, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("wiggleSort() = %#v, want %#v", tt.nums, tt.want)
			}
		})
	}
}
