package question_421_430

import (
	"reflect"
	"testing"
)

func Test_construct(t *testing.T) {
	tests := []struct {
		grid [][]int
		want *Node2
	}{
		{[][]int{{0}}, &Node2{IsLeaf: true, Val: false}},
		{[][]int{{1}}, &Node2{IsLeaf: true, Val: true}},
		{[][]int{{1, 1}, {1, 1}}, &Node2{IsLeaf: true, Val: true}},
		{[][]int{{0, 1}, {1, 0}}, &Node2{IsLeaf: false, Val: true, TopLeft: &Node2{IsLeaf: true, Val: false}, TopRight: &Node2{IsLeaf: true, Val: true}, BottomLeft: &Node2{IsLeaf: true, Val: true}, BottomRight: &Node2{IsLeaf: true, Val: false}}},
		{[][]int{
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0}}, &Node2{IsLeaf: false, Val: true,
			TopLeft: &Node2{IsLeaf: true, Val: true},
			TopRight: &Node2{IsLeaf: false, Val: true,
				TopLeft:     &Node2{IsLeaf: true, Val: false},
				TopRight:    &Node2{IsLeaf: true, Val: false},
				BottomLeft:  &Node2{IsLeaf: true, Val: true},
				BottomRight: &Node2{IsLeaf: true, Val: true},
			},
			BottomLeft:  &Node2{IsLeaf: true, Val: true},
			BottomRight: &Node2{IsLeaf: true, Val: false},
		}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := construct(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}
