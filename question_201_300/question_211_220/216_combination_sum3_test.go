package question_211_220

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{"test", 55, 7, nil},
		{"test", 1, 1, [][]int{{1}}},
		{"test", 2, 3, [][]int{{1, 2}}},
		{"test", 3, 7, [][]int{{1, 2, 4}}},
		{"test", 3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{"test", 9, 45, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.k, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
