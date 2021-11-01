package gglm

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas32"
)

var _ Mat = &Mat2{}
var _ fmt.Stringer = &Mat2{}

type Mat2 struct {
	*blas32.General
}

func (m *Mat2) At(row, col int) float32 {
	return m.Data[row*2+col]
}

func (m *Mat2) Set(row, col int, val float32) {
	m.Data[row*2+col] = val
}

func (m *Mat2) Size() int {
	return 2
}

func (m *Mat2) String() string {
	return fmt.Sprintf("| %f %f |\n| %f %f |", m.Data[0], m.Data[1], m.Data[2], m.Data[3])
}

//Scale m *= x (element wise multiplication)
func (m *Mat2) Scale(x float32) {
	m.Data[0] *= x
	m.Data[1] *= x
	m.Data[2] *= x
	m.Data[3] *= x
}

//NewMat2 returns the identity matrix if data=nil, otherwise data is used
//as the backing array.
//
//Note: data must be row major
func NewMat2(data []float32) *Mat2 {

	if data == nil {
		data = []float32{
			1, 0,
			0, 1,
		}
	} else if len(data) != 4 {
		panic("Data must be nil or size 4")
	}

	return &Mat2{
		General: &blas32.General{
			Rows:   2,
			Cols:   2,
			Stride: 2,
			Data:   data,
		},
	}
}
