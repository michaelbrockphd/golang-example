package matrix

// Interface describing all methods a matrix segment object must implement.
type MatrixSegment interface {
	// Return the number of elements in the row/column.
	Length() int

	// Get the element in the row/column
	Element(i int) (int64, error)

	// Set the element in the row/column
	SetElement(i int, v int64) error

	// Return a string of the row/column.
	String() string
}
