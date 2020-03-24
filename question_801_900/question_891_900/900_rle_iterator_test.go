package question_891_900

import (
	"fmt"
	"testing"
)

func TestRLEIterator(t *testing.T) {

	iterator := Constructor([]int{3, 8, 0, 9, 2, 5})

	if got := iterator.Next(2); got != 8 {
		t.Errorf("Next() = %v", got)
	}

	if got := iterator.Next(1); got != 8 {
		t.Errorf("Next() = %v", got)
	}

	if got := iterator.Next(1); got != 5 {
		t.Errorf("Next() = %v", got)
	}

	if got := iterator.Next(2); got != -1 {
		t.Errorf("Next() = %v", got)
	}

	iterator = Constructor([]int{784, 303, 477, 583, 909, 505})
	if got := iterator.Next(130); got != 303 {
		t.Errorf("Next() = %v", got)
	}
	fmt.Println(iterator.A) // 654, 303, 477, 583, 909, 505

	if got := iterator.Next(333); got != 303 {
		t.Errorf("Next() = %v", got)
	}
	fmt.Println(iterator.A) // 321, 303, 477, 583, 909, 505

	if got := iterator.Next(238); got != 303 {
		t.Errorf("Next() = %v", got)
	}
	fmt.Println(iterator.A) // 83, 303, 477, 583, 909, 505

	if got := iterator.Next(87); got != 583 {
		t.Errorf("Next() = %v", got)
	}
	fmt.Println(iterator.A) // 473, 583, 909, 505

	if got := iterator.Next(301); got != 583 {
		t.Errorf("Next() = %v", got)
	}
	fmt.Println(iterator.A)

	if got := iterator.Next(276); got != 505 {
		t.Errorf("Next() = %v", got)
	}
	fmt.Println(iterator.A)
}
