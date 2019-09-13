package question_411_420

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{"test", 1, []string{"1"}},
		{"test", 3, []string{"1", "2", "Fizz"}},
		{"test", 5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"test", 6, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
