package solutions

import "container/list"

type Current struct {
	node, distance int
}

func MaximumInvitations(favorite []int) int {
	n := len(favorite)
	reversedGraph := make([][]int, n)
	for person := 0; person < n; person++ {
		reversedGraph[favorite[person]] = append(
			reversedGraph[favorite[person]],
			person,
		)

	}

	bfs := func(startNode int, visitedNodes map[int]bool) int {
		q := list.New()
		q.PushBack(Current{startNode, 0})
		maxDistance := 0
		for q.Len() > 0 {
			element := q.Remove(q.Front()).(Current)
			currentNode, currentDistance := element.node, element.distance
			for _, neighbor := range reversedGraph[currentNode] {
				if visitedNodes[neighbor] {
					continue
				}

				visitedNodes[neighbor] = true
				q.PushBack(Current{neighbor, currentDistance + 1})
				maxDistance = max(maxDistance, currentDistance+1)
			}
		}

		return maxDistance
	}

	longestCycle, twoCycleInvitations := 0, 0
	visited := make([]bool, n)

	for person := 0; person < n; person++ {
		if visited[person] {
			continue
		}

		visitedPersons := make(map[int]int)
		current, distance := person, 0
		for {
			if visited[current] {
				break
			}
			visited[current], visitedPersons[current] = true, distance
			distance++
			nextPerson := favorite[current]
			if _, found := visitedPersons[nextPerson]; found {
				cycleLength := distance - visitedPersons[nextPerson]
				longestCycle = max(longestCycle, cycleLength)
				if cycleLength == 2 {
					vistedNodes := map[int]bool{
						current:    true,
						nextPerson: true,
					}
					twoCycleInvitations += 2 + bfs(nextPerson, vistedNodes) + bfs(current, vistedNodes)
				}
				break
			}
			current = nextPerson
		}
	}

	return max(longestCycle, twoCycleInvitations)
}
