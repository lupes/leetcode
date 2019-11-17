package question_111_120

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{"test", 0, []int{1}},
		{"test", 1, []int{1, 1}},
		{"test", 2, []int{1, 2, 1}},
		{"test", 3, []int{1, 3, 3, 1}},
		{"test", 4, []int{1, 4, 6, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
