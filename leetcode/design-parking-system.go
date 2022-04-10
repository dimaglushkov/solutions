package main

// source: https://leetcode.com/problems/design-parking-system/

type ParkingSystem [3]int

// NewParkingSystem rename to Constructor for leetcode
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	if ps[carType-1] == 0 {
		return false
	}
	ps[carType-1]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
func main() {
	parkingSystem := Constructor(1, 1, 0)
	println(parkingSystem.AddCar(1))
	println(parkingSystem.AddCar(2))
	println(parkingSystem.AddCar(3))
	println(parkingSystem.AddCar(1))
}
