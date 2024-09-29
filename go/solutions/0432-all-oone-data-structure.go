package solutions

import "sort"

type AllOne struct {
	table map[string]int
	st    []AllOnePair
}

type AllOnePair struct {
	count int
	key   string
}

func AllOneConstructor() AllOne {
	return AllOne{
		table: make(map[string]int),
		st:    []AllOnePair{},
	}
}

func (this *AllOne) Inc(key string) {
	if val, exists := this.table[key]; exists {
		this.remove(val, key)
	}

	this.table[key]++

	this.st = append(this.st, AllOnePair{this.table[key], key})

	sort.Slice(this.st, func(i, j int) bool {
		if this.st[i].count == this.st[j].count {
			return this.st[i].key < this.st[j].key
		}
		return this.st[i].count < this.st[j].count
	})
}

func (this *AllOne) Dec(key string) {
	if val, exists := this.table[key]; exists {
		this.remove(val, key)

		this.table[key]--

		if this.table[key] == 0 {
			delete(this.table, key)
		} else {
			this.st = append(this.st, AllOnePair{this.table[key], key})

			// Sort the slice again
			sort.Slice(this.st, func(i, j int) bool {
				if this.st[i].count == this.st[j].count {
					return this.st[i].key < this.st[j].key
				}
				return this.st[i].count < this.st[j].count
			})
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if len(this.st) == 0 {
		return ""
	}
	return this.st[len(this.st)-1].key
}

func (this *AllOne) GetMinKey() string {
	if len(this.st) == 0 {
		return ""
	}

	return this.st[0].key
}

func (this *AllOne) remove(count int, key string) {
	for i, p := range this.st {
		if p.count == count && p.key == key {
			this.st = append(this.st[:i], this.st[i+1:]...)
			break
		}
	}
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
