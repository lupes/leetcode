package question_1611_1620

import (
	"reflect"
	"testing"
)

func Test_bestCoordinate(t *testing.T) {
	tests := []struct {
		towers [][]int
		radius int
		want   []int
	}{
		{[][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2, []int{2, 1}},
		{[][]int{{23, 11, 21}}, 9, []int{23, 11}},
		{[][]int{{1, 2, 13}, {2, 1, 7}, {0, 1, 9}}, 2, []int{1, 2}},
		{[][]int{{2, 1, 9}, {0, 1, 9}}, 2, []int{0, 1}},
		{[][]int{{42, 0, 0}}, 7, []int{0, 0}},
		{[][]int{{0, 1, 2}, {2, 1, 2}, {1, 0, 2}, {1, 2, 2}}, 1, []int{1, 1}},
		{[][]int{{31, 13, 33}, {24, 45, 38}, {28, 32, 23}, {7, 23, 22}, {41, 50, 33}, {47, 21, 3}, {3, 33, 39}, {11, 38, 5}, {26, 20, 28}, {48, 39, 16}, {34, 29, 25}}, 21, []int{24, 45}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := bestCoordinate(tt.towers, tt.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}
