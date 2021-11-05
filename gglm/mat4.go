package gglm

import (
	"fmt"
)

var _ Mat = &Mat4{}
var _ fmt.Stringer = &Mat4{}

type Mat4 struct {
	Data [16]float32
}

func (m *Mat4) At(row, col int) float32 {
	return m.Data[row*4+col]
}

func (m *Mat4) Set(row, col int, val float32) {
	m.Data[row*4+col] = val
}

func (m *Mat4) Size() MatSize {
	return MatSize4x4
}

func (m *Mat4) String() string {
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n",
		m.Data[0], m.Data[1], m.Data[2], m.Data[3],
		m.Data[4], m.Data[5], m.Data[6], m.Data[7],
		m.Data[8], m.Data[9], m.Data[10], m.Data[11],
		m.Data[12], m.Data[13], m.Data[14], m.Data[15],
	)
}

//Add m += m2
func (m *Mat4) Add(m2 *Mat4) {

	m.Data[0] += m2.Data[0]
	m.Data[1] += m2.Data[1]
	m.Data[2] += m2.Data[2]
	m.Data[3] += m2.Data[3]

	m.Data[4] += m2.Data[4]
	m.Data[5] += m2.Data[5]
	m.Data[6] += m2.Data[6]
	m.Data[7] += m2.Data[7]

	m.Data[8] += m2.Data[8]
	m.Data[9] += m2.Data[9]
	m.Data[10] += m2.Data[10]
	m.Data[11] += m2.Data[11]

	m.Data[12] += m2.Data[12]
	m.Data[13] += m2.Data[13]
	m.Data[14] += m2.Data[14]
	m.Data[15] += m2.Data[15]
}

//Add m -= m2
func (m *Mat4) Sub(m2 *Mat4) {

	m.Data[0] -= m2.Data[0]
	m.Data[1] -= m2.Data[1]
	m.Data[2] -= m2.Data[2]
	m.Data[3] -= m2.Data[3]

	m.Data[4] -= m2.Data[4]
	m.Data[5] -= m2.Data[5]
	m.Data[6] -= m2.Data[6]
	m.Data[7] -= m2.Data[7]

	m.Data[8] -= m2.Data[8]
	m.Data[9] -= m2.Data[9]
	m.Data[10] -= m2.Data[10]
	m.Data[11] -= m2.Data[11]

	m.Data[12] -= m2.Data[12]
	m.Data[13] -= m2.Data[13]
	m.Data[14] -= m2.Data[14]
	m.Data[15] -= m2.Data[15]
}

//Mul m *= m2
func (m *Mat4) Mul(m2 *Mat4) {

	//Indices:
	// 00, 01, 02, 03,
	// 04, 05, 06, 07,
	// 08, 09, 10, 11,
	// 12, 13, 14, 15,

	//Seems to improve performance by ~5% (18ns/op -> 17ns/op).
	//Works by improving cache usage by putting 0,4,8,12 together instead of faraway in the array?
	a := m2.Data[0]
	b := m2.Data[4]
	c := m2.Data[8]
	d := m2.Data[12]

	m.Data = [16]float32{
		m.Data[0]*a + m.Data[1]*b + m.Data[2]*c + m.Data[3]*d,
		m.Data[0]*m2.Data[1] + m.Data[1]*m2.Data[5] + m.Data[2]*m2.Data[9] + m.Data[3]*m2.Data[13],
		m.Data[0]*m2.Data[2] + m.Data[1]*m2.Data[6] + m.Data[2]*m2.Data[10] + m.Data[3]*m2.Data[14],
		m.Data[0]*m2.Data[3] + m.Data[1]*m2.Data[7] + m.Data[2]*m2.Data[11] + m.Data[3]*m2.Data[15],

		m.Data[4]*a + m.Data[5]*b + m.Data[6]*c + m.Data[7]*d,
		m.Data[4]*m2.Data[1] + m.Data[5]*m2.Data[5] + m.Data[6]*m2.Data[9] + m.Data[7]*m2.Data[13],
		m.Data[4]*m2.Data[2] + m.Data[5]*m2.Data[6] + m.Data[6]*m2.Data[10] + m.Data[7]*m2.Data[14],
		m.Data[4]*m2.Data[3] + m.Data[5]*m2.Data[7] + m.Data[6]*m2.Data[11] + m.Data[7]*m2.Data[15],

		m.Data[8]*a + m.Data[9]*b + m.Data[10]*c + m.Data[11]*d,
		m.Data[8]*m2.Data[1] + m.Data[9]*m2.Data[5] + m.Data[10]*m2.Data[9] + m.Data[11]*m2.Data[13],
		m.Data[8]*m2.Data[2] + m.Data[9]*m2.Data[6] + m.Data[10]*m2.Data[10] + m.Data[11]*m2.Data[14],
		m.Data[8]*m2.Data[3] + m.Data[9]*m2.Data[7] + m.Data[10]*m2.Data[11] + m.Data[11]*m2.Data[15],

		m.Data[12]*a + m.Data[13]*b + m.Data[14]*c + m.Data[15]*d,
		m.Data[12]*m2.Data[1] + m.Data[13]*m2.Data[5] + m.Data[14]*m2.Data[9] + m.Data[15]*m2.Data[13],
		m.Data[12]*m2.Data[2] + m.Data[13]*m2.Data[6] + m.Data[14]*m2.Data[10] + m.Data[15]*m2.Data[14],
		m.Data[12]*m2.Data[3] + m.Data[13]*m2.Data[7] + m.Data[14]*m2.Data[11] + m.Data[15]*m2.Data[15],
	}
}

