package question_441_450

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test", []int{1, 2, 2}, []int{2}},
		{"test", []int{1, 2, 2, 3}, []int{2}},
		{"test", []int{1, 2, 2, 3}, []int{2}},
		{"test", []int{1, 2, 2, 3, 4}, []int{2}},
		{"test", []int{1, 2, 2, 3, 4, 3}, []int{2, 3}},
		{"test", []int{1, 2, 2, 3, 4, 3, 1}, []int{2, 3, 1}},
		{"test", []int{4, 3, 2, 7, 8, 2, 3, 1}, []int{3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
