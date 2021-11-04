package gglm

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas32"
)

var _ Mat = &Mat4{}
var _ fmt.Stringer = &Mat4{}

type Mat4 struct {
	Arr [16]float32
	*blas32.General
}

func (m *Mat4) At(row, col int) float32 {
	return m.Arr[row*4+col]
}

func (m *Mat4) Set(row, col int, val float32) {
	m.Arr[row*4+col] = val
}

func (m *Mat4) Size() MatSize {
	return MatSize4x4
}

func (m *Mat4) String() string {
	return fmt.Sprintf("\n| %f %f %f %f |\n| %f %f %f %f |\n| %f %f %f %f |\n| %f %f %f %f |\n",
		m.Arr[0], m.Arr[1], m.Arr[2], m.Arr[3],
		m.Arr[4], m.Arr[5], m.Arr[6], m.Arr[7],
		m.Arr[8], m.Arr[9], m.Arr[10], m.Arr[11],
		m.Arr[12], m.Arr[13], m.Arr[14], m.Arr[15],
	)
}

//Add m += m2
func (m *Mat4) Add(m2 *Mat4) {

	m.Arr[0] += m2.Arr[0]
	m.Arr[1] += m2.Arr[1]
	m.Arr[2] += m2.Arr[2]
	m.Arr[3] += m2.Arr[3]

	m.Arr[4] += m2.Arr[4]
	m.Arr[5] += m2.Arr[5]
	m.Arr[6] += m2.Arr[6]
	m.Arr[7] += m2.Arr[7]

	m.Arr[8] += m2.Arr[8]
	m.Arr[9] += m2.Arr[9]
	m.Arr[10] += m2.Arr[10]
	m.Arr[11] += m2.Arr[11]

	m.Arr[12] += m2.Arr[12]
	m.Arr[13] += m2.Arr[13]
	m.Arr[14] += m2.Arr[14]
	m.Arr[15] += m2.Arr[15]
}

//Add m -= m2
func (m *Mat4) Sub(m2 *Mat4) {

	m.Arr[0] -= m2.Arr[0]
	m.Arr[1] -= m2.Arr[1]
	m.Arr[2] -= m2.Arr[2]
	m.Arr[3] -= m2.Arr[3]

	m.Arr[4] -= m2.Arr[4]
	m.Arr[5] -= m2.Arr[5]
	m.Arr[6] -= m2.Arr[6]
	m.Arr[7] -= m2.Arr[7]

	m.Arr[8] -= m2.Arr[8]
	m.Arr[9] -= m2.Arr[9]
	m.Arr[10] -= m2.Arr[10]
	m.Arr[11] -= m2.Arr[11]

	m.Arr[12] -= m2.Arr[12]
	m.Arr[13] -= m2.Arr[13]
	m.Arr[14] -= m2.Arr[14]
	m.Arr[15] -= m2.Arr[15]
}

//Scale m *= x (element wise multiplication)
func (m *Mat4) Scale(x float32) {

	m.Arr[0] *= x
	m.Arr[1] *= x
	m.Arr[2] *= x
	m.Arr[3] *= x

	m.Arr[4] *= x
	m.Arr[5] *= x
	m.Arr[6] *= x
	m.Arr[7] *= x

	m.Arr[8] *= x
	m.Arr[9] *= x
	m.Arr[10] *= x
	m.Arr[11] *= x

	m.Arr[12] *= x
	m.Arr[13] *= x
	m.Arr[14] *= x
	m.Arr[15] *= x
}

//AddMat4 m3 = m1 + m2
func AddMat4(m1, m2 *Mat4) *Mat4 {
	return NewMat4([]float32{
		m1.Arr[0] + m2.Arr[0],
		m1.Arr[1] + m2.Arr[1],
		m1.Arr[2] + m2.Arr[2],
		m1.Arr[3] + m2.Arr[3],

		m1.Arr[4] + m2.Arr[4],
		m1.Arr[5] + m2.Arr[5],
		m1.Arr[6] + m2.Arr[6],
		m1.Arr[7] + m2.Arr[7],

		m1.Arr[8] + m2.Arr[8],
		m1.Arr[9] + m2.Arr[9],
		m1.Arr[10] + m2.Arr[10],
		m1.Arr[11] + m2.Arr[11],

		m1.Arr[12] + m2.Arr[12],
		m1.Arr[13] + m2.Arr[13],
		m1.Arr[14] + m2.Arr[14],
		m1.Arr[15] + m2.Arr[15],
	})
}

//SubMat4 m3 = m1 - m2
func SubMat4(m1, m2 *Mat4) *Mat4 {
	return NewMat4([]float32{
		m1.Arr[0] - m2.Arr[0],
		m1.Arr[1] - m2.Arr[1],
		m1.Arr[2] - m2.Arr[2],
		m1.Arr[3] - m2.Arr[3],

		m1.Arr[4] - m2.Arr[4],
		m1.Arr[5] - m2.Arr[5],
		m1.Arr[6] - m2.Arr[6],
		m1.Arr[7] - m2.Arr[7],

		m1.Arr[8] - m2.Arr[8],
		m1.Arr[9] - m2.Arr[9],
		m1.Arr[10] - m2.Arr[10],
		m1.Arr[11] - m2.Arr[11],

		m1.Arr[12] - m2.Arr[12],
		m1.Arr[13] - m2.Arr[13],
		m1.Arr[14] - m2.Arr[14],
		m1.Arr[15] - m2.Arr[15],
	})
}

//NewMat4 returns the identity matrix if data is nil, otherwise data is copied.
//
//Note: data must be row major
func NewMat4(data []float32) *Mat4 {

	arr := [16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	if data != nil {
		if len(data) != 16 {
			panic("Data must be nil or size 16")
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
		arr[9] = data[9]
		arr[10] = data[10]
		arr[11] = data[11]

		arr[12] = data[12]
		arr[13] = data[13]
		arr[14] = data[14]
		arr[15] = data[15]
	}

	return &Mat4{
		Arr: arr,
		General: &blas32.General{
			Rows:   4,
			Cols:   4,
			Stride: 4,
			Data:   arr[:],
		},
	}
}
