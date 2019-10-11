package question_491_500

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{"test", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{"test", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 3, 5, 4, 6}},
		{"test", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 4, 5, 3, 6}},
		{"test", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
