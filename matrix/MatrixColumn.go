package matrix

// Alias so matrix columns are not confused for matrix rows.
type MatrixColumn = matrixSegment

// Create a new matrix column.
func NewMatrixColumn(n int) *MatrixColumn {
	return createSegment(n)
}

// Create a new matrix column with existing elements.
func NewMatrixColumnWithElements(elements []int64) *MatrixColumn {
	return createSegmentWithElements(elements)
}
