package question_421_430

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		root *Node5
		want [][]int
	}{
		{&Node5{
			Val: 1,
			Children: []*Node5{
				{
					Val: 3,
					Children: []*Node5{
						{
							Val:      5,
							Children: nil,
						},
						{
							Val:      6,
							Children: nil,
						},
					},
				},
				{
					Val: 2,
				},
				{
					Val: 4,
				},
			},
		},
			[][]int{{1}, {3, 2, 4}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := levelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
