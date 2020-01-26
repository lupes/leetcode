package question_721_730

import (
	"fmt"
	"testing"
)

func TestMyCalendar_Book(t *testing.T) {
	type arg struct {
		start int
		end   int
		want  bool
	}
	tests := []struct {
		args [][]arg
	}{
		{
			args: [][]arg{
				{
					{start: 10, end: 20, want: true},
					{start: 15, end: 25, want: false},
					{start: 20, end: 30, want: true},
					{start: 5, end: 6, want: true},
					{start: 6, end: 7, want: true},
					{start: 1, end: 8, want: false},
				},
				{
					{start: 37, end: 50, want: true},
					{start: 33, end: 50, want: false},
					{start: 4, end: 17, want: true},
					{start: 35, end: 48, want: false},
					{start: 8, end: 25, want: false},
				},
				{
					{start: 47, end: 50, want: true},
					{start: 33, end: 41, want: true},
					{start: 39, end: 45, want: false},
					{start: 33, end: 42, want: false},
					{start: 25, end: 32, want: true},
					{start: 26, end: 35, want: false},
					{start: 19, end: 25, want: true},
					{start: 3, end: 8, want: true},
					{start: 8, end: 13, want: true},
					{start: 18, end: 27, want: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			for _, args := range tt.args {
				myCalendar := Constructor()
				for _, arg := range args {
					if got := myCalendar.Book(arg.start, arg.end); got != arg.want {
						t.Errorf("Book(%d, %d) = %v, want %v", arg.start, arg.end, got, arg.want)
					}
				}
				fmt.Println()
			}
		})
	}
}
