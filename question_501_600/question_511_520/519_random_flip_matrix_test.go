package qustion_511_520

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	solution := Constructor(2, 3)
	for i := 0; i < 12; i++ {
		fmt.Printf("%+v\n", solution.Flip())
	}
}
