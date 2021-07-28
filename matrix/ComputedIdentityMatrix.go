package matrix

// Private details of a computed identity matrix.
//
// Rather than store a whole matrix, only the dimensions and identity values are
// kept.
type computedIdentityMatrix struct {
	n   int
	val int64
}

// Get the number of rows in the matrix.
func (m *computedIdentityMatrix) Rows() int {
	return m.n
}

// Get the number of columns in the matrix.
func (m *computedIdentityMatrix) Columns() int {
	return m.n
}

// Get the value of the specified element.
func (m *computedIdentityMatrix) Element(r int, c int) int64 {
	rtn := int64(0)

	if r == c {
		rtn = m.val
	}

	return rtn
}

// Set the value of the specified element.
func (m *computedIdentityMatrix) SetElement(r int, c int, v int64) {
	// Throw an error!
}

// Get a given row form the matrix.
func (m *computedIdentityMatrix) Row(r int) *MatrixRow {
	row := make([]int64, m.n)

	row[r] = m.val

	return NewMatrixRowWithElements(row)
}

// Get a given column from the matrix.
func (m *computedIdentityMatrix) Column(c int) *MatrixColumn {
	col := make([]int64, m.n)

	col[c] = m.val

	return NewMatrixColumnWithElements(col)
}

// Create a new computed identity matrix for the given value and dimensions.
func NewComputedIdentityMatrix(n int, val int64) Matrix {
	matrix := new(computedIdentityMatrix)

	matrix.n = n

	matrix.val = val

	return matrix
}