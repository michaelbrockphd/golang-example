package main

import (
	"fmt"
	"os"

	"learningOwl.org/matrix"
)

func printMatrix(f *os.File, m matrix.Matrix) {
	r := 0

	for r < m.Rows() {
		fmt.Fprintf(f, "%v\n", m.Row(r).String())

		r += 1
	}
}

func main() {
	stdout := os.Stdout

	size := 50

	a := matrix.NewRowAlignedMatrix(size, size)

	matrix.InitializeMatrix(a)

	printMatrix(stdout, a)

	fmt.Fprintln(stdout)

	//b := matrix.NewComputedIdentityMatrix(size, 2)
	//b := matrix.NewRowAlignedMatrix(size, size)
	b := matrix.NewColumnAlignedMatrix(size, size)

	matrix.InitializeIdentityMatrix(b, 2)

	printMatrix(stdout, b)

	fmt.Fprintln(stdout)

	c := matrix.Multiply(a, b)

	printMatrix(stdout, c)

	fmt.Fprintln(stdout)
}
