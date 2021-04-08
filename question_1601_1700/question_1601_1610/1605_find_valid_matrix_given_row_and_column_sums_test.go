package question_1601_1610

import (
	"reflect"
	"testing"
)

func Test_restoreMatrix(t *testing.T) {
	tests := []struct {
		rowSum []int
		colSum []int
		want   [][]int
	}{
		{
			[]int{3, 8},
			[]int{4, 7},
			[][]int{
				{3, 0},
				{1, 7},
			},
		},
		{
			[]int{5, 7, 10},
			[]int{8, 6, 8},
			[][]int{
				{5, 0, 0},
				{3, 4, 0},
				{0, 2, 8},
			},
		},
		{
			[]int{14, 9},
			[]int{6, 9, 8},
			[][]int{
				{6, 8, 0},
				{0, 1, 8},
			},
		},
		{
			[]int{1, 0},
			[]int{1},
			[][]int{
				{1},
				{0},
			},
		},
		{
			[]int{0},
			[]int{0},
			[][]int{
				{0},
			},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := restoreMatrix(tt.rowSum, tt.colSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
