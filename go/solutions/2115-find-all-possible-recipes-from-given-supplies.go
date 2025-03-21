package solutions

func FindAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	supplySet := make(map[string]bool)
	for _, supply := range supplies {
		supplySet[supply] = true
	}

	graph := make(map[string][]string) // ingredient -> recipes that need it
	inDegree := make(map[string]int)   // How many unavailable ingredients each recipe needs
	available := make(map[string]bool)
	queue := []string{}

	for i, recipe := range recipes {
		unavailableCount := 0 // count recipes not in supply set
		for _, ingredient := range ingredients[i] {
			if !supplySet[ingredient] { // ingredient not in initial supplies
				graph[ingredient] = append(graph[ingredient], recipe)
				unavailableCount++
			}
		}

		if unavailableCount == 0 { // All ingredients are available
			queue = append(queue, recipe)
			available[recipe] = true
		} else {
			inDegree[recipe] = unavailableCount
		}
	}

	res := []string{}
	for len(queue) > 0 {
		currRecipe := queue[0]
		queue = queue[1:]

		res = append(res, currRecipe)
		for _, nextRecipe := range graph[currRecipe] {
			inDegree[nextRecipe]--
			if inDegree[nextRecipe] == 0 {
				queue = append(queue, nextRecipe)
				available[nextRecipe] = true
			}
		}
	}

	return res
}
