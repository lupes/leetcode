package question_491_500

import "testing"

func Test_findPoisonedDuration(t *testing.T) {
	tests := []struct {
		name       string
		timeSeries []int
		duration   int
		want       int
	}{
		{"test", []int{1}, 2, 2},
		{"test", []int{1, 4}, 2, 4},
		{"test", []int{1, 2}, 2, 3},
		{"test", []int{1, 2, 3}, 2, 4},
		{"test", []int{1, 2, 3, 5}, 2, 6},
		{"test", []int{1, 2, 3, 4, 5}, 5, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoisonedDuration(tt.timeSeries, tt.duration); got != tt.want {
				t.Errorf("findPoisonedDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
