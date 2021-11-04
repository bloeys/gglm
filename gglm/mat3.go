package gglm

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas32"
)

var _ Mat = &Mat3{}
var _ fmt.Stringer = &Mat3{}

type Mat3 struct {
	Arr [9]float32
	*blas32.General
}

func (m *Mat3) At(row, col int) float32 {
	return m.Arr[row*3+col]
}

func (m *Mat3) Set(row, col int, val float32) {
	m.Arr[row*3+col] = val
}

func (m *Mat3) Size() MatSize {
	return MatSize3x3
}

func (m *Mat3) String() string {
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f |\n",
		m.Arr[0], m.Arr[1], m.Arr[2],
		m.Arr[3], m.Arr[4], m.Arr[5],
		m.Arr[6], m.Arr[7], m.Arr[8],
	)
}

//Add m += m2
func (m *Mat3) Add(m2 *Mat3) {

	m.Arr[0] += m2.Arr[0]
	m.Arr[1] += m2.Arr[1]
	m.Arr[2] += m2.Arr[2]

	m.Arr[3] += m2.Arr[3]
	m.Arr[4] += m2.Arr[4]
	m.Arr[5] += m2.Arr[5]

	m.Arr[6] += m2.Arr[6]
	m.Arr[7] += m2.Arr[7]
	m.Arr[8] += m2.Arr[8]
}

//Add m -= m2
func (m *Mat3) Sub(m2 *Mat3) {

	m.Arr[0] -= m2.Arr[0]
	m.Arr[1] -= m2.Arr[1]
	m.Arr[2] -= m2.Arr[2]

	m.Arr[3] -= m2.Arr[3]
	m.Arr[4] -= m2.Arr[4]
	m.Arr[5] -= m2.Arr[5]

	m.Arr[6] -= m2.Arr[6]
	m.Arr[7] -= m2.Arr[7]
	m.Arr[8] -= m2.Arr[8]
}

//Scale m *= x (element wise multiplication)
func (m *Mat3) Scale(x float32) {

	m.Arr[0] *= x
	m.Arr[1] *= x
	m.Arr[2] *= x

	m.Arr[3] *= x
	m.Arr[4] *= x
	m.Arr[5] *= x

	m.Arr[6] *= x
	m.Arr[7] *= x
	m.Arr[8] *= x
}

//AddMat3 m3 = m1 + m2
func AddMat3(m1, m2 *Mat3) *Mat3 {
	return NewMat3([]float32{
		m1.Arr[0] + m2.Arr[0],
		m1.Arr[1] + m2.Arr[1],
		m1.Arr[2] + m2.Arr[2],

		m1.Arr[3] + m2.Arr[3],
		m1.Arr[4] + m2.Arr[4],
		m1.Arr[5] + m2.Arr[5],

		m1.Arr[6] + m2.Arr[6],
		m1.Arr[7] + m2.Arr[7],
		m1.Arr[8] + m2.Arr[8],
	})
}

//SubMat3 m3 = m1 - m2
func SubMat3(m1, m2 *Mat3) *Mat3 {
	return NewMat3([]float32{
		m1.Arr[0] - m2.Arr[0],
		m1.Arr[1] - m2.Arr[1],
		m1.Arr[2] - m2.Arr[2],

		m1.Arr[3] - m2.Arr[3],
		m1.Arr[4] - m2.Arr[4],
		m1.Arr[5] - m2.Arr[5],

		m1.Arr[6] - m2.Arr[6],
		m1.Arr[7] - m2.Arr[7],
		m1.Arr[8] - m2.Arr[8],
	})
}

//NewMat3 returns the identity matrix if data is nil, otherwise data is copied.
//
//Note: data must be row major
func NewMat3(data []float32) *Mat3 {

	arr := [9]float32{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	if data != nil {
		if len(data) != 9 {
			panic("Data must be nil or size 9")
		}

		arr[0] = data[0]
		arr[1] = data[1]
		arr[2] = data[2]

		arr[3] = data[3]
		arr[4] = data[4]
		arr[5] = data[5]

		arr[6] = data[6]
		arr[7] = data[7]
		arr[8] = data[8]
	}

	return &Mat3{
		Arr: arr,
		General: &blas32.General{
			Rows:   3,
			Cols:   3,
			Stride: 3,
			Data:   arr[:],
		},
	}
}
