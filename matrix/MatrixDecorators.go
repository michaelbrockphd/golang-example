package matrix

/*
// Get the value of the specified element.
func (m *linearMatrix) Element(r int, c int) int64 {
	i := (r * m.size) + c

	return m.elements[i]
}

// Set the value of the specified element.
func (m *linearMatrix) SetElement(r int, c int, v int64) {
	i := (r * m.size) + c

	m.elements[i] = v
}

// Get a given row form the matrix.
func (m *linearMatrix) Row(r int) *MatrixRow {
	tmp := m.elements[r*m.size : m.size]

	row := NewMatrixRowWithElements(tmp)

	return row
}

// Get a given column from the matrix.
func (m *linearMatrix) Column(c int) *MatrixColumn {
	col := make([]int64, m.count)

	r := 0

	i := 0

	for r < m.count {
		i = (r * m.size) + c

		col[r] = m.elements[i]

		i += 1
	}

	return NewMatrixColumnWithElements(col)
}
*/
