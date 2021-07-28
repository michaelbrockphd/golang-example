package matrix

// Matrix implementation where all elements are on a continuous array.
//
// In effect, it keeps track of a count sections, each of of a given size.
type linearMatrix struct {
	count int
	size  int

	elements []int64
}

// Read a given element value.
func (m *linearMatrix) readElement(c int, offset int) int64 {
	i := (c * m.size) + offset

	return m.elements[i]
}

// Write a given element value.
func (m *linearMatrix) writeElement(c int, offset int, val int64) {
	i := (c * m.size) + offset

	m.elements[i] = val
}

// Get a given section
func (m *linearMatrix) section(n int) []int64 {
	sec := m.elements[n*m.size : m.size]

	return sec
}

// Get a given cross section
func (m *linearMatrix) crossSection(n int) []int64 {
	sec := make([]int64, m.count)

	i := 0

	offset := n

	for i < m.count {
		sec[i] = m.elements[offset]

		i += 1

		offset += m.size
	}

	return sec
}

// Create a new matrix using a linear array for storage.
func createLinearMatrix(count int, size int) *linearMatrix {
	rtn := new(linearMatrix)

	rtn.count = count

	rtn.size = size

	rtn.elements = make([]int64, count*size)

	return rtn
}
