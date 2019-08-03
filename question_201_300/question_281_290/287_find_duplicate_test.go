package question_281_290

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", []int{1, 3, 4, 2, 2}, 2},
		{"test", []int{3, 1, 3, 4, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
