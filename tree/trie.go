package tree

// Trie ...
type Trie struct {
	size int
	root *trieNode
}

type trieNode struct {
	key        rune
	isTerminal bool
	value      interface{}
	children   map[rune]*trieNode
}

// NewTrie ...
func NewTrie() *Trie {
	t := &Trie{}
	return t
}

// Put ...
func (t *Trie) Put(key string, value interface{}) {
	if t.root == nil {
		t.root = &trieNode{
			children: make(map[rune]*trieNode, 0),
		}
	}
	n := t.root
	for _, c := range key {
		next, ok := n.children[c]
		if !ok {
			next = &trieNode{key: c}
			n.children[c] = next
		}
		n = next
	}
	n.isTerminal = true
}

// Has ...
func (t *Trie) Has(key string) bool {
	if t.root == nil {
		return false
	}
	n := t.root
	for _, c := range key {
		next, ok := n.children[c]
		if !ok {
			return false
		}
		n = next
	}
	return n.isTerminal
}

// Del ...
func (t *Trie) Del(key string) {
	if t.root == nil {
		return
	}
}
