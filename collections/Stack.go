package collections

// Interface describing all methods a stack must implement.
type Stack interface {
	// Get the number of items in the stack.
	Lenght() int

	// The top most item in the stack.
	Top() interface{}

	// Add an new item to the stack.
	Push(i interface{})

	// Remove the top most item from the stack.
	Pop() error
}
