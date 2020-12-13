package question_581_590

import (
	"reflect"
	"testing"
)

func Test_postorder(t *testing.T) {
	tests := []struct {
		root *Node
		want []int
	}{
		{&Node{Val: 1, Children: []*Node{{3, []*Node{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4}}}, []int{5, 6, 3, 2, 4, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := postorder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
