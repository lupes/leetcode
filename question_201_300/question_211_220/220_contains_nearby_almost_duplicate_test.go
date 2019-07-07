package question_211_220

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		t    int
		want bool
	}{
		{"test", []int{1, 2, 3, 1}, 3, 1, true},
		{"test", []int{1, 0, 1, 1}, 1, 2, true},
		{"test", []int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		{"test", []int{3, 6, 0, 4}, 2, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
