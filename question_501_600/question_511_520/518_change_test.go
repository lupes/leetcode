package qustion_511_520

import "testing"

func Test_change(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{"test", 5, []int{1, 2, 5}, 4},
		{"test", 3, []int{2}, 0},
		{"test", 10, []int{10}, 1},
		{"test", 10, []int{1, 2, 5}, 10},
		{"test", 500, []int{3, 5, 7, 8, 9, 10, 11}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.amount, tt.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
1*10
1*8+2*1
1*6+2*2
1*5+5*1
1*4+2*3
1*3+2*1+5*1
1*2+2*4
1*1+2*2+5*1
1*0+2*5
1*0+5*2
*/
