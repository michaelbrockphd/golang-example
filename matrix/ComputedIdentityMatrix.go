package matrix

import (
	"errors"
)

// Private details of a computed identity matrix.
//
// Rather than store a whole matrix, only the dimensions and identity value are
// kept.
type ComputedIdentityMatrix struct {
	// The number of rows and columns in the matrix (they are the same)
	n int

	// The identity value (what diagonally appears on the matrix)
	val int64
}

// Get the number of rows in the matrix.
func (m *ComputedIdentityMatrix) Rows() int {
	return m.n
}

// Get the number of columns in the matrix.
func (m *ComputedIdentityMatrix) Columns() int {
	return m.n
}

// Get the value of the specified element.
func (m *ComputedIdentityMatrix) Element(r int, c int) (int64, error) {
	rtn := int64(0)
	var err error = nil

	if r < 0 || c < 0 || m.n <= r || m.n <= c {
		err = NewMatrixIndexError(r, c)

	} else {
		if r == c {
			rtn = m.val
		}
	}

	return rtn, err
}

// Set the value of the specified element.
func (m *ComputedIdentityMatrix) SetElement(r int, c int, v int64) error {
	return errors.New("Cannot modify a computed matrix segment.")
}

// Get a given row form the matrix.
func (m *ComputedIdentityMatrix) Row(r int) (MatrixSegment, error) {
	var rtn MatrixSegment = nil
	var err error = nil

	if r < 0 || m.n <= r {
		err = NewMatrixSegmentIndexError(r)

	} else {
		rtn = createComputedMatrixSegment(m.n, r, m.val)
	}

	return rtn, err
}

// Get a given column from the matrix.
func (m *ComputedIdentityMatrix) Column(c int) (MatrixSegment, error) {
	var rtn MatrixSegment = nil
	var err error = nil

	if c < 0 || m.n <= c {
		err = NewMatrixSegmentIndexError(c)

	} else {
		rtn = createComputedMatrixSegment(m.n, c, m.val)
	}

	return rtn, err
}

// Create a new computed identity matrix for the given value and dimensions.
func NewComputedIdentityMatrix(n int, val int64) Matrix {
	matrix := new(ComputedIdentityMatrix)

	matrix.n = n

	matrix.val = val

	return matrix
}
