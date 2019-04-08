package misc

// LinkCutTree ...
type LinkCutTree struct {
	nodes []*linkCutNode
}

// NewLinkCutTree ...
func NewLinkCutTree(size int) *LinkCutTree {
	t := &LinkCutTree{
		nodes: make([]*linkCutNode, size),
	}
	for i := range t.nodes {
		t.nodes[i].id = i
		t.nodes[i].update()
	}
	return t
}

// Link ...
func (t *LinkCutTree) Link(i, j int) {
	a := t.nodes[i]
	b := t.nodes[j]
	a.link(b)
}

// Cut ...
func (t *LinkCutTree) Cut(i int) {
	t.nodes[i].cut()
}

// IsConnected ...
func (t *LinkCutTree) IsConnected(i, j int) bool {
	a := t.nodes[i]
	b := t.nodes[j]
	return a.id == b.id
}

// Root ...
func (t *LinkCutTree) Root(i int) int {
	return t.nodes[i].root().id
}

// Depth ...
func (t *LinkCutTree) Depth(i int) int {
	return t.nodes[i].depth()
}

// LCA ...
func (t *LinkCutTree) LCA(i, j int) int {
	a := t.nodes[i]
	b := t.nodes[j]
	return a.lca(b).id
}

type linkCutNode struct {
	id          int
	size        int
	parent      *linkCutNode
	pathParent  *linkCutNode
	left, right *linkCutNode
}

func newNode() *linkCutNode {
	n := &linkCutNode{}
	return n
}

func (x *linkCutNode) update() {
	x.size = 1
	if x.left != nil {
		x.size += x.left.size
	}
	if x.right != nil {
		x.size += x.right.size
	}
}

func (x *linkCutNode) rotateRight() {
	y := x.parent
	z := y.parent

	y.left = x.right
	if y.left != nil {
		y.left.parent = y
	}

	x.right = y
	y.parent = x

	x.parent = z
	if x.parent != nil {
		if y == z.left {
			z.left = x
		} else {
			z.right = x
		}
	}
	x.pathParent = y.pathParent
	y.pathParent = nil
	y.update()
}

func (x *linkCutNode) rotateLeft() {
	y := x.parent
	z := y.parent

	y.right = x.left
	if y.right != nil {
		y.right.parent = y
	}
	x.left = y
	y.parent = x

	x.parent = z
	if x.parent != nil {
		if y == z.left {
			z.left = x
		} else {
			z.right = x
		}
	}
	x.pathParent = y.pathParent
	y.pathParent = nil
	y.update()
}

func (x *linkCutNode) splay() {
	for x.parent != nil {
		y := x.parent
		if y.parent == nil {
			if x == y.left {
				x.rotateRight()
			} else {
				x.rotateLeft()
			}
		} else {
			z := y.parent
			if y == z.left {
				if x == y.left {
					y.rotateRight()
					x.rotateRight()
				} else {
					x.rotateLeft()
					x.rotateRight()
				}
			} else {
				if x == y.right {
					y.rotateLeft()
					x.rotateLeft()
				} else {
					x.rotateRight()
					x.rotateLeft()
				}
			}
		}
	}
	x.update()
}

func (x *linkCutNode) access() *linkCutNode {
	x.splay()
	if x.right != nil {
		x.right.pathParent = x
		x.right.parent = nil
		x.right = nil
		x.update()
	}

	last := x
	for x.pathParent != nil {
		y := x.pathParent
		last = y
		y.splay()
		if y.right != nil {
			y.right.pathParent = y
			y.right.parent = nil
		}

		y.right = x
		x.parent = y
		x.pathParent = nil
		y.update()
		x.splay()
	}
	return last
}

func (x *linkCutNode) root() *linkCutNode {
	x.access()
	for x.left != nil {
		x = x.left
	}
	x.splay()
	return x
}

func (x *linkCutNode) cut() {
	x.access()
	x.left.parent = nil
	x.left = nil
	x.update()
}

func (x *linkCutNode) link(y *linkCutNode) {
	x.access()
	y.access()
	x.left = y
	y.parent = x
	x.update()
}

func (x *linkCutNode) lca(y *linkCutNode) *linkCutNode {
	x.access()
	return y.access()
}

func (x *linkCutNode) depth() int {
	x.access()
	return x.size - 1
}
