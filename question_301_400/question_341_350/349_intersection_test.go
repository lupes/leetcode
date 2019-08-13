package question_341_350

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"test", []int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{"test", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
