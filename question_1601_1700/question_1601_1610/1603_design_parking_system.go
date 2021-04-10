package question_1601_1610

// 1603. 设计停车系统
// https://leetcode-cn.com/problems/design-parking-system/
// Topics: 设计

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.Big > 0 {
		this.Big--
	} else if carType == 2 && this.Medium > 0 {
		this.Medium--
	} else if carType == 3 && this.Small > 0 {
		this.Small--
	} else {
		return false
	}
	return true
}
