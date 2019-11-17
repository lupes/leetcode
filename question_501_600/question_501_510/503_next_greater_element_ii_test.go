package question_501_510

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test", []int{4, 1, 2}, []int{-1, 2, 4}},
		{"test", []int{2, 4}, []int{4, -1}},
		{"test", []int{1, 2, 1}, []int{2, -1, 2}},
		{"test", []int{2, 1, 3, 6, 5, 4, 2, 1, 0}, []int{3, 3, 6, -1, 6, 6, 3, 2, 2}},
		{"test", []int{1, 8, -1, -100, -1, 222, 1111111, -111111}, []int{8, 222, 222, -1, 222, 1111111, -1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
