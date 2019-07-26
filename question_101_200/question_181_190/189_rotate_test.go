package question_181_190

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"test", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"test", []int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{"test", []int{1, 2, 3, 4, 5, 6, 7}, 1, []int{7, 1, 2, 3, 4, 5, 6}},
		{"test", []int{1, 2, 3, 4, 5, 6, 7}, 2, []int{6, 7, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("got=%v want=%v\n", tt.nums, tt.want)
			}
		})
	}
}
