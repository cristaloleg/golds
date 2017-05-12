package misc

// UnionFind XXX
type UnionFind struct {
	count  int
	parent []int
	rank   []int
}

// NewUnionFind XXX
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

// Count XXX
func (u *UnionFind) Count() int {
	return u.count
}

// Union XXX
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

// IsUnited XXX
func (u *UnionFind) IsUnited(x, y int) bool {
	return u.find(x) == u.find(y)
}
func (u *UnionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return x
}
