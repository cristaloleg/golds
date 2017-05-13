package misc

// SparseUnionFind sparse disjoint set structure
type SparseUnionFind struct {
	count  int
	parent map[int]int
	rank   map[int]int
}

// NewSparseUnionFind returns pointer to SparseUnionFind
func NewSparseUnionFind() *SparseUnionFind {
	u := &SparseUnionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
	return u
}

// Count returns number of independent sets
func (u *SparseUnionFind) Count() int {
	return u.count
}

// Create creates new independent set
func (u *SparseUnionFind) Create(x int) {
	u.parent[x] = x
	u.rank[x] = 0
	u.count++
}

// Union unites two sets
func (u *SparseUnionFind) Union(x, y int) {
	xx, yy := u.find(x), u.find(y)
	if xx != -1 && yy != -1 && xx == yy {
		return
	}
	if xx == -1 {
		u.Create(x)
	}
	if yy == -1 {
		u.Create(y)
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
func (u *SparseUnionFind) IsUnited(x, y int) bool {
	x, y = u.find(x), u.find(y)
	return x != -1 && y != -1 && x == y
}

func (u *SparseUnionFind) find(x int) int {
	value, ok := u.parent[x]
	if !ok {
		return -1
	}
	if value != x {
		u.parent[x] = u.find(value)
	}
	return value
}
