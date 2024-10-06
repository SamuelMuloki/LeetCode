package solutions

import "sort"

func DividePlayers(skill []int) int64 {
	sort.Ints(skill)

	totalSkill := skill[0] + skill[len(skill)-1]
	var chemistrySum int64 = 0

	for i := 0; i < len(skill)/2; i++ {
		if skill[i]+skill[len(skill)-i-1] != totalSkill {
			return -1
		}
		chemistrySum += int64(skill[i]) * int64(skill[len(skill)-i-1])
	}

	return chemistrySum
}
