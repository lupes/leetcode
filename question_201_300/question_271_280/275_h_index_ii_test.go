package question_271_280

import "testing"

func Test_hIndex2(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{"test", []int{0}, 0},
		{"test", []int{1}, 1},
		{"test", []int{100}, 1},
		{"test", []int{100, 102}, 2},
		{"test", []int{0, 1}, 1},
		{"test", []int{0, 1, 2}, 1},
		{"test", []int{0, 1, 2, 3}, 2},
		{"test", []int{0, 0, 1, 3, 5, 6}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex2(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
