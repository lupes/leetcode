package question_71_80

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{"test#0", 1, 1, [][]int{{1}}},
		{"test#1", 2, 1, [][]int{{1}, {2}}},
		{"test#2", 3, 2, [][]int{{1, 2}, {1, 3}, {2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.n, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
