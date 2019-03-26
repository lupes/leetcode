package question_31_40

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test#0", args{nil, 0}, -1},
		{"test#1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"test#2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
