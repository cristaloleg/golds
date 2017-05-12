package heap

// FibHeap XXX
type FibHeap struct {
	count int
	min   *fibHeapNode
	comp  func(a, b interface{}) bool
}

type fibHeapNode struct {
	value  interface{}
	marked bool
	degree int
	left   *fibHeapNode
	right  *fibHeapNode
	parent *fibHeapNode
	child  *fibHeapNode
}

// NewFibHeap XXX
func NewFibHeap(comp func(interface{}, interface{}) bool) *FibHeap {
	h := &FibHeap{
		comp: comp,
	}
	return h
}

// Size XXX
func (h *FibHeap) Size() int {
	return h.count
}

// IsEmpty XXX
func (h *FibHeap) IsEmpty() bool {
	return h.count == 0
}

// Push XXX
func (h *FibHeap) Push(value interface{}) {
	n := &fibHeapNode{
		value: value,
	}
	h.count++
	h.insert(n)
}

// Pop XXX
func (h *FibHeap) Pop() (value interface{}, ok bool) {
	if h.min == nil {
		return nil, false
	}
	minNode := h.min

	if minNode.child != nil {
		child := minNode.child
		for minNode == child.parent {
			nextChild := child.right
			h.insert(child)
			child = nextChild
		}
	}

	minNode.left.right = minNode.right
	minNode.right.left = minNode.left

	if minNode.right == minNode {
		h.min = nil
	} else {
		h.min = h.min.right
		h.consolidate()
	}
	h.count--

	return minNode.value, true
}

// Top XXX
func (h *FibHeap) Top() (value interface{}, ok bool) {
	if h.min == nil {
		return nil, false
	}
	return h.min.value, true
}

func (h *FibHeap) insert(x *fibHeapNode) {
	x.parent = nil
	x.marked = false

	if h.min == nil {
		h.min = x
		h.min.right = h.min
		h.min.left = h.min
		return
	}

	x.left = h.min.left
	x.right = h.min
	x.left.right = x
	x.right.left = x

	if h.comp(x.value, h.min.value) {
		h.min = x
	}
}

func (h *FibHeap) consolidate() {
	nodes := make([]*fibHeapNode, h.count+1)

	rootNodes := make([]*fibHeapNode, 0)
	rootNodes = append(rootNodes, h.min)
	for n := h.min.right; n != h.min; n = n.right {
		rootNodes = append(rootNodes, n)
	}

	for _, node := range rootNodes {
		if node.parent != nil {
			continue
		}

		d := node.degree
		for nodes[d] != nil {
			y := nodes[d]
			if !h.comp(node.value, y.value) {
				node, y = y, node
			}
			h.link(y, node)
			nodes[d] = nil
			d++
		}
		nodes[d] = node
	}

	h.min = nil

	for _, node := range nodes {
		if node != nil {
			h.insert(node)
		}
	}
}

func (h *FibHeap) link(y, x *fibHeapNode) {
	y.left.right = y.right
	y.right.left = y.left

	if x.child == nil {
		y.right = y
		y.left = y
	} else {
		y.left = x.child.left
		y.right = x.child
		y.right.left = y
		y.left.right = y
	}
	x.child = y
	y.parent = x

	x.degree++
	y.marked = false
}

func (h *FibHeap) cut(x, y *fibHeapNode) {
	x.left.right = x.right
	x.right.left = x.left

	if x.right == x {
		y.child = nil
	} else {
		y.child = x.right
	}

	y.degree--
	h.insert(x)
}

func (h *FibHeap) cascadingCut(x *fibHeapNode) {
	parent := x.parent
	if parent == nil {
		return
	}
	if !parent.marked {
		parent.marked = true
		return
	}
	h.cut(x, parent)
	h.cascadingCut(parent)
}
