package maps_test

import "testing"
import . "github.com/cristaloleg/golds/maps"

func TestCounter(t *testing.T) {
	t.Parallel()

	c := NewCounter()
	if c == nil {
		t.Error("cannot instantiate a Counter")
	}

	if c.Size() != 0 {
		t.Error("must be empty")
	}

	c.Inc("key")
	c.Put("key", 2)

	if c.Size() != 1 || !c.Has("key") {
		t.Error("must be non-empty")
	}

	if value := c.Get("key"); value != 2 {
		t.Errorf("must be 2, but was %v", value)
	}

	c.Inc("key")
	if value := c.Get("key"); value != 3 {
		t.Errorf("must be 3, but was %v", value)
	}

	c.Dec("key")
	if value := c.Get("key"); value != 2 {
		t.Errorf("must be 2, but was %v", value)
	}

	c.Del("key")
	if value := c.Get("key"); value != -1 {
		t.Errorf("must be -1, but was %v", value)
	}

	c.Dec("key")
	if value := c.Get("key"); value != 0 {
		t.Errorf("must be 0, but was %v", value)
	}

	if keys := c.Keys(); len(keys) != 1 {
		t.Errorf("should have %v keys, got %v", 1, len(keys))
	}
}
