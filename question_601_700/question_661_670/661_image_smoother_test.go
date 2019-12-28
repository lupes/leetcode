package question_661_670

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	tests := []struct {
		M    [][]int
		want [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{[][]int{
			{2, 3, 4},
			{5, 6, 7},
			{8, 9, 10},
			{11, 12, 13},
			{14, 15, 16}},
			[][]int{{4, 4, 5}, {5, 6, 6}, {8, 9, 9}, {11, 12, 12}, {13, 13, 14}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := imageSmoother(tt.M); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
