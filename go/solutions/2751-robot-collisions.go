package solutions

import "slices"

type robot struct {
	dir    byte
	health int
	pos    int
	ind    int
}

func SurvivedRobotsHealths(positions []int, healths []int, directions string) []int {
	robots := make([]robot, len(positions))
	for i := range positions {
		robots[i] = robot{
			dir:    directions[i],
			health: healths[i],
			pos:    positions[i],
			ind:    i,
		}
	}

	slices.SortFunc(robots, func(a, b robot) int {
		return a.pos - b.pos
	})

	stack := []robot{}
	for _, r := range robots {
		if r.dir == 'R' {
			stack = append(stack, r)
			continue
		}
		for {
			if len(stack) == 0 || stack[len(stack)-1].dir == 'L' {
				stack = append(stack, r)
				break
			}
			if stack[len(stack)-1].health == r.health {
				stack = stack[:len(stack)-1]
				break
			}
			if stack[len(stack)-1].health > r.health {
				stack[len(stack)-1].health--
				break
			}
			stack = stack[:len(stack)-1]
			r.health--
		}
	}

	slices.SortFunc(stack, func(a, b robot) int {
		return a.ind - b.ind
	})

	res := make([]int, len(stack))
	for i := range res {
		res[i] = stack[i].health
	}
	return res
}
