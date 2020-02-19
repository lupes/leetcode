package question_821_830

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	tests := []struct {
		S    string
		want [][]int
	}{
		{"abbxxxxzzy", [][]int{{3, 6}}},
		{"abc", nil},
		{"aaa", [][]int{{0, 2}}},
		{"aaabb", [][]int{{0, 2}}},
		{"abcdddeeeeaabbbcd", [][]int{{3, 5}, {6, 9}, {12, 14}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := largeGroupPositions(tt.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
