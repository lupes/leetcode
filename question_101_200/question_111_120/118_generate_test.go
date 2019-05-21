package question_111_120

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	want := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
		{1, 5, 10, 10, 5, 1},
		{1, 6, 15, 20, 15, 6, 1},
	}
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{"test", -1, nil},
		{"test", 0, nil},
		{"test", 1, want[0:1]},
		{"test", 2, want[0:2]},
		{"test", 3, want[0:3]},
		{"test", 4, want[0:4]},
		{"test", 5, want[0:5]},
		{"test", 6, want[0:6]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
