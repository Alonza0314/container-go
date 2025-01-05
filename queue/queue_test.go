package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	// test for checking queue is empty
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}

	// test for pushing 2 items and the Len function
	q.Push(1)
	q.Push(2)
	if q.Len() != 2 {
		t.Errorf("Expected length of queue to be 2, got %d", q.Len())
	}

	// test for getting front
	front, err := q.Front()
	if err != nil {
		t.Errorf("Error while getting front element: %v", err)
	} else if front != 1 {
		t.Errorf("Expected front element to be 1, got %v", front)
	}

	// test for poping item and the Len function
	item, err := q.Pop()
	if err != nil {
		t.Errorf("Error while popping element: %v", err)
	} else if item != 1 {
		t.Errorf("Expected popped item to be 1, got %v", item)
	}
	if q.Len() != 1 {
		t.Errorf("Expected length of queue to be 1 after pop, got %d", q.Len())
	}

	// test for popping item and checking queue is empty
	q.Pop()
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after pop, but it's not")
	}

	// test for popping on empty queue
	_, err = q.Pop()
	if err == nil || err.Error() != "Queue is empty." {
		t.Errorf("Expected 'Queue is empty.' error, got %v", err)
	}
}