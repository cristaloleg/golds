package interval

// SegmentTree X
type SegmentTree struct {
	size int
	data []int
}

// NewSegmentTree X
func NewSegmentTree(size int) *SegmentTree {
	t := &SegmentTree{
		size: size,
		data: make([]int, size*2),
	}
	return t
}

// Build X
func (t *SegmentTree) Build(values []int) {
	t.size = len(values)
	for i := 0; i < t.size; i++ {
		t.data[t.size+i] = values[i]
	}
	for i := t.size - 1; i > 0; i-- {
		t.data[i] = t.data[i<<1] + t.data[i<<1|1]
	}
}

// Modify X
func (t *SegmentTree) Modify(i int, value int) {
	i += t.size
	t.data[i] = value
	for ; i > 1; i >>= 1 {
		t.data[i>>1] = t.data[i] + t.data[i^1]
	}
}

// ModifyRange X
// func (t *SegmentTree) ModifyRange(i, j int, value int) {
// 	i += t.size
// 	j += t.size
// 	for i < j {
// 		if (i & 1) == 1 {
// 			t.data[i] += value
// 			i++
// 		}
// 		if (j & 1) == 1 {
// 			t.data[j] += value
// 			j--
// 		}
// 		i >>= 1
// 		j >>= 1
// 	}
// }

// Query X
// func (t *SegmentTree) Query(i int) int {
// 	var res int
// 	for i += t.size; i > 0; i >>= 1 {
// 		res = res + t.data[i]
// 	}
// 	return res
// }

// QueryRange X
func (t *SegmentTree) QueryRange(i, j int) int {
	var res int
	i += t.size
	j += t.size
	for i < j {
		if (i & 1) == 1 {
			res = res + t.data[i]
			i++
		}
		if (j & 1) == 1 {
			j--
			res = res + t.data[j]

		}
		i >>= 1
		j >>= 1
	}
	return res
}
