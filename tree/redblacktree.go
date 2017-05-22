package tree

// RedBlackTree X
type RedBlackTree struct {
	root *rbNode
	comp func(a, b interface{}) bool
}

type rbNode struct {
	size  int
	isRed bool
	value interface{}
	left  *rbNode
	right *rbNode
}

// NewRedBlackTree X
func NewRedBlackTree(comp func(a, b interface{}) bool) *RedBlackTree {
	t := &RedBlackTree{
		comp: comp,
	}
	return t
}

// Size X
func (t *RedBlackTree) Size() int {
	return t.root.size
}

// IsEmpty X
func (t *RedBlackTree) IsEmpty() bool {
	return t.root == nil
}

// Clear X
func (t *RedBlackTree) Clear() {
	t.root = nil
}

// Put X
func (t *RedBlackTree) Put(value interface{}) {
	t.root = t.put(value, t.root)
	t.root.isRed = false
}

func (t *RedBlackTree) put(value interface{}, node *rbNode) *rbNode {
	if node == nil {
		return &rbNode{
			size:  1,
			value: value,
		}
	}

	if t.comp(value, node.value) {
		node.left = t.put(value, node.left)
	} else if t.comp(node.value, value) {
		node.right = t.put(value, node.right)
	} else {
		node.value = value
	}

	return t.balance(node)
}

func (t *RedBlackTree) balance(node *rbNode) *rbNode {
	// colorL := t.isRed(node.left)
	// colorR := t.isRed(node.right)
	if !t.isRed(node.left) && t.isRed(node.right) {
		node = t.rotL(node)
	}

	if t.isRed(node.left) && t.isRed(node.left.left) {
		node = t.rotR(node)
	}

	if t.isRed(node.left) && t.isRed(node.right) {
		node = t.flip(node)
	}

	node.size = t.size(node.left) + t.size(node.right) + 1
	return node
}

// Has X
func (t *RedBlackTree) Has(value interface{}) bool {
	return t.has(value, t.root)
}

func (t *RedBlackTree) has(value interface{}, node *rbNode) bool {
	// for node != nil {
	// 	b1 := t.comp(value, node.value)
	// 	b2 := t.comp(node.value, value)
	// 	if !b1 && !b2 {
	// 		return true
	// 	}
	// 	if b1 {
	// 		node = node.left
	// 	} else if b2 {
	// 		node = node.right
	// 	}
	// }
	// return false

	if node == nil {
		return false
	}
	if t.comp(value, node.value) {
		return t.has(value, node.left)
	}
	if t.comp(node.value, value) {
		return t.has(value, node.right)
	}
	return true
}

// Min X
func (t *RedBlackTree) Min() interface{} {
	if t.root == nil {
		return nil
	}
	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node.value
}

// Max X
func (t *RedBlackTree) Max() interface{} {
	if t.root == nil {
		return nil
	}
	node := t.root
	for node.right != nil {
		node = node.right
	}
	return node.value
}

func (t *RedBlackTree) isRed(node *rbNode) bool {
	return node != nil && node.isRed
}

func (t *RedBlackTree) size(node *rbNode) int {
	if node == nil {
		return 0
	}
	return node.size
}

func (t *RedBlackTree) rotL(n *rbNode) *rbNode {
	x := n.right
	n.right = x.left
	x.left = n

	x.isRed = n.isRed
	n.isRed = true

	x.size = n.size
	n.size = t.size(n.left) + t.size(n.right) + 1

	return x
}

func (t *RedBlackTree) rotR(n *rbNode) *rbNode {
	x := n.left
	n.left = x.right
	x.right = n

	x.isRed = n.isRed
	n.isRed = true

	x.size = n.size
	n.size = t.size(n.left) + t.size(n.right) + 1

	return x
}

func (t *RedBlackTree) flip(node *rbNode) *rbNode {
	node.left.isRed = false
	node.right.isRed = false
	node.isRed = true
	return node
}
