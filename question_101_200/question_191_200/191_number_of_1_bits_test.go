package question_191_200

import "testing"

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want int
	}{
		{"test", 1, 1},
		{"test", 2, 1},
		{"test", 3, 2},
		{"test", 4, 1},
		{"test", 5, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
