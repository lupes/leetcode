package question_1991_2000

import (
	"reflect"
	"testing"
)

func Test_findFarmland(t *testing.T) {
	tests := []struct {
		land [][]int
		want [][]int
	}{
		{[][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}, [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}}},
		{[][]int{{1, 1}, {1, 1}}, [][]int{{0, 0, 1, 1}}},
		{[][]int{{0}}, nil},
		{[][]int{{0}, {0}, {1}, {1}, {1}, {0}, {1}, {1}}, [][]int{{2, 0, 4, 0}, {6, 0, 7, 0}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findFarmland(tt.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFarmland() = %v, want %v", got, tt.want)
			}
		})
	}
}
