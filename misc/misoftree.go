package misc

// MisofTree represents Misof's Tree.
type MisofTree struct {
	data [15][1 << 15]int
}

// NewMisofTree instantiates a new MisofTree.
func NewMisofTree(size int) *MisofTree {
	t := &MisofTree{}
	return t
}

func (t *MisofTree) Inc(x int) {
	for i := 0; i < 15; i, x = i+1, x>>1 {
		t.data[i][x]++
	}
}

func (t *MisofTree) Dec(x int) {
	for i := 0; i < 15; i, x = i+1, x>>1 {
		t.data[i][x]--
	}
}

func (t *MisofTree) Kth(k int) int {
	res := 0
	for i := 15 - 1; i >= 0; i-- {
		res <<= 1
		if t.data[i][res] <= k {
			k -= t.data[i][res]
			res |= 1
		}
	}
	return res
}
