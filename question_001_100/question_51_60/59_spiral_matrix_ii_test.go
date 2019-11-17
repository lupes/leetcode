package question_51_60

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{"test#0", 1, [][]int{{1}}},
		{"test#1", 2, [][]int{{1, 2}, {4, 3}}},
		{"test#2", 3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{"test#3", 4, [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
