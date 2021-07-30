package collections

import (
	"errors"
)

// Simple stack implementation using single-linked nodes.
type NodeStack struct {
	length int

	topNode *SingleLinkedNode
}

// Get the number of items in the stack.
func (ns *NodeStack) Length() int {
	return ns.length
}

// The top most item in the stack.
func (ns *NodeStack) Top() interface{} {
	if ns.topNode != nil {
		return ns.topNode.value
	}

	return nil
}

// Add an new item to the stack.
func (ns *NodeStack) Push(i interface{}) {
	newTop := NewSingleLinkedNode()

	newTop.value = i

	newTop.next = ns.topNode

	ns.topNode = newTop

	ns.length += 1
}

// Remove the top most item from the stack.
func (ns *NodeStack) Pop() error {
	if ns.length < 1 {
		return errors.New("cannot pop from an empty stack")
	}

	oldTop := ns.topNode

	ns.topNode = oldTop.next

	oldTop.next = nil

	ns.length -= 1

	return nil
}

// Create a new node-backed stack
func NewNodeStack() Stack {
	rtn := new(NodeStack)

	rtn.length = 0

	rtn.topNode = nil

	return rtn
}
