package solutions

func DestCity(paths [][]string) string {
	graph := make(map[string]string)
	for i := range paths {
		graph[paths[i][0]] = paths[i][1]
	}

	for i := range paths {
		if _, ok := graph[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}

	return ""
}
