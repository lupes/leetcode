package question_491_500

import (
	"reflect"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"test", []int{4, 6}, [][]int{{4, 6}}},
		{"test", []int{4, 6, 7, 7}, [][]int{{4, 6}, {6, 7}, {4, 6, 7}, {7, 7}, {4, 7, 7}, {6, 7, 7}, {4, 6, 7, 7}, {4, 7}}},
		{"test", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubsequences(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
