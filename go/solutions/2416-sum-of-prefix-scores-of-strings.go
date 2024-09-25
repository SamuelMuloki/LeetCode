package solutions

type PrefixTrie struct {
	children [26]*PrefixTrie
	count    int
	root     *PrefixTrie
}

func NewPrefixTrie() *PrefixTrie {
	return &PrefixTrie{
		children: [26]*PrefixTrie{},
		root:     &PrefixTrie{},
	}
}

func (trie *PrefixTrie) Insert(word string) {
	curr := trie.root
	for _, ch := range word {
		if curr.children[ch-'a'] == nil {
			curr.children[ch-'a'] = &PrefixTrie{}
		}
		curr = curr.children[ch-'a']
		curr.count++
	}
}

func (trie *PrefixTrie) PrefixCount(word string) int {
	curr := trie.root
	sum := 0
	for _, ch := range word {
		curr = curr.children[ch-'a']
		sum += curr.count
	}

	return sum
}

func SumPrefixScores(words []string) []int {
	prefTrie := NewPrefixTrie()

	for _, word := range words {
		prefTrie.Insert(word)
	}

	ans := make([]int, len(words))
	for i, word := range words {
		ans[i] = prefTrie.PrefixCount(word)
	}

	return ans
}
