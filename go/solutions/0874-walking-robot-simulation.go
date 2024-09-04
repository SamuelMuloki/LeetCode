package solutions

func RobotSim(commands []int, obstacles [][]int) int {
	obstacleMap := make(map[int]bool)
	HASH_MULTIPLIER := 60001
	var hashCoordinates = func(x, y int) int {
		return x + HASH_MULTIPLIER*y
	}

	for i := range obstacles {
		obstacleMap[hashCoordinates(obstacles[i][0], obstacles[i][1])] = true
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currentPosition := []int{0, 0}
	maxDistanceSquared, currentDirection := 0, 0
	for _, command := range commands {
		if command == -1 {
			currentDirection = (currentDirection + 1) % 4
			continue
		}

		if command == -2 {
			currentDirection = (currentDirection + 3) % 4
			continue
		}

		direction := directions[currentDirection]
		for step := 0; step < command; step++ {
			nextX := currentPosition[0] + direction[0]
			nextY := currentPosition[1] + direction[1]
			if obstacleMap[hashCoordinates(nextX, nextY)] {
				break
			}
			currentPosition[0] = nextX
			currentPosition[1] = nextY
		}

		maxDistanceSquared = max(
			maxDistanceSquared,
			currentPosition[0]*currentPosition[0]+currentPosition[1]*currentPosition[1],
		)
	}

	return maxDistanceSquared
}
