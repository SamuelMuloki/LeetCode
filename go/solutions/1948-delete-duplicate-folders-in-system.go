package solutions

import (
	"sort"
	"strings"
)

type Trie2 struct {
	serial   string
	children map[string]*Trie2
}

func DeleteDuplicateFolder(paths [][]string) [][]string {
	root := &Trie2{children: make(map[string]*Trie2)}
	for _, path := range paths {
		cur := root
		for _, node := range path {
			if _, ok := cur.children[node]; !ok {
				cur.children[node] = &Trie2{children: make(map[string]*Trie2)}
			}
			cur = cur.children[node]
		}
	}

	freq := make(map[string]int)
	var construct func(*Trie2)
	construct = func(node *Trie2) {
		if len(node.children) == 0 {
			return
		}
		v := make([]string, 0, len(node.children))
		for folder, child := range node.children {
			construct(child)
			v = append(v, folder+"("+child.serial+")")
		}
		sort.Strings(v)
		node.serial = strings.Join(v, "")
		freq[node.serial]++
	}

	construct(root)
	ans := make([][]string, 0)
	path := make([]string, 0)

	var operate func(*Trie2)
	operate = func(node *Trie2) {
		if freq[node.serial] > 1 {
			return
		}

		if len(path) > 0 {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}

		for folder, child := range node.children {
			path = append(path, folder)
			operate(child)
			path = path[:len(path)-1]
		}
	}

	operate(root)
	return ans
}
