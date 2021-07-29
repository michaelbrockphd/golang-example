package matrix

// Return the product of the two provided matrix segments.
func MultiplySegments(a MatrixSegment, b MatrixSegment) int64 {
	rtn := int64(0)

	len := a.Length()

	i := 0

	for i < len {
		rtn += (a.Element(i) * b.Element(i))

		i += 1
	}

	return rtn
}

// Return the product of the two provided matrixes.
func Multiply(a Matrix, b Matrix) Matrix {
	numR := a.Rows()
	numC := b.Columns()

	rtn := NewRowAlignedMatrix(numR, numC)

	idxR := 0
	idxC := 0

	var currR MatrixSegment = nil
	var currC MatrixSegment = nil

	for idxR < numR {
		currR = a.Row(idxR)

		idxC = 0

		for idxC < numC {
			currC = b.Column(idxC)

			rtn.SetElement(idxR, idxC, MultiplySegments(currR, currC))

			idxC += 1
		}

		idxR += 1
	}

	return rtn
}
