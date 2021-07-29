package collections

// Interface describing all methods a queue must implement.
type Queue interface {
	// Get the number of items in the queue.
	Length() int

	// Return the item at the head of the queue.
	Head() interface{}

	// Return the item at the end of the queue.
	Tail() interface{}

	// Add a item to the tail of the queue.
	Enqueue(i interface{})

	// Remove the item at the head of the queue/
	Dequeue() error
}
