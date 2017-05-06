package set

import "testing"

func TestUnion(t *testing.T) {
	a := NewHashSet()
	b := NewHashSet()

	for i := 0; i < 10; i++ {
		a.Put(i)
	}

	for i := 5; i < 15; i++ {
		b.Put(i)
	}

	res := Union(a, b)

	for i := 0; i < 15; i++ {
		if !res.Has(i) {
			t.Errorf("must contain %v", i)
		}
	}
}

func TestIntersection(t *testing.T) {
	a := NewHashSet()
	b := NewHashSet()

	for i := 0; i < 10; i++ {
		a.Put(i)
	}

	for i := 5; i < 15; i++ {
		b.Put(i)
	}

	res := Intersection(a, b)

	for i := 5; i < 10; i++ {
		if !res.Has(i) {
			t.Errorf("must contain %v", i)
		}
	}
}

func TestDifference(t *testing.T) {
	a := NewHashSet()
	b := NewHashSet()

	for i := 0; i < 10; i++ {
		a.Put(i)
	}

	for i := 5; i < 15; i++ {
		b.Put(i)
	}

	res := Difference(a, b)

	for i := 0; i < 5; i++ {
		if !res.Has(i) {
			t.Errorf("must contain %v", i)
		}
	}

	res = Difference(b, a)

	for i := 10; i < 15; i++ {
		if !res.Has(i) {
			t.Errorf("must contain %v", i)
		}
	}
}
