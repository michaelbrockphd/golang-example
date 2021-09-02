package matrix

// Interface describing all methods a matrix object must implement.
type Matrix interface {
	// Get the number of rows in the matrix.
	Rows() int

	// Get the number of columns in the matrix.
	Columns() int

	// Get the value of the specified element.
	Element(r int, c int) (int64, error)

	// Set the value of the specified element.
	SetElement(r int, c int, v int64) error

	// Get a given row form the matrix.
	Row(r int) (MatrixSegment, error)

	// Get a given column from the matrix.
	Column(c int) (MatrixSegment, error)
}
