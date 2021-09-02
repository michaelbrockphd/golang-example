package matrix

// Adapter to create a column aligned matrix
type ColumnAlignedMatrix struct {
	mtx *linearMatrix
}

// Get the number of rows in the matrix.
func (m *ColumnAlignedMatrix) Rows() int {
	return m.mtx.size
}

// Get the number of columns in the matrix.
func (m *ColumnAlignedMatrix) Columns() int {
	return m.mtx.count
}

// Get the value of the specified element.
func (m *ColumnAlignedMatrix) Element(r int, c int) (int64, error) {
	rtn := int64(0)
	var err error = nil

	if r < 0 || c < 0 || m.mtx.count <= c || m.mtx.size <= r {
		err = NewMatrixIndexError(r, c)

	} else {
		rtn = m.mtx.readElement(c, r)
	}

	return rtn, err
}

// Set the value of the specified element.
func (m *ColumnAlignedMatrix) SetElement(r int, c int, v int64) error {
	var err error = nil

	if r < 0 || c < 0 || m.mtx.count <= c || m.mtx.size <= r {
		err = NewMatrixIndexError(r, c)

	} else {
		m.mtx.writeElement(c, r, v)
	}

	return err
}

// Get a given row form the matrix.
func (m *ColumnAlignedMatrix) Row(r int) (MatrixSegment, error) {
	var row MatrixSegment = nil
	var err error = nil

	if r < 0 || m.mtx.count <= r {
		err = NewMatrixSegmentIndexError(r)

	} else {
		sec := m.mtx.crossSection(r)

		row = createSegmentWithElements(sec)
	}

	return row, err
}

// Get a given column from the matrix.
func (m *ColumnAlignedMatrix) Column(c int) (MatrixSegment, error) {
	var col MatrixSegment = nil
	var err error = nil

	if c < 0 || m.mtx.count <= c {
		err = NewMatrixSegmentIndexError(c)

	} else {
		sec := m.mtx.section(c)

		col = createSegmentWithElements(sec)
	}

	return col, err
}

// Create a new row aligned matrix
//
// In memory, the matrix is kept as a sequence of columns.
func NewColumnAlignedMatrix(rows int, columns int) Matrix {
	m := new(ColumnAlignedMatrix)

	m.mtx = createLinearMatrix(columns, rows)

	return m
}
