package question_81_90

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{"test#0", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"test#1", []int{1, 2, 3, 0, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6, 0}},
		{"test#2", []int{0}, 0, []int{1}, 1, []int{1}},
		{"test#3", []int{1}, 1, []int{}, 0, []int{1}},
		{"test#4", []int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("got = %+v want= %+v", tt.nums1, tt.want)
			}
		})
	}
}
