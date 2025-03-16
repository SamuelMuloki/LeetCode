package solutions

import "math"

func RepairCars(ranks []int, cars int) int64 {
	l, r := int64(1), int64(cars*cars*ranks[0])
	for l < r {
		mid := l + (r-l)/2
		carsRepaired := int64(0)
		for _, rank := range ranks {
			carsRepaired += int64(math.Sqrt(float64(mid) / float64(rank)))
		}

		if carsRepaired >= int64(cars) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
