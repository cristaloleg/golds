package misc

// UnionFind disjoint set structure
type UnionFind struct {
	count  int
	parent []int
}

// NewUnionFind returns pointer to UnionFind
func NewUnionFind(size int) *UnionFind {
	u := &UnionFind{
		count:  size,
		parent: make([]int, size),
	}
	for i := 0; i < size; i++ {
		u.parent[i] = -1
	}
	return u
}

// Count returns number of independent sets
func (u *UnionFind) Count() int {
	return u.count
}

// Size returns number of independent sets
func (u *UnionFind) Size(x int) int {
	return -u.parent[x]
}

// Union unites two sets
func (u *UnionFind) Union(x, y int) {
	x, y = u.find(x), u.find(y)
	if x == y {
		return
	}
	if u.parent[x] > u.parent[y] {
		x, y = y, x
	}
	u.parent[x] += u.parent[y]
	u.parent[y] = x
	u.count--
}

// IsUnited returns true if two sets are connected
func (u *UnionFind) IsUnited(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *UnionFind) find(x int) int {
	for u.parent[x] >= 0 {
		x = u.parent[x]
	}
	return x
}
