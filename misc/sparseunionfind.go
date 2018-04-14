package misc

// SparseUnionFind sparse disjoint set structure
type SparseUnionFind struct {
	count  int
	parent map[int]int
}

// NewSparseUnionFind returns pointer to SparseUnionFind
func NewSparseUnionFind() *SparseUnionFind {
	u := &SparseUnionFind{
		parent: make(map[int]int),
	}
	return u
}

// Count returns number of independent sets
func (u *SparseUnionFind) Count() int {
	return u.count
}

// Size returns number of independent sets
func (u *UnionFind) Size(x int) int {
	return -u.parent[u.find(x)]
}

// Create creates new independent set
func (u *SparseUnionFind) Create(x int) {
	if !u.IsExists(x) {
		u.parent[x] = -1
		u.count++
	}
}

// Union unites two sets
func (u *SparseUnionFind) Union(x, y int) {
	u.Create(x)
	u.Create(y)
	xx, yy := u.find(x), u.find(y)
	if xx == yy {
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
func (u *SparseUnionFind) IsUnited(x, y int) bool {
	return u.IsExists(x) && u.IsExists(y) && u.find(x) == u.find(y)
}

// IsExists returns true if given element exists, false otherwise
func (u *SparseUnionFind) IsExists(x int) bool {
	_, ok := u.parent[x]
	return ok
}

func (u *SparseUnionFind) find(x int) int {
	for u.parent[x] >= 0 {
		x = u.parent[x]
	}
	return x
}
