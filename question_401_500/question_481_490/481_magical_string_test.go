package question_481_490

import (
	"testing"
)

//  5 1 2 2 1 1
// 10 2 1 2 2 1
// 15 2 2 1 1 2
// 20 1 1 2 2 1
// 25 2 1 1 2 1
// 30 2 2 1 1 2
// 1 1 2 1 2
// 2 1 2 2 1
// 1 2 1 2 2
// 1 2 1 1 2
// 1 1 2 2 1 2 2 1 1 2 1 2 2 1 2 2 1 1 2 1 1 2 1 2 2 1 2 1 1 2 2 1 2 2 1 1 2 1 2 2 1 2 2 1 1 2 1 1 2 2 1 0

// 12211
// 21221
// 22112
// 1122

func Test_magicalString(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 3},
		{7, 4},
		{8, 4},
		{9, 4},
		{10, 5},
		{11, 5},
		{12, 5},
		{13, 6},
		{14, 7},
		{15, 7},
		{16, 8},
		{17, 9},
		{18, 9},
		{19, 9},
		//{20, 10},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := magicalString(tt.n); got != tt.want {
				t.Errorf("magicalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
