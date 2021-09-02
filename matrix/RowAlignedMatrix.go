package matrix

// Adapter to create a row aligned matrix
type RowAlignedMatrix struct {
	mtx *linearMatrix
}

// Get the number of rows in the matrix.
func (m *RowAlignedMatrix) Rows() int {
	return m.mtx.count
}

// Get the number of columns in the matrix.
func (m *RowAlignedMatrix) Columns() int {
	return m.mtx.size
}

// Get the value of the specified element.
func (m *RowAlignedMatrix) Element(r int, c int) (int64, error) {
	rtn := int64(0)
	var err error = nil

	if r < 0 || c < 0 || m.mtx.count <= r || m.mtx.size <= c {
		err = NewMatrixIndexError(r, c)

	} else {
		rtn = m.mtx.readElement(r, c)
	}

	return rtn, err
}

// Set the value of the specified element.
func (m *RowAlignedMatrix) SetElement(r int, c int, v int64) error {
	var err error = nil

	if r < 0 || c < 0 || m.mtx.count <= r || m.mtx.size <= c {
		err = NewMatrixIndexError(r, c)

	} else {
		m.mtx.writeElement(r, c, v)
	}

	return err
}

// Get a given row form the matrix.
func (m *RowAlignedMatrix) Row(r int) (MatrixSegment, error) {
	var row MatrixSegment = nil
	var err error = nil

	if r < 0 || m.mtx.count <= r {
		err = NewMatrixSegmentIndexError(r)

	} else {
		sec := m.mtx.section(r)

		row = createSegmentWithElements(sec)
	}

	return row, err
}

// Get a given column from the matrix.
func (m *RowAlignedMatrix) Column(c int) (MatrixSegment, error) {
	var col MatrixSegment = nil
	var err error = nil

	if c < 0 || m.mtx.count <= c {
		err = NewMatrixSegmentIndexError(c)

	} else {
		sec := m.mtx.crossSection(c)

		col = createSegmentWithElements(sec)
	}

	return col, err
}

// Create a new row aligned matrix
//
// In memory, the matrix is kept as a sequence of rows.
func NewRowAlignedMatrix(rows int, columns int) Matrix {
	m := new(RowAlignedMatrix)

	m.mtx = createLinearMatrix(rows, columns)

	return m
}
