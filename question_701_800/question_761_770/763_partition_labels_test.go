package question_761_770

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	tests := []struct {
		S    string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := partitionLabels(tt.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
