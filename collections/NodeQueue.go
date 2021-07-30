package collections

import "errors"

// Simple queue implementation using single-linked nodes.
type NodeQueue struct {
	length int

	head *SingleLinkedNode

	tail *SingleLinkedNode
}

// Get the number of items in the queue.
func (nq *NodeQueue) Length() int {
	return nq.length
}

// Return the item at the head of the queue.
func (nq *NodeQueue) Head() interface{} {
	if nq.head != nil {
		return nq.head.value
	}

	return nil
}

// Return the item at the end of the queue.
func (nq *NodeQueue) Tail() interface{} {
	if nq.tail != nil {
		return nq.tail.value
	}

	return nil
}

// Add a item to the tail of the queue.
func (nq *NodeQueue) Enqueue(i interface{}) {
	newTail := NewSingleLinkedNode()

	newTail.value = i

	if nq.tail != nil {
		nq.tail.next = newTail

		nq.tail = newTail

	} else {
		nq.head = newTail
		nq.tail = newTail
	}

	nq.length += 1
}

// Remove the item at the head of the queue/
func (nq *NodeQueue) Dequeue() error {
	if nq.length < 1 {
		return errors.New("cannot dequeue from an empty queue")
	}

	oldHead := nq.head

	nq.head = oldHead.next

	oldHead.next = nil

	if nq.head == nil {
		nq.tail = nil
	}

	nq.length -= 1

	return nil
}

func NewNodeQueue() Queue {
	rtn := new(NodeQueue)

	return rtn
}
