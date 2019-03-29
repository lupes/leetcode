package question_31_40

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test#0", args{[]int{}, 1}, nil},
		{"test#1", args{[]int{1}, 2}, [][]int{{1, 1}}},
		{"test#2", args{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{"test#3", args{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"test#4", args{[]int{2, 8, 5, 4}, 10}, [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 4}, {2, 4, 4}, {2, 8}, {5, 5}}},
		{"test#5", args{[]int{7, 12, 5, 10, 9, 4, 6, 8}, 32}, [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 4}, {2, 4, 4}, {2, 8}, {5, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
