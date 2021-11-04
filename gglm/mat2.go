package gglm

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas32"
)

var _ Mat = &Mat2{}
var _ fmt.Stringer = &Mat2{}

type Mat2 struct {
	Arr [4]float32
	*blas32.General
}

func (m *Mat2) At(row, col int) float32 {
	return m.Arr[row*2+col]
}

func (m *Mat2) Set(row, col int, val float32) {
	m.Arr[row*2+col] = val
}

func (m *Mat2) Size() MatSize {
	return MatSize2x2
}

func (m *Mat2) String() string {
	//+ always shows +/- sign; - means pad to the right; 9 means total of 9 digits (or padding if less); .3 means 3 decimals
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f |\n", m.Arr[0], m.Arr[1], m.Arr[2], m.Arr[3])
}

//Add m += m2
func (m *Mat2) Add(m2 *Mat2) {
	m.Arr[0] += m2.Arr[0]
	m.Arr[1] += m2.Arr[1]
	m.Arr[2] += m2.Arr[2]
	m.Arr[3] += m2.Arr[3]
}

//Add m -= m2
func (m *Mat2) Sub(m2 *Mat2) {
	m.Arr[0] -= m2.Arr[0]
	m.Arr[1] -= m2.Arr[1]
	m.Arr[2] -= m2.Arr[2]
	m.Arr[3] -= m2.Arr[3]
}

//Scale m *= x (element wise multiplication)
func (m *Mat2) Scale(x float32) {
	m.Arr[0] *= x
	m.Arr[1] *= x
	m.Arr[2] *= x
	m.Arr[3] *= x
}

//AddMat2 m3 = m1 + m2
func AddMat2(m1, m2 *Mat2) *Mat2 {
	return NewMat2([]float32{
		m1.Arr[0] + m2.Arr[0],
		m1.Arr[1] + m2.Arr[1],
		m1.Arr[2] + m2.Arr[2],
		m1.Arr[3] + m2.Arr[3],
	})
}

//SubMat2 m3 = m1 - m2
func SubMat2(m1, m2 *Mat2) *Mat2 {
	return NewMat2([]float32{
		m1.Arr[0] - m2.Arr[0],
		m1.Arr[1] - m2.Arr[1],
		m1.Arr[2] - m2.Arr[2],
		m1.Arr[3] - m2.Arr[3],
	})
}

//NewMat2 returns the identity matrix if data is nil, otherwise data is copied.
//
//Note: data must be row major
func NewMat2(data []float32) *Mat2 {

	arr := [4]float32{
		1, 0,
		0, 1,
	}
	if data != nil {
		if len(data) != 4 {
			panic("Data must be nil or size 4")
		}

		arr[0] = data[0]
		arr[1] = data[1]
		arr[2] = data[2]
		arr[3] = data[3]
	}

	return &Mat2{
		Arr: arr,
		General: &blas32.General{
			Rows:   2,
			Cols:   2,
			Stride: 2,
			Data:   arr[:],
		},
	}
}
