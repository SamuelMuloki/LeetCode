package solutions

func MaxBottlesDrunk(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		numBottles -= numExchange
		res++
		numExchange++
		numBottles++
	}

	return res
}
