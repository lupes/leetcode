package question_301_310

import "testing"

func TestNumArray_SumRange(t *testing.T) {

	self := Constructor([]int{-2, 0, 3, -5, 2, -1})
	tests := []struct {
		name string
		this NumArray
		i    int
		j    int
		want int
	}{
		{"test", self, 0, 2, 1},
		{"test", self, 2, 5, -1},
		{"test", self, 0, 5, -3},
		{"test", self, -1, 5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.SumRange(tt.i, tt.j); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
