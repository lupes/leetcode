package question_391_400

import "testing"

func TestSolution_Pick(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"test", []int{1, 2, 3, 3, 3}, 3, []int{2, 3, 4}},
		{"test", []int{1, 2, 3, 3, 3}, 1, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.nums)
			got := this.Pick(tt.target)
			for _, n := range tt.want {
				if n == got {
					goto Next
				}
			}
			t.Errorf("Solution.Pick() = %v, want %v", got, tt.want)
			return
		Next:
			t.Logf("Solution.Pick() = %v, want %v", got, tt.want)
		})
	}
}
