package question_51_60

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{"test#0", [][]int{{1}}, []int{1}},
		{"test#1", [][]int{{1, 2}}, []int{1, 2}},
		{"test#2", [][]int{{1}, {2}}, []int{1, 2}},
		{"test#3", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
		{"test#4", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}, []int{1, 2, 3, 4, 8, 7, 6, 5}},
		{"test#5", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"test#6", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
