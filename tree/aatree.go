package tree

import "github.com/cristaloleg/golds"

// AATree ...
type AATree struct {
	root *aaNode
}

type aaNode struct {
	level int
	left  *aaNode
	right *aaNode
	value golds.Comparable
}

// var nullNode *aaNode

// func init() {
// 	nullNode = new(aaNode)
// 	nullNode.level = 0
// 	nullNode.left = nullNode
// 	nullNode.right = nullNode
// }

// NewAATree ...
func NewAATree() *AATree {
	t := &AATree{}
	return t
}

// Size returns number of elements in the tree
func (t *AATree) Size() int {
	if t.root == nil {
		return 0
	}
	// return t.root
	return 0
}

// IsEmpty true if tree is empty
func (t *AATree) IsEmpty() bool {
	return t.root == nil
}

// Clear removes all elements from the tree
func (t *AATree) Clear() {
	t.root = nil
}

// Put ...
func (t *AATree) Put(value golds.Comparable) {
	t.put(value, t.root)
}

func (t *AATree) put(value golds.Comparable, node *aaNode) {
	if node == nil {
		node = &aaNode{0, nil, nil, value}
	} else if value.Less(node.value) {
		t.put(value, node.left)
	} else if node.value.Less(value) {
		t.put(value, node.right)
	} else {
		return
	}

	t.skew(node)
	t.split(node)
}

func (t *AATree) skew(node *aaNode) {
	if node.left.level == node.level {
		t.rotateWithLeftChild(node)
	}
}

func (t *AATree) split(node *aaNode) {
	if node.right.right.level == node.level {
		t.rotateWithRightChild(node)
		node.level++
	}
}

func (t *AATree) rotateWithLeftChild(node *aaNode) {
	tmp := node.left
	node.left = tmp.right
	tmp.right = node
	node = tmp
}

func (t *AATree) rotateWithRightChild(node *aaNode) {
	tmp := node.right
	node.right = tmp.left
	tmp.left = node
	node = tmp
}

// Has ...
func (t *AATree) Has(value golds.Comparable) bool {
	node := t.root
	for node != nil {
		if value.Less(node.value) {
			node = node.left
		} else {
			if node.value.Less(value) {
				node = node.right
			} else {
				return true
			}
		}
	}
	return false
}

// Min ...
func (t *AATree) Min() (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node.value, true
}

// Max ...
func (t *AATree) Max() (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	node := t.root
	for node.right != nil {
		node = node.right
	}
	return node.value, true
}
