package solutions

import "math/rand"

type RandomizedSet struct {
	set     map[int]int
	indices []int
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{
		set:     make(map[int]int),
		indices: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}

	this.indices = append(this.indices, val)
	this.set[val] = len(this.indices) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.set[val]
	if !ok {
		return false
	}

	lastIdx := len(this.indices) - 1
	lastElem := this.indices[lastIdx]
	this.indices[idx] = lastElem
	this.set[lastElem] = idx
	delete(this.set, val)
	this.indices = this.indices[:len(this.indices)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.indices[rand.Intn(len(this.indices))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
