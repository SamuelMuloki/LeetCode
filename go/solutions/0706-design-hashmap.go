package solutions

type Pair struct {
	key, value int
}

type MyHashMap struct {
	Pairs []Pair
}

func HashConstructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {
	for i := range this.Pairs {
		if this.Pairs[i].key == key {
			this.Pairs[i].value = value
			return
		}
	}

	this.Pairs = append(this.Pairs, Pair{
		key:   key,
		value: value,
	})
}

func (this *MyHashMap) Get(key int) int {
	for i := range this.Pairs {
		if this.Pairs[i].key == key {
			return this.Pairs[i].value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	for i := 0; i < len(this.Pairs); i++ {
		if this.Pairs[i].key == key {
			rem := this.Pairs[:i]
			if i < len(this.Pairs) {
				rem = append(rem, this.Pairs[i+1:]...)
			}
			this.Pairs = rem
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
