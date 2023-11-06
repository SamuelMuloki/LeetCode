package solutions

type Trie struct {
	children  [26]*Trie
	endOfWord bool
	root      *Trie
}

func TrieConstructor() Trie {
	return Trie{
		children: [26]*Trie{},
		root:     &Trie{},
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, ch := range word {
		if curr.children[ch-'a'] == nil {
			curr.children[ch-'a'] = &Trie{}
		}
		curr = curr.children[ch-'a']
	}
	curr.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, ch := range word {
		if curr.children[ch-'a'] == nil {
			return false
		}
		curr = curr.children[ch-'a']
	}
	return curr.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, ch := range prefix {
		if curr.children[ch-'a'] == nil {
			return false
		}
		curr = curr.children[ch-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
