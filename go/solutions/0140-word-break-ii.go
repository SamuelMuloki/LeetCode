package solutions

type WordBreakTrie struct {
	isEndOfWord bool
	children    [26]*WordBreakTrie
	word        string
}

func WordBreak2(s string, wordDict []string) []string {
	var trie WordBreakTrie
	for _, word := range wordDict {
		curr := &trie
		for i := range word {
			idx := word[i] - 'a'
			if curr.children[idx] != nil {
				curr = curr.children[idx]
			} else {
				curr.children[idx] = &WordBreakTrie{}
				curr = curr.children[idx]
			}
		}
		curr.isEndOfWord = true
		curr.word = word
	}

	result := make([]string, 0)
	var backtrack func(int, string)
	backtrack = func(start int, sentence string) {
		if start == len(s) {
			if sentence != "" {
				result = append(result, sentence)
			}
			return
		}

		curr := &trie
		for i := start; i < len(s); i++ {
			curr = curr.children[s[i]-'a']
			if curr == nil {
				return
			}

			if curr.isEndOfWord {
				oldSentence := sentence
				if sentence != "" {
					sentence += " "
				}
				sentence += curr.word
				backtrack(i+1, sentence)
				sentence = oldSentence
			}
		}
	}
	backtrack(0, "")

	return result
}
