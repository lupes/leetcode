package common

import (
	"fmt"
)

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Printf("%#v\n", row)
	}
	fmt.Println()
}
