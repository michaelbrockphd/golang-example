package matrix

// Adapter to create a row aligned matrix
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
func (m *ColumnAlignedMatrix) Element(r int, c int) int64 {
	return m.mtx.readElement(c, r)
}

// Set the value of the specified element.
func (m *ColumnAlignedMatrix) SetElement(r int, c int, v int64) {
	m.mtx.writeElement(c, r, v)
}

// Get a given row form the matrix.
func (m *ColumnAlignedMatrix) Row(r int) *MatrixRow {
	sec := m.mtx.crossSection(r)

	row := NewMatrixRowWithElements(sec)

	return row
}

// Get a given column from the matrix.
func (m *ColumnAlignedMatrix) Column(c int) *MatrixColumn {
	sec := m.mtx.section(c)

	col := NewMatrixColumnWithElements(sec)

	return col
}

// Create a new row aligned matrix
//
// In memory, the matrix is kept as a sequence of rows.
func NewColumnAlignedMatrix(rows int, columns int) Matrix {
	m := new(ColumnAlignedMatrix)

	m.mtx = createLinearMatrix(columns, rows)

	return m
}
