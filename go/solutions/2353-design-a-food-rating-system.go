package solutions

import "container/heap"

type Food struct {
	Name   string
	Rating int
}

type FoodRatings struct {
	CuisineRatings map[string]*FoodRatingsMaxHeap
	FoodCuisine    map[string]string
	FoodRating     map[string]int
}

func FoodRatingsConstructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodRating := make(map[string]int)
	foodCuisine := make(map[string]string)
	cuisineRatings := make(map[string]*FoodRatingsMaxHeap)

	for i := 0; i < len(foods); i++ {
		foodRating[foods[i]] = ratings[i]
		foodCuisine[foods[i]] = cuisines[i]

		if _, ok := cuisineRatings[cuisines[i]]; !ok {
			cuisineRatings[cuisines[i]] = &FoodRatingsMaxHeap{}
		}

		heap.Push(cuisineRatings[cuisines[i]], Food{foods[i], ratings[i]})
	}

	return FoodRatings{
		CuisineRatings: cuisineRatings,
		FoodCuisine:    foodCuisine,
		FoodRating:     foodRating,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.FoodRating[food] = newRating
	cuisine := this.FoodCuisine[food]

	heap.Push(this.CuisineRatings[cuisine], Food{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	highestRated := (*this.CuisineRatings[cuisine])[0]

	for highestRated.Rating != this.FoodRating[highestRated.Name] {
		heap.Pop(this.CuisineRatings[cuisine])
		highestRated = (*this.CuisineRatings[cuisine])[0]
	}

	return highestRated.Name
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

type FoodRatingsMaxHeap []Food

func (h FoodRatingsMaxHeap) Len() int { return len(h) }
func (h FoodRatingsMaxHeap) Less(i, j int) bool {
	if h[i].Rating == h[j].Rating {
		return h[i].Name < h[j].Name
	}

	return h[i].Rating > h[j].Rating
}
func (h FoodRatingsMaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *FoodRatingsMaxHeap) Push(x any)   { *h = append(*h, x.(Food)) }
func (h *FoodRatingsMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