//Scale m *= x (element wise multiplication)
func (m *Mat4) Scale(x float32) {

	m.Data[0] *= x
	m.Data[1] *= x
	m.Data[2] *= x
	m.Data[3] *= x

	m.Data[4] *= x
	m.Data[5] *= x
	m.Data[6] *= x
	m.Data[7] *= x

	m.Data[8] *= x
	m.Data[9] *= x
	m.Data[10] *= x
	m.Data[11] *= x

	m.Data[12] *= x
	m.Data[13] *= x
	m.Data[14] *= x
	m.Data[15] *= x
}

//AddMat4 m3 = m1 + m2
func AddMat4(m1, m2 *Mat4) *Mat4 {
	return &Mat4{
		Data: [16]float32{
			m1.Data[0] + m2.Data[0],
			m1.Data[1] + m2.Data[1],
			m1.Data[2] + m2.Data[2],
			m1.Data[3] + m2.Data[3],

			m1.Data[4] + m2.Data[4],
			m1.Data[5] + m2.Data[5],
			m1.Data[6] + m2.Data[6],
			m1.Data[7] + m2.Data[7],

			m1.Data[8] + m2.Data[8],
			m1.Data[9] + m2.Data[9],
			m1.Data[10] + m2.Data[10],
			m1.Data[11] + m2.Data[11],

			m1.Data[12] + m2.Data[12],
			m1.Data[13] + m2.Data[13],
			m1.Data[14] + m2.Data[14],
			m1.Data[15] + m2.Data[15],
		},
	}
}

//SubMat4 m3 = m1 - m2
func SubMat4(m1, m2 *Mat4) *Mat4 {
	return &Mat4{
		Data: [16]float32{
			m1.Data[0] - m2.Data[0],
			m1.Data[1] - m2.Data[1],
			m1.Data[2] - m2.Data[2],
			m1.Data[3] - m2.Data[3],

			m1.Data[4] - m2.Data[4],
			m1.Data[5] - m2.Data[5],
			m1.Data[6] - m2.Data[6],
			m1.Data[7] - m2.Data[7],

			m1.Data[8] - m2.Data[8],
			m1.Data[9] - m2.Data[9],
			m1.Data[10] - m2.Data[10],
			m1.Data[11] - m2.Data[11],

			m1.Data[12] - m2.Data[12],
			m1.Data[13] - m2.Data[13],
			m1.Data[14] - m2.Data[14],
			m1.Data[15] - m2.Data[15],
		},
	}
}

//MulMat4 m3 = m1 * m2
func MulMat4(m1, m2 *Mat4) *Mat4 {
	return &Mat4{
		Data: [16]float32{
			m1.Data[0]*m2.Data[0] + m1.Data[1]*m2.Data[4] + m1.Data[2]*m2.Data[8] + m1.Data[3]*m2.Data[12],
			m1.Data[0]*m2.Data[1] + m1.Data[1]*m2.Data[5] + m1.Data[2]*m2.Data[9] + m1.Data[3]*m2.Data[13],
			m1.Data[0]*m2.Data[2] + m1.Data[1]*m2.Data[6] + m1.Data[2]*m2.Data[10] + m1.Data[3]*m2.Data[14],
			m1.Data[0]*m2.Data[3] + m1.Data[1]*m2.Data[7] + m1.Data[2]*m2.Data[11] + m1.Data[3]*m2.Data[15],

			m1.Data[4]*m2.Data[0] + m1.Data[5]*m2.Data[4] + m1.Data[6]*m2.Data[8] + m1.Data[7]*m2.Data[12],
			m1.Data[4]*m2.Data[1] + m1.Data[5]*m2.Data[5] + m1.Data[6]*m2.Data[9] + m1.Data[7]*m2.Data[13],
			m1.Data[4]*m2.Data[2] + m1.Data[5]*m2.Data[6] + m1.Data[6]*m2.Data[10] + m1.Data[7]*m2.Data[14],
			m1.Data[4]*m2.Data[3] + m1.Data[5]*m2.Data[7] + m1.Data[6]*m2.Data[11] + m1.Data[7]*m2.Data[15],

			m1.Data[8]*m2.Data[0] + m1.Data[9]*m2.Data[4] + m1.Data[10]*m2.Data[8] + m1.Data[11]*m2.Data[12],
			m1.Data[8]*m2.Data[1] + m1.Data[9]*m2.Data[5] + m1.Data[10]*m2.Data[9] + m1.Data[11]*m2.Data[13],
			m1.Data[8]*m2.Data[2] + m1.Data[9]*m2.Data[6] + m1.Data[10]*m2.Data[10] + m1.Data[11]*m2.Data[14],
			m1.Data[8]*m2.Data[3] + m1.Data[9]*m2.Data[7] + m1.Data[10]*m2.Data[11] + m1.Data[11]*m2.Data[15],

			m1.Data[12]*m2.Data[0] + m1.Data[13]*m2.Data[4] + m1.Data[14]*m2.Data[8] + m1.Data[15]*m2.Data[12],
			m1.Data[12]*m2.Data[1] + m1.Data[13]*m2.Data[5] + m1.Data[14]*m2.Data[9] + m1.Data[15]*m2.Data[13],
			m1.Data[12]*m2.Data[2] + m1.Data[13]*m2.Data[6] + m1.Data[14]*m2.Data[10] + m1.Data[15]*m2.Data[14],
			m1.Data[12]*m2.Data[3] + m1.Data[13]*m2.Data[7] + m1.Data[14]*m2.Data[11] + m1.Data[15]*m2.Data[15],
		},
	}
}

//NewMat4Id returns the 4x4 identity matrix
func NewMat4Id() *Mat4 {
	return &Mat4{
		Data: [16]float32{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		},
	}
}
