package gglm

import (
	"fmt"
)

var _ Mat = &Mat3{}
var _ fmt.Stringer = &Mat3{}

type Mat3 struct {
	Data [9]float32
}

func (m *Mat3) Get(row, col int) float32 {
	return m.Data[row*3+col]
}

func (m *Mat3) Set(row, col int, val float32) {
	m.Data[row*3+col] = val
}

func (m *Mat3) Size() MatSize {
	return MatSize3x3
}

func (m *Mat3) String() string {
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f |\n",
		m.Data[0], m.Data[1], m.Data[2],
		m.Data[3], m.Data[4], m.Data[5],
		m.Data[6], m.Data[7], m.Data[8],
	)
}

//Add m += m2
func (m *Mat3) Add(m2 *Mat3) {

	m.Data[0] += m2.Data[0]
	m.Data[1] += m2.Data[1]
	m.Data[2] += m2.Data[2]

	m.Data[3] += m2.Data[3]
	m.Data[4] += m2.Data[4]
	m.Data[5] += m2.Data[5]

	m.Data[6] += m2.Data[6]
	m.Data[7] += m2.Data[7]
	m.Data[8] += m2.Data[8]
}

//Add m -= m2
func (m *Mat3) Sub(m2 *Mat3) {

	m.Data[0] -= m2.Data[0]
	m.Data[1] -= m2.Data[1]
	m.Data[2] -= m2.Data[2]

	m.Data[3] -= m2.Data[3]
	m.Data[4] -= m2.Data[4]
	m.Data[5] -= m2.Data[5]

	m.Data[6] -= m2.Data[6]
	m.Data[7] -= m2.Data[7]
	m.Data[8] -= m2.Data[8]
}

//Mul m *= m2
func (m *Mat3) Mul(m2 *Mat3) {

	//Indices:
	// 0, 1, 2,
	// 3, 4, 5,
	// 6, 7, 8,

	m.Data = [9]float32{
		m.Data[0]*m2.Data[0] + m.Data[1]*m2.Data[3] + m.Data[2]*m2.Data[6],
		m.Data[0]*m2.Data[1] + m.Data[1]*m2.Data[4] + m.Data[2]*m2.Data[7],
		m.Data[0]*m2.Data[2] + m.Data[1]*m2.Data[5] + m.Data[2]*m2.Data[8],

		m.Data[3]*m2.Data[0] + m.Data[4]*m2.Data[3] + m.Data[5]*m2.Data[6],
		m.Data[3]*m2.Data[1] + m.Data[4]*m2.Data[4] + m.Data[5]*m2.Data[7],
		m.Data[3]*m2.Data[2] + m.Data[4]*m2.Data[5] + m.Data[5]*m2.Data[8],

		m.Data[6]*m2.Data[0] + m.Data[7]*m2.Data[3] + m.Data[8]*m2.Data[6],
		m.Data[6]*m2.Data[1] + m.Data[7]*m2.Data[4] + m.Data[8]*m2.Data[7],
		m.Data[6]*m2.Data[2] + m.Data[7]*m2.Data[5] + m.Data[8]*m2.Data[8],
	}
}

//Scale m *= x (element wise multiplication)
func (m *Mat3) Scale(x float32) {

	m.Data[0] *= x
	m.Data[1] *= x
	m.Data[2] *= x

	m.Data[3] *= x
	m.Data[4] *= x
	m.Data[5] *= x

	m.Data[6] *= x
	m.Data[7] *= x
	m.Data[8] *= x
}

func (m *Mat3) Eq(m2 *Mat3) bool {
	return m.Data == m2.Data
}

//AddMat3 m3 = m1 + m2
func AddMat3(m1, m2 *Mat3) *Mat3 {
	return &Mat3{
		Data: [9]float32{
			m1.Data[0] + m2.Data[0],
			m1.Data[1] + m2.Data[1],
			m1.Data[2] + m2.Data[2],

			m1.Data[3] + m2.Data[3],
			m1.Data[4] + m2.Data[4],
			m1.Data[5] + m2.Data[5],

			m1.Data[6] + m2.Data[6],
			m1.Data[7] + m2.Data[7],
			m1.Data[8] + m2.Data[8],
		},
	}
}

//SubMat3 m3 = m1 - m2
func SubMat3(m1, m2 *Mat3) *Mat3 {
	return &Mat3{
		Data: [9]float32{
			m1.Data[0] - m2.Data[0],
			m1.Data[1] - m2.Data[1],
			m1.Data[2] - m2.Data[2],

			m1.Data[3] - m2.Data[3],
			m1.Data[4] - m2.Data[4],
			m1.Data[5] - m2.Data[5],

			m1.Data[6] - m2.Data[6],
			m1.Data[7] - m2.Data[7],
			m1.Data[8] - m2.Data[8],
		},
	}
}

//MulMat3 m3 = m1 * m2
func MulMat3(m1, m2 *Mat3) *Mat3 {
	return &Mat3{
		Data: [9]float32{
			m1.Data[0]*m2.Data[0] + m1.Data[1]*m2.Data[3] + m1.Data[2]*m2.Data[6],
			m1.Data[0]*m2.Data[1] + m1.Data[1]*m2.Data[4] + m1.Data[2]*m2.Data[7],
			m1.Data[0]*m2.Data[2] + m1.Data[1]*m2.Data[5] + m1.Data[2]*m2.Data[8],

			m1.Data[3]*m2.Data[0] + m1.Data[4]*m2.Data[3] + m1.Data[5]*m2.Data[6],
			m1.Data[3]*m2.Data[1] + m1.Data[4]*m2.Data[4] + m1.Data[5]*m2.Data[7],
			m1.Data[3]*m2.Data[2] + m1.Data[4]*m2.Data[5] + m1.Data[5]*m2.Data[8],

			m1.Data[6]*m2.Data[0] + m1.Data[7]*m2.Data[3] + m1.Data[8]*m2.Data[6],
			m1.Data[6]*m2.Data[1] + m1.Data[7]*m2.Data[4] + m1.Data[8]*m2.Data[7],
			m1.Data[6]*m2.Data[2] + m1.Data[7]*m2.Data[5] + m1.Data[8]*m2.Data[8],
		},
	}
}

//MulMat3Vec3 v2 = m1 * v1
func MulMat3Vec3(m1 *Mat3, v1 *Vec3) *Vec3 {
	return &Vec3{
		Data: [3]float32{
			m1.Data[0]*v1.Data[0] + m1.Data[1]*v1.Data[1] + m1.Data[2]*v1.Data[2],
			m1.Data[3]*v1.Data[0] + m1.Data[4]*v1.Data[1] + m1.Data[5]*v1.Data[2],
			m1.Data[6]*v1.Data[0] + m1.Data[7]*v1.Data[1] + m1.Data[8]*v1.Data[2],
		},
	}
}

//NewMat3Id returns the 3x3 identity matrix
func NewMat3Id() *Mat3 {
	return &Mat3{
		Data: [9]float32{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		},
	}
}
