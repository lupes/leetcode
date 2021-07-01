package question_131_140

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node1.Neighbors, node1)

	tests := []struct {
		node *Node
		want *Node
	}{
		{&Node{Val: 1}, &Node{Val: 1}},
		{&Node{Val: 1}, &Node{Val: 1}},
		{node1, &Node{Val: 1}},
	}
	for _, tt := range tests {

		t.Run("test", func(t *testing.T) {
			if got := cloneGraph(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
