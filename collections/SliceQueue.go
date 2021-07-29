package collections

import "errors"

type SliceQueue struct {
	length int

	items []element
}

// Get the number of items in the queue.
func (q *SliceQueue) Length() int {
	return q.length
}

// Return the item at the head of the queue.
func (q *SliceQueue) Head() interface{} {
	if q.length < 1 {
		return nil

	} else {
		return q.items[0].value
	}
}

// Return the item at the end of the queue.
func (q *SliceQueue) Tail() interface{} {
	if q.length < 1 {
		return nil

	} else {
		return q.items[q.length-1].value
	}
}

// Add a item to the tail of the queue.
func (q *SliceQueue) Enqueue(i interface{}) {
	e := element{i}

	q.items = append(q.items, e)

	q.length += 1
}

// Remove the item at the head of the queue/
func (q *SliceQueue) Dequeue() error {
	if q.length < 1 {
		return errors.New("cannot dequeue from an empty queue")
	}

	q.items = q.items[1:q.length]

	q.length -= 1

	return nil
}
