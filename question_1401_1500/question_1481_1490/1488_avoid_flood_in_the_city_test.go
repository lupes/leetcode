package question_1481_1490

import (
	"reflect"
	"testing"
)

func Test_avoidFlood(t *testing.T) {
	tests := []struct {
		rains []int
		want  []int
	}{
		//{[]int{1, 2, 3, 4}, []int{-1, -1, -1, -1}},
		//{[]int{1, 2, 0, 0, 2, 1}, []int{-1, -1, 2, 1, -1, -1}},
		//{[]int{1, 2, 0, 1, 2}, nil},
		//{[]int{69, 0, 0, 0, 69}, []int{-1, 69, 1, 1, -1}},
		//{[]int{10, 20, 20}, nil},
		{[]int{1, 0, 2, 0, 3, 0, 2, 0, 0, 0, 1, 2, 3}, []int{-1, 1, -1, 2, -1, 3, -1, 2, 1, 1, -1, -1, -1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := avoidFlood(tt.rains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
