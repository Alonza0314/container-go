package Heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	// use minHeap to test, so child value should be larger than parent value
	cmpFunc := func(child, parent interface{}) bool {
		return child.(int) > parent.(int)
	}

	heap := NewHeap(cmpFunc)

	// test for checking heap is empty
	if !heap.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}

	// Test for pushing elements into the heap
	heap.Push(3)
	if heap.Len() != 1 {
		t.Errorf("Expected heap length to be 1 after pushing an element, got %d", heap.Len())
	}

	heap.Push(5)
	if heap.Len() != 2 {
		t.Errorf("Expected heap length to be 2 after pushing an element, got %d", heap.Len())
	}

	heap.Push(2)
	if heap.Len() != 3 {
		t.Errorf("Expected heap length to be 3 after pushing an element, got %d", heap.Len())
	}

	// Test for popping elements from the heap
	top, err := heap.Pop()
	if err != nil {
		t.Errorf("Unexpected error while popping from heap: %v", err)
	}
	if top != 2 {
		t.Errorf("Expected top element to be 2, got %v", top)
	}

	top, err = heap.Pop()
	if err != nil {
		t.Errorf("Unexpected error while popping from heap: %v", err)
	}
	if top != 3 {
		t.Errorf("Expected top element to be 3, got %v", top)
	}

	top, err = heap.Pop()
	if err != nil {
		t.Errorf("Unexpected error while popping from heap: %v", err)
	}
	if top != 5 {
		t.Errorf("Expected top element to be 5 got %v", top)
	}

	// Test for popping from an empty heap
	_, err = heap.Pop()
	if err == nil {
		t.Error("Expected an error while popping from an empty heap, but got none")
	}
}

func TestMaxHeap(t *testing.T) {
	// use maxHeap to test, so child value should be larger than parent value
	cmpFunc := func(child, parent interface{}) bool {
		return child.(int) < parent.(int)
	}

	heap := NewHeap(cmpFunc)

	// test for checking heap is empty
	if !heap.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}

	// Test for pushing elements into the heap
	heap.Push(3)
	if heap.Len() != 1 {
		t.Errorf("Expected heap length to be 1 after pushing an element, got %d", heap.Len())
	}

	heap.Push(5)
	if heap.Len() != 2 {
		t.Errorf("Expected heap length to be 2 after pushing an element, got %d", heap.Len())
	}

	heap.Push(2)
	if heap.Len() != 3 {
		t.Errorf("Expected heap length to be 3 after pushing an element, got %d", heap.Len())
	}

	// Test for popping elements from the heap
	top, err := heap.Pop()
	if err != nil {
		t.Errorf("Unexpected error while popping from heap: %v", err)
	}
	if top != 5 {
		t.Errorf("Expected top element to be 5, got %v", top)
	}

	top, err = heap.Pop()
	if err != nil {
		t.Errorf("Unexpected error while popping from heap: %v", err)
	}
	if top != 3 {
		t.Errorf("Expected top element to be 3, got %v", top)
	}

	top, err = heap.Pop()
	if err != nil {
		t.Errorf("Unexpected error while popping from heap: %v", err)
	}
	if top != 2 {
		t.Errorf("Expected top element to be 2 got %v", top)
	}

	// Test for popping from an empty heap
	_, err = heap.Pop()
	if err == nil {
		t.Error("Expected an error while popping from an empty heap, but got none")
	}
}
