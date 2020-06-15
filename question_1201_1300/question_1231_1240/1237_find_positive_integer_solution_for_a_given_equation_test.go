package question_1231_1240

import (
	"reflect"
	"testing"
)

func Test_findSolution(t *testing.T) {
	tests := []struct {
		fun  func(int, int) int
		z    int
		want [][]int
	}{
		{
			func(i int, j int) int {
				return i + j
			},
			5,
			[][]int{{1, 4}, {2, 3}, {3, 2}, {4, 1}}}, {
			func(i int, j int) int {
				return i * j
			},
			5,
			[][]int{{1, 5}, {5, 1}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findSolution(tt.fun, tt.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
