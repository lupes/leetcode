package question_41_50

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	matrix1 := [][]int{{1}}
	want1 := [][]int{{1}}

	matrix2 := [][]int{{1, 2}, {3, 4}}
	want2 := [][]int{{3, 1}, {4, 2}}

	matrix3 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want3 := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	matrix4 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	want4 := [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}

	matrix5 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	want5 := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}

	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{"test#1", matrix1, want1},
		{"test#2", matrix2, want2},
		{"test#3", matrix3, want3},
		{"test#4", matrix4, want4},
		{"test#5", matrix5, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("%+v != %+v", tt.matrix, tt.want)
			}
		})
	}
}
