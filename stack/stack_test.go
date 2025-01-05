package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	// test for getting empty stack's length
	if l := stack.Len(); l != 0 {
		t.Errorf("Expected length of stack to be 0, but got %d", l)
	}

	// test for pushing one item
	stack.Push(1)
	if l := stack.Len(); l != 1 {
		t.Errorf("Expected length of stack to be 1, but got %d", l)
	}

	// test for pushing multi items
	stack.Push(2)
	stack.Push(3)
	if l := stack.Len(); l != 3 {
		t.Errorf("Expected length of stack to be 3, but got %d", l)
	}

	// test for popping one item
	_, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error when popping one element, but got %v", err)
	}
	if l := stack.Len(); l != 2 {
		t.Errorf("Expected length of stack to be 2 after popping one element, but got %d", l)
	}

	// test for popping multi items
	_, err = stack.Pop()
	if err != nil {
		t.Errorf("Expected no error when popping one element, but got %v", err)
	}
	_, err = stack.Pop()
	if err != nil {
		t.Errorf("Expected no error when popping one element, but got %v", err)
	}
	if l := stack.Len(); l != 0 {
		t.Errorf("Expected length of stack to be 0 after popping all elements, but got %d", l)
	}

	// test for popping one item from empty stack
	_, err = stack.Pop()
	if err == nil || err.Error() != "stack is empty" {
		t.Errorf("Expected error 'stack is empty', but got %v", err)
	}

	// test for getting the top item from empty stack
	top, err := stack.Top()
	if err == nil || top != nil {
		t.Errorf("Expected top element to be nil in an empty stack, but got %v", top)
	}

	// test for getting the top item
	stack.Push(10)
	top, err = stack.Top()
	if err != nil || top != 10 {
		t.Errorf("Expected top element to be 10, but got %v", top)
	}

	// test for checking stack is not emoty
	if isEmpty := stack.IsEmpty(); isEmpty {
		t.Errorf("Expected stack to be not empty, but it's empty")
	}
}
