package gglm

import (
	"gonum.org/v1/gonum/blas/blas32"
)

type Mat2 struct {
	*blas32.GeneralCols
}

func NewMat2(data []float32) *Mat2 {

	if data == nil {
		data = make([]float32, 2*2)
	} else if len(data) != 4 {
		panic("Data must be nil or size 4")
	}

	gc := &blas32.GeneralCols{}
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
