package matrixlib_test

import "testing"
import "github.com/jkraj/jkraj-golang/matrixlib"

func TestNewMatrix(t *testing.T) {

	_, err := matrixlib.NewMatrix(-1)
	if err == nil {
		t.Log("Invalid size should throw an error")
		t.Fail()
	}

	m1, err := matrixlib.NewMatrix(2)
	if err != nil {
		t.Log("Valid size should not cause an error")
		t.Fail()
	}

	if m1.GetSize() != 2 {
		t.Log("Size not properly set")
		t.Fail()
	}
	_, err = matrixlib.NewMatrix(0)
	if err == nil {
		t.Log("Invalid size should throw an error")
		t.Fail()
	}
}

func TestGet(t *testing.T) {

	m1, _ := matrixlib.NewMatrix(2)
	if value, err := m1.Get(2,2); err != nil || value != 0  {
		t.Log("can't access the values")
		t.Fail()
	}
}
