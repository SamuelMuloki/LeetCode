package solutions

func CheckIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adj := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		adj[prerequisites[i][0]] = append(
			adj[prerequisites[i][0]],
			prerequisites[i][1],
		)
	}

	var dfs func(node, target int, visited []bool) bool
	dfs = func(node, target int, visited []bool) bool {
		visited[node] = true
		if node == target {
			return true
		}

		val := false
		for _, nei := range adj[node] {
			if !visited[nei] {
				val = val || dfs(nei, target, visited)
			}
		}

		return val
	}

	ans := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		visited := make([]bool, numCourses)
		ans[i] = dfs(queries[i][0], queries[i][1], visited)
	}

	return ans
}
