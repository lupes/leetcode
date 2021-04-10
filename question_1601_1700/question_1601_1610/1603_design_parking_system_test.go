package question_1601_1610

import (
	"fmt"
	"testing"
)

func TestParkingSystem(t *testing.T) {
	this := Constructor(1, 1, 0)
	fmt.Printf("%v\n", this.AddCar(1))
	fmt.Printf("%v\n", this.AddCar(2))
	fmt.Printf("%v\n", this.AddCar(3))
	fmt.Printf("%v\n", this.AddCar(1))
}
