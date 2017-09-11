package heap

// FibHeap Fibonacci Heap structure
type FibHeap struct {
	count int
	min   *fibHeapNode
	comp  func(a, b interface{}) bool
}

// fibHeapNode represents node of FibHeap
type fibHeapNode struct {
	value  interface{}
	marked bool
	degree int
	left   *fibHeapNode
	right  *fibHeapNode
	parent *fibHeapNode
	child  *fibHeapNode
}

// NewFibHeap returns a pointer to FibHeap
func NewFibHeap(comp func(interface{}, interface{}) bool) *FibHeap {
	h := &FibHeap{
		comp: comp,
	}
	return h
}

// Size returns size of a heap
func (h *FibHeap) Size() int {
	return h.count
}

// IsEmpty returns true if heap is empty
func (h *FibHeap) IsEmpty() bool {
	return h.count == 0
}

// Clear removes all elements from the FibHeap
func (h *FibHeap) Clear() {
	h.min = nil
	h.count = 0
}

// Build X
func (h *FibHeap) Build(values []interface{}) {
	for _, v := range values {
		h.Push(v)
	}
}

// Push adds value to the heap in O(1) time
func (h *FibHeap) Push(value interface{}) {
	n := &fibHeapNode{
		value: value,
	}
	h.count++
	h.insert(n)
}

// PushMany adds elements to the heap
func (h *FibHeap) PushMany(values ...interface{}) {
	for _, v := range values {
		h.Push(v)
	}
}

// Pop removes top element from the heap in O(log(N)) time
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
		h.min = minNode.right
		h.consolidate()
	}
	h.count--

	return minNode.value, true
}

// Top returns top element from the heap in O(1) time
func (h *FibHeap) Top() (value interface{}, ok bool) {
	if h.min == nil {
		return nil, false
	}
	return h.min.value, true
}

// insert appends new node to the root list
func (h *FibHeap) insert(x *fibHeapNode) {
	x.parent = nil
	x.marked = false

	if h.min == nil {
		h.min = x
		h.min.left = h.min
		h.min.right = h.min
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

// consolidate changes structure of the heap
func (h *FibHeap) consolidate() {
	nodes := make(map[int]*fibHeapNode)

	rootNodes := make([]*fibHeapNode, 0)
	rootNodes = append(rootNodes, h.min)
	for n := h.min.right; n != h.min; n = n.right {
		rootNodes = append(rootNodes, n)
	}

	for _, node := range rootNodes {
		d := node.degree
		for nodes[d] != nil {
			y := nodes[d]
			if h.comp(y.value, node.value) {
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

// link make y child of x
func (h *FibHeap) link(y, x *fibHeapNode) {
	y.left.right = y.right
	y.right.left = y.left

	if x.child == nil {
		y.left = y
		y.right = y
	} else {
		y.left = x.child.left
		y.right = x.child
		y.left.right = y
		y.right.left = y
	}
	x.child = y
	y.parent = x

	x.degree++
	y.marked = false
}
