package matrix

// Matrix implementation where all elements are on a continuous array.
type linearMatrix struct {
	rows    int
	columns int

	elements []int64
}

// Get the number of rows in the matrix.
func (m *linearMatrix) Rows() int {
	return m.rows
}

// Get the number of columns in the matrix.
func (m *linearMatrix) Columns() int {
	return m.columns
}

// Get the value of the specified element.
func (m *linearMatrix) Element(r int, c int) int64 {
	i := (r * m.columns) + c

	return m.elements[i]
}

// Set the value of the specified element.
func (m *linearMatrix) SetElement(r int, c int, v int64) {
	i := (r * m.columns) + c

	m.elements[i] = v
}

// Get a given row form the matrix.
func (m *linearMatrix) Row(r int) *MatrixRow {
	tmp := m.elements[r*m.columns : m.columns]

	row := NewMatrixRowWithElements(tmp)

	return row
}

// Get a given column from the matrix.
func (m *linearMatrix) Column(c int) *MatrixColumn {
	col := make([]int64, m.rows)

	r := 0

	i := 0

	for r < m.rows {
		i = (r * m.columns) + c

		col[r] = m.elements[i]

		i += 1
	}

	return NewMatrixColumnWithElements(col)
}

// Create a new matrix using a linear array for storage.
func NewLinearMatrix(rows int, cols int) Matrix {
	rtn := new(linearMatrix)

	rtn.rows = rows

	rtn.columns = cols

	rtn.elements = make([]int64, rows*cols)

	return rtn
}
