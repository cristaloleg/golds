package tree

// AvlTree represents AVL tree
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

// NewAvlTree returns a pointer to the AvlTree
func NewAvlTree(comp func(a, b interface{}) bool) *AvlTree {
	t := &AvlTree{
		comp: comp,
	}
	return t
}

// Size returns number of elements in the tree
func (t *AvlTree) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size
}

// IsEmpty true if tree is empty
func (t *AvlTree) IsEmpty() bool {
	return t.root == nil
}

// Clear removes all elements from the tree
func (t *AvlTree) Clear() {
	t.root = nil
}

// Put adds element to the tree
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

// Has returns true if tree contains element
func (t *AvlTree) Has(value interface{}) bool {
	return t.has(value, t.root)
}

func (t *AvlTree) has(value interface{}, node *avlNode) bool {
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

// Min returns min element in O(log(N)) time
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

// Max returns max element in O(log(N)) time
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
