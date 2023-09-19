package solutions

import "sort"

func FindItinerary(tickets [][]string) []string {
	output := make([]string, 0)
	ticketGraph := make(map[string][]string)

	for _, ticket := range tickets {
		ticketGraph[ticket[0]] = append(ticketGraph[ticket[0]], ticket[1])
	}

	for key := range ticketGraph {
		sort.Sort(sort.Reverse(sort.StringSlice(ticketGraph[key])))
	}

	var dfs func(vertex string)
	dfs = func(vertex string) {
		for len(ticketGraph[vertex]) > 0 {
			next := ticketGraph[vertex][len(ticketGraph[vertex])-1]
			ticketGraph[vertex] = ticketGraph[vertex][:len(ticketGraph[vertex])-1]
			dfs(next)
		}
		output = append(output, vertex)
	}

	dfs("JFK")

	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}

	return output
}
