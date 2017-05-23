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
	child  [2]*avlNode
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

	return t.balance(node.update())
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

func (t *AvlTree) balance2(node *avlNode) *avlNode {
	diff := t.balanceFactor(node)
	if diff == 0 {
		return node
	}

	idx, flag := bool2int(diff < -1), diff > 0

	if t.balanceFactor(node.child[idx]) != 0 {
		node.child[idx] = t.rot(flag, node.child[idx])
	}
	return t.rot(!flag, node)
}

func (t *AvlTree) balanceFactor(node *avlNode) int {
	return node.HeightOf(node.left) - node.HeightOf(node.right)
}

func (t *AvlTree) rot(isLeft bool, node *avlNode) *avlNode {
	idx, inv := bool2int(isLeft), bool2int(!isLeft)
	y := node.child[inv]
	node.child[inv] = y.child[idx]
	y.child[idx] = node
	y.size = node.size

	node.update()
	return y.update()
}

func (t *AvlTree) rotL(node *avlNode) *avlNode {
	y := node.right
	node.right = y.left
	y.left = node
	y.size = node.size

	node.update()
	return y.update()
}

func (t *AvlTree) rotR(node *avlNode) *avlNode {
	y := node.left
	node.left = y.right
	y.right = node
	y.size = node.size

	node.update()
	return y.update()
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
	for node != nil {
		if t.comp(value, node.value) {
			node = node.left
		} else {
			if t.comp(node.value, value) {
				node = node.right
			} else {
				return true
			}
		}
	}
	return false
}

// Min returns min element in O(log(N)) time
func (t *AvlTree) Min() (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.min(t.root).value, true
}

// Max returns max element in O(log(N)) time
func (t *AvlTree) Max() (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.max(t.root).value, true
}

// Del X
func (t *AvlTree) Del(value interface{}) {
	t.del(value, t.root)
}

func (t *AvlTree) del(value interface{}, node *avlNode) *avlNode {
	b2 := t.comp(node.value, value)

	if t.comp(value, node.value) {
		node.left = t.del(value, node.left)
	} else if b2 {
		node.right = t.del(value, node.right)
	} else {
		// if node.left == nil {
		// 	return node.right
		// } else if node.right == nil {
		// 	return node.left
		// } else {
		// 	y := node
		// 	node = t.min(y.right)
		// 	node.right = t.deleteMin(y.right)
		// }

		if node.left != nil && node.right != nil {
			// node to delete found with both children;
			// replace values with smallest node of the right sub-tree
			rightMinNode := t.min(node.right)
			// node.key = rightMinNode.key
			node.value = rightMinNode.value
			// delete smallest node that we replaced
			node.right = t.del(value, node.right)
		} else if node.left != nil {
			// node only has left child
			node = node.left
		} else if node.right != nil {
			// node only has right child
			node = node.right
		} else {
			// node has no children
			node = nil
			return node
		}
	}
	return t.balance(node.update())
}

func (t *AvlTree) min(node *avlNode) *avlNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *AvlTree) max(node *avlNode) *avlNode {
	for node.right != nil {
		node = node.right
	}
	return node
}

func (t *AvlTree) deleteMin(node *avlNode) *avlNode {
	return node
}

///////////

func (n *avlNode) update() *avlNode {
	n.size = 1 + n.Size()
	n.height = 1 + n.Height()
	return n
}

func (n *avlNode) Size() int {
	val := 0
	if n.left != nil {
		val += n.left.size
	}
	if n.right != nil {
		val += n.right.size
	}
	n.size = val
	return val
}

func (n *avlNode) Height() int {
	val := -1
	if n.left != nil {
		val = n.left.height
	}
	if n.right != nil && n.right.size > val {
		val = n.right.height
	}
	n.height = val
	return val
}

func (n *avlNode) HeightOf(node *avlNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

func bool2int(a bool) int {
	x := 0
	if a {
		x = 1
	}
	return x & 1
}
