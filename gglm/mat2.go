package gglm

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas32"
)

var _ fmt.Stringer = &Mat2{}

type Mat2 struct {
	*blas32.GeneralCols
}

func (m *Mat2) String() string {
	return fmt.Sprintf("| %f %f |\n| %f %f |", m.Data[0], m.Data[2], m.Data[1], m.Data[3])
}

func NewMat2(data []float32) *Mat2 {

	if data == nil {
		data = make([]float32, 2*2)
	} else if len(data) != 4 {
		panic("Data must be nil or size 4")
	}

	gc := &blas32.GeneralCols{
		Rows:   2,
		Cols:   2,
		Stride: 2,
		Data:   data,
	}

	gc.From(blas32.General{
		Rows:   2,
		Cols:   2,
		Stride: 2,
		Data:   data,
	})

	return &Mat2{
		GeneralCols: gc,
	}
}

//NewMat2ID returns an identity Mat2
func NewMat2ID() *Mat2 {
	return NewMat2([]float32{
		1, 0,
		0, 1,
	})
}
