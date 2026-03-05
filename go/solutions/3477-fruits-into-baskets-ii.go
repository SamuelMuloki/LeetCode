package solutions

func NumOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	res := 0
	for i := 0; i < n; i++ {
		canBePlaced := false
		for j := 0; j < n; j++ {
			if fruits[i] <= baskets[j] {
				canBePlaced = true
				baskets[j] = -1
				break
			}
		}

		if !canBePlaced {
			res++
		}
	}

	return res
}
