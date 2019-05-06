package question_71_80

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test#0", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("got %+v want: %+v", tt.nums, tt.want)
			}
		})
	}
}
