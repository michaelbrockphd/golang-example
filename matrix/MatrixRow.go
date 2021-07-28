package matrix

// Alias so matrix rows are not confused for matrix columns.
type MatrixRow = matrixSegment

// Create a new matrix row.
func NewMatrixRow(n int) *MatrixRow {
	return createSegment(n)
}

// Create a new matrix row with existing elements.
func NewMatrixRowWithElements(elements []int64) *MatrixRow {
	return createSegmentWithElements(elements)
}
