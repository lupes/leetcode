package question_341_350

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"test", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"test", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.nums1, tt.nums2)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
