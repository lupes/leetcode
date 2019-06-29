package question_211_220

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"test", []int{1, 2, 3, 1}, 3, true},
		{"test", []int{1, 0, 1, 1}, 1, true},
		{"test", []int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.nums, tt.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
