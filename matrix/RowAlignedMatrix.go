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
func (m *RowAlignedMatrix) Element(r int, c int) int64 {
	return m.mtx.readElement(r, c)
}

// Set the value of the specified element.
func (m *RowAlignedMatrix) SetElement(r int, c int, v int64) {
	m.mtx.writeElement(r, c, v)
}

// Get a given row form the matrix.
func (m *RowAlignedMatrix) Row(r int) MatrixSegment {
	sec := m.mtx.section(r)

	row := createSegmentWithElements(sec)

	return row
}

// Get a given column from the matrix.
func (m *RowAlignedMatrix) Column(c int) MatrixSegment {
	sec := m.mtx.crossSection(c)

	col := createSegmentWithElements(sec)

	return col
}

// Create a new row aligned matrix
//
// In memory, the matrix is kept as a sequence of rows.
func NewRowAlignedMatrix(rows int, columns int) Matrix {
	m := new(RowAlignedMatrix)

	m.mtx = createLinearMatrix(rows, columns)

	return m
}
