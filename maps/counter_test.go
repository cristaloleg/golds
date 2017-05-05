package maps

import "testing"

func TestCounter(t *testing.T) {
	c := NewCounter()
	if c == nil {
		t.Error("cannot instantiate a Counter")
	}

	if c.Size() != 0 {
		t.Error("must be empty")
	}

	c.Put("key", 2)

	if c.Size() != 1 {
		t.Error("must be empty")
	}

	v := c.Get("key")

	if v != 2 {
		t.Errorf("must be 2, but was %v", v)
	}
}
