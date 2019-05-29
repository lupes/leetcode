package question_161_170

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"test", []int{-2, -1}, -3, []int{1, 2}},
		{"test", []int{-1, 0}, -1, []int{1, 2}},
		{"test", []int{0, 1}, 1, []int{1, 2}},
		{"test", []int{1, 2, 3}, 3, []int{1, 2}},
		{"test", []int{1, 2}, 3, []int{1, 2}},
		{"test", []int{1, 2, 3}, 3, []int{1, 2}},
		{"test", []int{1, 2, 3, 4}, 3, []int{1, 2}},
		{"test", []int{1, 2, 3, 4}, 6, []int{2, 4}},
		{"test", []int{1, 2, 3, 5, 11, 12}, 16, []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
