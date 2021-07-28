package matrix

type MatrixSegment interface {
	// Return the number of elements in the row/column.
	Length() int

	// Get the element in the row/column
	Element(i int) int64

	// Set the element in the row/column
	SetElement(i int, v int64)

	// Return a string of the row/column.
	String() string
}
