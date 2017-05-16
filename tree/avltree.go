package tree

// AvlTree X
type AvlTree struct {
	root *avlNode
	comp func(a, b interface{}) bool
}

type avlNode struct {
	size   int
	height int
	left   *avlNode
	right  *avlNode
	value  interface{}
}

// NewAvlTree X
func NewAvlTree(comp func(a, b interface{}) bool) *AvlTree {
	t := &AvlTree{
		comp: comp,
	}
	return t
}

// Size X
func (t *AvlTree) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size
}

// IsEmpty X
func (t *AvlTree) IsEmpty() bool {
	return t.root == nil
}

// Clear X
func (t *AvlTree) Clear() {
	t.root = nil
}

// Put X
func (t *AvlTree) Put(value interface{}) {
	t.root = t.put(value, t.root)
}

func (t *AvlTree) put(value interface{}, node *avlNode) *avlNode {
	if node == nil {
		return &avlNode{
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

	node.size = 1 + t.size(node.left) + t.size(node.right)
	node.height = 1 + max(t.height(node.left), t.height(node.right))

	return t.balance(node)
}

func (t *AvlTree) size(node *avlNode) int {
	if node == nil {
		return 0
	}
	return node.size
}

func (t *AvlTree) height(node *avlNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (t *AvlTree) balance(node *avlNode) *avlNode {
	diff := t.balanceFactor(node)
	if diff < -1 {
		if t.balanceFactor(node.right) > 0 {
			node.right = t.rotR(node.right)
		}
		node = t.rotL(node)
	} else if diff > 1 {
		if t.balanceFactor(node.left) < 0 {
			node.left = t.rotL(node.left)
		}
		node = t.rotR(node)
	}
	return node
}

func (t *AvlTree) balanceFactor(node *avlNode) int {
	return t.height(node.left) - t.height(node.right)
}

func (t *AvlTree) rotL(node *avlNode) *avlNode {
	y := node.right
	node.right = y.left
	y.left = node
	y.size = node.size
	node.size = 1 + t.size(node.left) + t.size(node.right)
	node.height = 1 + max(t.height(node.left), t.height(node.right))
	y.height = 1 + max(t.height(y.left), t.height(y.right))
	return y
}

func (t *AvlTree) rotR(node *avlNode) *avlNode {
	y := node.left
	node.left = y.right
	y.right = node
	y.size = node.size
	node.size = 1 + t.size(node.left) + t.size(node.right)
	node.height = 1 + max(t.height(node.left), t.height(node.right))
	y.height = 1 + max(t.height(y.left), t.height(y.right))
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Has X
func (t *AvlTree) Has(value interface{}) bool {
	return t.has(value, t.root)
}

func (t *AvlTree) has(value interface{}, node *avlNode) bool {
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
func (t *AvlTree) Min() (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node.value, true
}

// Max X
func (t *AvlTree) Max() (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	node := t.root
	for node.right != nil {
		node = node.right
	}
	return node.value, true
}

// // PopMin X
// func (t *AvlTree) PopMin() (interface{},bool) {
// 	if t.root == nil {
// 		return nil, false
// 	}
// 	node := t.root
// 	for node.left != nil {
// 		node = node.left
// 	}
// 	t.del(node.value, node)
// 	return node.value, true
// }

// // PopMax X
// func (t *AvlTree) PopMax() (interface{},bool) {
// 	if t.root == nil {
// 		return nil, false
// 	}
// 	node := t.root
// 	for node.right != nil {
// 		node = node.right
// 	}
// 	t.del(node.value, node)
// 	return node.value, true
// }

// // Del X
// func (t *AvlTree) Del(value interface{}) {
// 	t.del(value, t.root)
// }

// func (t *AvlTree) del(value interface{}, node *avlNode) {
// 	b1 := t.comp(value, node.value)
// 	b2 := t.comp(node.value, value)

// 	if b1 {
// 		// node.left = t.del(value, node.left)
// 	} else if b2 {
// 		// node.right = t.del(value, node.right)
// 	} else {
// 		if node.left == nil {
// 			// return node.right
// 		} else if node.right == nil {
// 			// return node.left
// 		} else {
// 			// y := node
// 			// node = min(y.right);
// 			// node.right = deleteMin(y.right);
// 			// node.left = node.left
// 		}
// 	}
// 	node.size = 1 + t.size(node.left) + t.size(node.right)
// 	node.height = 1 + max(t.height(node.left), t.height(node.right))
// 	// return t.balance(node)
// }
