package solutions

type ParkingSystem struct {
	CarType map[int]int
}

func ParkingConstructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		CarType: map[int]int{1: big, 2: medium, 3: small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.CarType[carType] > 0 {
		this.CarType[carType]--
		return true
	}

	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
