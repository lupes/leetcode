package common

import (
	"math"
	"reflect"
	"testing"
)

func TestNewNodeV2(t *testing.T) {
	tests := []struct {
		v    []int
		want *TreeNode
	}{
		//{[]int{1, 2, 3, 4, -99, -99, 7, 8, 9, -99, -99, 12, 13, -99, 14},
		//	NewNode(1,
		//		NewNode(2,
		//			NewNode(4, NewNode(8), NewNode(9)),
		//			NewNode(-99, NewNode(-99), NewNode(-99))),
		//		NewNode(3,
		//			NewNode(-99, NewNode(12), NewNode(13)),
		//			NewNode(7, NewNode(-99), NewNode(14))))},
		{[]int{1, 2, 3, 4, math.MaxInt32, math.MaxInt32, 7, 8, 9, math.MaxInt32, 14},
			NewNode(1,
				NewNode(2, NewNode(4, NewNode(8), NewNode(9))),
				NewNode(3, nil, NewNode(7, nil, NewNode(14)))),
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := NewNodeV2(tt.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNodeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
