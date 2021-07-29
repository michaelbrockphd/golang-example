package collections

import (
	"errors"
)

// Stack implementation using a Golang slice.
//
// Conceptually, a stack should only allow a programmer to add or remove items
// at the head of the collection.
//
// Internally, this implementation adds and removes items to the end of slice
// to keep memory and CPU usage down (try and avoid shifting items, etc.)
type SliceStack struct {
	length int

	items []element
}

// Get the number of items in the stack.
func (s *SliceStack) Lenght() int {
	return s.length
}

// The top most item in the stack.
func (s *SliceStack) Top() interface{} {
	if s.length < 1 {
		return nil

	} else {
		return s.items[s.length-1].value
	}
}

// Add an new item to the stack.
func (s *SliceStack) Push(i interface{}) {
	e := element{i}

	s.items = append(s.items, e)

	s.length += 1
}

// Remove the top most item from the stack.
func (s *SliceStack) Pop() error {
	if s.length < 1 {
		return errors.New("cannot pop from an empty stack")
	}

	newSlice := s.items[:len(s.items)-1]

	s.items = newSlice

	s.length -= 1

	return nil
}

// Create a new stack with a slice as its internal structure
func NewSliceStack() Stack {
	rtn := new(SliceStack)

	rtn.items = make([]element, 0)

	return rtn
}
