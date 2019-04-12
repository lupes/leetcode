package question_51_60

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name      string
		intervals []Interval
		want      []Interval
	}{
		{"test#0", nil, nil},
		{"test#1", []Interval{{1, 2}}, []Interval{{1, 2}}},
		{"test#2", []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, []Interval{{1, 6}, {8, 10}, {15, 18}}},
		{"test#3", []Interval{{1, 4}, {4, 5}}, []Interval{{1, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
