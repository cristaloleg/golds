package golds

// DisjointSet represents disjoint set (union-find) structure.
//
type DisjointSet struct {
	count  int
	parent []int // (u.parent >= 0) ? (parent) : (root's size)
}

// NewDisjointSet instantiates a new DisjointSet.
func NewDisjointSet(size int) *DisjointSet {
	u := &DisjointSet{
		count:  size,
		parent: make([]int, size),
	}
	for i := 0; i < size; i++ {
		u.parent[i] = -1
	}
	return u
}

// Count returns number of independent sets.
func (u *DisjointSet) Count() int {
	return u.count
}

// Size returns number of independent sets for a given node.
func (u *DisjointSet) Size(x int) int {
	return -u.parent[u.find(x)]
}

// Union unites two sets, return true if sets are already connected.
func (u *DisjointSet) Union(x, y int) bool {
	x, y = u.find(x), u.find(y)
	if x == y {
		return false
	}
	if u.parent[x] > u.parent[y] {
		x, y = y, x
	}
	u.parent[x] += u.parent[y]
	u.parent[y] = x
	u.count--
	return true
}

// IsUnited returns true if two sets are connected.
func (u *DisjointSet) IsUnited(x, y int) bool {
	return u.find(x) == u.find(y)
}

// find returns a parent of a x node.
func (u *DisjointSet) find(x int) int {
	for u.parent[x] >= 0 {
		x = u.parent[x]
	}
	return x
}
