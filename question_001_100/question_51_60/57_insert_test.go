package question_51_60

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{"test#0", [][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{"test#1", [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{"test#2", [][]int{{1, 3}, {6, 9}}, []int{10, 11}, [][]int{{1, 3}, {6, 9}, {10, 11}}},
		{"test#3", [][]int{{1, 3}, {6, 9}}, []int{-2, 0}, [][]int{{-2, 0}, {1, 3}, {6, 9}}},
		{"test#4", [][]int{{1, 5}}, []int{2, 7}, [][]int{{1, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.intervals, tt.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
