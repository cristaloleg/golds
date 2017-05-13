package misc

// UnionFind disjoint set structure
type UnionFind struct {
	count  int
	parent []int
	rank   []int
}

// NewUnionFind returns pointer to UnionFind
func NewUnionFind(size int) *UnionFind {
	u := &UnionFind{
		count:  size,
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := 0; i < size; i++ {
		u.parent[i] = i
	}
	return u
}

// Count returns number of independent sets
func (u *UnionFind) Count() int {
	return u.count
}

// Union unites two sets
func (u *UnionFind) Union(x, y int) {
	x, y = u.find(x), u.find(y)
	if x == y {
		return
	}
	if u.rank[x] < u.rank[y] {
		u.parent[x] = y
		u.rank[y] += u.rank[x]
	} else {
		u.parent[y] = x
		u.rank[x] += u.rank[y]
	}
	u.count--
}

// IsUnited returns true if two sets are connected
func (u *UnionFind) IsUnited(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *UnionFind) find(x int) int {
	parent := u.parent[x]
	if parent != x {
		parent = u.find(parent)
		u.parent[x] = parent
	}
	return parent
}
