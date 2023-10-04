package solutions

type node struct {
	key, val int
	next     *node
}

type MyHashMap struct {
	size, mult int
	data       []*node
}

func HashConstructor() MyHashMap {
	return MyHashMap{size: 19997, mult: 12582917, data: make([]*node, 19997)}
}

func (this *MyHashMap) Hash(key int) int {
	return (key * this.mult) % this.size
}

func (this *MyHashMap) Put(key int, value int) {
	this.Remove(key)
	pos := this.Hash(key)
	this.data[pos] = &node{key: key, val: value, next: this.data[pos]}
}

func (this *MyHashMap) Get(key int) int {
	pos := this.Hash(key)
	head := this.data[pos]

	for ; head != nil; head = head.next {
		if head.key == key {
			return head.val
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	pos := this.Hash(key)
	head := this.data[pos]

	if head == nil {
		return
	}

	if head.key == key {
		this.data[pos] = head.next
	} else {
		for ; head.next != nil; head = head.next {
			if head.next.key == key {
				head.next = head.next.next
				return
			}
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
