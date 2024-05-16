package Set

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()

	// test for checking if set is empty
	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty, but it's not")
	}

	// test for length of an empty set
	if set.Len() != 0 {
		t.Errorf("Expected length of set to be 0, got %d", set.Len())
	}

	// test for adding elements
	set.Push(1)
	set.Push("hello")
	set.Push(3.14)

	// test for length after adding elements
	if set.Len() != 3 {
		t.Errorf("Expected length of set to be 3, got %d", set.Len())
	}

	// test for checking existence of added elements
	if !set.Find(1) || !set.Find("hello") || !set.Find(3.14) {
		t.Errorf("Expected all added elements to be found in the set")
	}

	// test for checking non-existence of an element
	if set.Find("world") {
		t.Errorf("Expected 'world' not to be found in the set, but it was found")
	}

	// test for adding duplicate element
	err := set.Push("hello")
	if err == nil {
		t.Errorf("Expected error when adding duplicate element, but no error occurred")
	}

	// test for removing an element
	err = set.Pop(1)
	if err != nil {
		t.Errorf("Unexpected error when removing element: %v", err)
	}

	// test for length after removing an element
	if set.Len() != 2 {
		t.Errorf("Expected length of set to be 2 after removing an element, got %d", set.Len())
	}

	// test for removing a non-existent element
	err = set.Pop("world")
	if err == nil {
		t.Errorf("Expected error when removing non-existent element, but no error occurred")
	}

	// test for checking non-existence of a removed element
	if set.Find(1) {
		t.Errorf("Expected 1 not to be found in the set after removal, but it was found")
	}
}
