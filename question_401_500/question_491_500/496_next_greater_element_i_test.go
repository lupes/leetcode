package question_491_500

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"test", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"test", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
