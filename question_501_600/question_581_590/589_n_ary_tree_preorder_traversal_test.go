package question_581_590

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	tests := []struct {
		root *Node
		want []int
	}{
		{&Node{Val: 1, Children: []*Node{{3, []*Node{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4}}}, []int{1, 3, 5, 6, 2, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := preorder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
