package question_41_50

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		//{"test#0", []int{1}, [][]int{{1}}},
		{"test#1", []int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{"test#2", []int{2, 2, 1, 1}, [][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}}},
		//{"test#3", []int{0, 1, 0, 0, 9}, [][]int{{0, 0, 0, 1, 9}, {0, 0, 0, 9, 1}, {0, 0, 1, 0, 9}, {0, 0, 1, 9, 0}, {0, 0, 9, 0, 1}, {0, 0, 9, 1, 0}, {0, 1, 0, 0, 9}, {0, 1, 0, 9, 0}, {0, 1, 9, 0, 0}, {0, 9, 0, 0, 1}, {0, 9, 0, 1, 0}, {0, 9, 1, 0, 0}, {1, 0, 0, 0, 9}, {1, 0, 0, 9, 0}, {1, 0, 9, 0, 0}, {1, 9, 0, 0, 0}, {9, 0, 0, 0, 1}, {9, 0, 0, 1, 0}, {9, 0, 1, 0, 0}, {9, 1, 0, 0, 0}}},
		//{"test#4", []int{-1, 2, 0, -1, 1, 0, 1}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		//{"test#5", []int{1, 1, 2, 2, 3, 3, 4}, [][]int{}},
		//{"test#6", []int{4, 2, 1, 3, 1}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
