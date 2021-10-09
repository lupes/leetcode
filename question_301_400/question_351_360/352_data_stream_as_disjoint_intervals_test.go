package question_351_360

import (
	"fmt"
	"testing"
)

func TestSummaryRanges(t *testing.T) {

	this := Constructor2()
	this.AddNum(1)
	fmt.Printf("1 %v\n", this.GetIntervals())
	this.AddNum(0)
	fmt.Printf("0 %v\n", this.GetIntervals())

	this.AddNum(3)
	fmt.Printf("3 %v\n", this.GetIntervals())
	this.AddNum(7)
	fmt.Printf("7 %v\n", this.GetIntervals())
	this.AddNum(2)
	fmt.Printf("2 %v\n", this.GetIntervals())
	this.AddNum(6)
	fmt.Printf("6 %v\n", this.GetIntervals())
}
