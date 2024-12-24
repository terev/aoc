package util

type TrieNode struct {
	children   [26]*TrieNode
	val        rune
	isTerminal bool
}

type Trie [26]*TrieNode

func (t *Trie) Insert(item string) {
	node := t[item[0]-'a']
	if node == nil {
		node = &TrieNode{}
		t[item[0]-'a'] = node
	}
	i := 1
	for i < len(item) {
		charIdx := item[i] - 'a'
		nextNode := node.children[charIdx]
		if nextNode == nil {
			nextNode = &TrieNode{}
			node.children[charIdx] = nextNode
		}
		node = nextNode
		i++
	}
	node.isTerminal = true
}

func (t *Trie) LongestPrefix(item string) string {
	if len(item) == 0 {
		return ""
	}

	var longest int

	node := t[item[0]-'a']
	if node == nil {
		return ""
	}
	if node.isTerminal {
		longest = 1
	}

	i := 1
	for i < len(item) {
		charIdx := item[i] - 'a'
		node = node.children[charIdx]

		if node == nil {
			break
		}
		if node.isTerminal {
			longest = i + 1
		}

		i++
	}
	if longest == 0 {
		return ""
	}
	return item[:longest]
}

func (t *Trie) Prefixes(item string) []string {
	if len(item) == 0 {
		return []string{}
	}

	node := t[item[0]-'a']
	if node == nil {
		return []string{}
	}

	prefixes := []string{}
	if node.isTerminal {
		prefixes = append(prefixes, item[:1])
	}
	i := 1
	for i < len(item) {
		charIdx := item[i] - 'a'
		nextNode := node.children[charIdx]
		if nextNode == nil {
			return prefixes
		}
		node = nextNode
		if node.isTerminal {
			prefixes = append(prefixes, item[:i+1])
		}
		i++
	}
	return prefixes
}
