package question_81_90

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"test", []int{1}, [][]int{{}, {1}}},
		{"test", []int{1, 1}, [][]int{{}, {1}, {1, 1}}},
		{"test", []int{1, 1, 1}, [][]int{{}, {1}, {1, 1}, {1, 1, 1}}},
		{"test", []int{1, 1, 1, 2}, [][]int{{}, {1}, {2}, {1, 1}, {1, 2}, {1, 1, 1}, {1, 1, 2}, {1, 1, 1, 2}}},
		{"test", []int{1, 2, 2}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
