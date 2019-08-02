package question_281_290

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"test", []int{0, 0, 1, 0, 3, 12, 0}, []int{1, 3, 12, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("moveZeroes got=%v want=%v\n", tt.nums, tt.want)
			}
		})
	}
}
