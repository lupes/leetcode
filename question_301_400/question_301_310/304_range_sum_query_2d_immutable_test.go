package question_301_310

import "testing"

func TestNumMatrix_SumRegion(t *testing.T) {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	tests := []struct {
		name   string
		matrix [][]int
		row1   int
		col1   int
		row2   int
		col2   int
		want   int
	}{
		{"test", matrix, 0, 0, 0, 0, 3},
		{"test", matrix, 2, 1, 4, 3, 8},
		{"test", matrix, 1, 1, 2, 2, 11},
		{"test", matrix, 1, 2, 2, 4, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := ConstructorMatrix(tt.matrix)
			if got := this.SumRegion(tt.row1, tt.col1, tt.row2, tt.col2); got != tt.want {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
