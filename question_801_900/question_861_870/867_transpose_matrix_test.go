package question_861_870

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {

	tests := []struct {
		A    [][]int
		want [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := transpose(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
