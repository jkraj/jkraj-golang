package matrixlib

import "errors"

type matrix struct {
	Size int "Size of matrix"
	values [][]int "Values of matrix"
}

/* we can use simillarly 
type other matrix 
m1 matrix
other m2 = other(m1)
*/

func (m *matrix)GetSize() int { // always will have 1 receiver
	return m.Size
}

func NewMatrix(size int) (result *matrix, err error) {
	if size < 1 {
		return result, errors.New("Invalid size")
	}
	result = &matrix{}
	result.Size = size
	result.values = make([][]int, size)
	for i, _ := range result.values {
		result.values[i] = make([]int, size)
	}

	// Assign the random value to the matrix
	return
}

func iterate(lhs matrix, rhs matrix, operation matrixop) (result *matrix, err error) {
	if lhs.Size != rhs.Size {
		err = errors.New("Both matrix have different size")
		return
	}

	if lhs.values == nil || rhs.values == nil {
		err = errors.New("matrix values are empty")
		return
	}
	result, _ = NewMatrix(lhs.Size)

	for i, values := range lhs.values {
		for j, innerval := range values {
			rhsValue, _ := rhs.Get(i, j)
			result.Set(i, j, operation(innerval, rhsValue))
		}
	}

	return
}

// First class functions
// variable of a function type.
// any value of a funtion type is pointers

type matrixop func(int, int) int

func addops(lval int, rval int) int {
	return lval + rval
}

func subops(lval int, rval int) int {
	return lval - rval
}

func Sub(lhs matrix, rhs matrix) (*matrix, error) {
	return iterate(lhs, rhs, subops)
}


func Add1(lhs matrix, rhs matrix) (*matrix, error) {
        return iterate(lhs, rhs, func(m int, n int) { // Lambda/Clousures 
					return m + n
					})
}

func Add2(lhs matrix, rhs matrix) (*matrix, error) {
	a := 22
        return iterate(lhs, rhs, func(m int, n int) { // variable scope
					return m + n + a
					})
}

// * is needed during declaration but while calling this method & will be used.
// error is pointer so we can return nil
func (m *matrix) Set(row int, col int, value int) error {
	return nil
}

func (m *matrix) Get(row int, col int) (int, error) {
	res := 0
	if row <= 0 && col <=0 {
		return res, errors.New("Row & Size should be same")
	}
	return m.values[row][col], nil
}
