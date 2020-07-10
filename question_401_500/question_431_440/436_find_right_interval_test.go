package question_431_440

import (
	"reflect"
	"testing"
)

func Test_findRightIntervalHelper(t *testing.T) {
	tests := []struct {
		keys  []int
		key   int
		want  bool
		want1 int
	}{
		{[]int{1, 3, 4, 7, 8}, 1, true, 1},
		{[]int{1, 3, 4, 7, 8}, 2, true, 3},
		{[]int{1, 3, 4, 7, 8}, 4, true, 4},
		{[]int{1, 3, 4, 7, 8}, 5, true, 7},
		{[]int{1, 3, 4, 7, 8}, 8, true, 8},
		{[]int{1, 3, 4, 7, 8}, 9, false, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got, got1 := findRightIntervalHelper(tt.keys, tt.key)
			if got != tt.want {
				t.Errorf("findRightIntervalHelper() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findRightIntervalHelper() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findRightInterval(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      []int
	}{
		{[][]int{{1, 2}}, []int{-1}},
		{[][]int{{3, 4}, {2, 3}, {1, 2}}, []int{-1, 0, 1}},
		{[][]int{{1, 4}, {2, 3}, {3, 4}}, []int{-1, 2, -1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findRightInterval(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRightInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
