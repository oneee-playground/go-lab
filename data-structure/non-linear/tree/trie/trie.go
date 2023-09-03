package trie

type node struct {
	children   map[rune]*node
	isLastRune bool
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: &node{children: make(map[rune]*node)},
	}
}

func (t *Trie) Insert(str string) {
	var current = t.root

	for _, r := range str {
		if _, ok := current.children[r]; !ok {
			current.children[r] = &node{
				children: make(map[rune]*node),
			}
		}
		current = current.children[r]
	}
	current.isLastRune = true
}

func (t *Trie) Find(str string) (ok bool) {
	var current = t.root

	for _, r := range str {
		current, ok = current.children[r]
		if !ok {
			return false
		}
	}
	return true
}

func (t *Trie) Delete(str string) {
	cursor := 0
	t.root.delete([]rune(str), &cursor)
}

func (n *node) delete(str []rune, cursor *int) bool {
	if *cursor == len(str) {
		return true
	}

	r := str[*cursor]

	target, found := n.children[r]
	if !found {
		return false
	}

	(*cursor)++
	ok := target.delete(str, cursor)
	if !ok {
		return false
	}

	if len(target.children) != 0 || (target.isLastRune && r != str[len(str)-1]) {
		return false
	}

	delete(n.children, r)

	return true
}
