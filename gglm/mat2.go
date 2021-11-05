package gglm

import (
	"fmt"
)

var _ Mat = &Mat2{}
var _ fmt.Stringer = &Mat2{}

type Mat2 struct {
	Data [4]float32
}

func (m *Mat2) At(row, col int) float32 {
	return m.Data[row*2+col]
}

func (m *Mat2) Set(row, col int, val float32) {
	m.Data[row*2+col] = val
}

func (m *Mat2) Size() MatSize {
	return MatSize2x2
}

func (m *Mat2) String() string {
	//+ always shows +/- sign; - means pad to the right; 9 means total of 9 digits (or padding if less); .3 means 3 decimals
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f |\n", m.Data[0], m.Data[1], m.Data[2], m.Data[3])
}

//Add m += m2
func (m *Mat2) Add(m2 *Mat2) {
	m.Data[0] += m2.Data[0]
	m.Data[1] += m2.Data[1]
	m.Data[2] += m2.Data[2]
	m.Data[3] += m2.Data[3]
}

//Add m -= m2
func (m *Mat2) Sub(m2 *Mat2) {
	m.Data[0] -= m2.Data[0]
	m.Data[1] -= m2.Data[1]
	m.Data[2] -= m2.Data[2]
	m.Data[3] -= m2.Data[3]
}

//Mul m *= m2
func (m *Mat2) Mul(m2 *Mat2) {
	m.Data = [4]float32{
		m.Data[0]*m2.Data[0] + m.Data[1]*m2.Data[2],
		m.Data[0]*m2.Data[1] + m.Data[1]*m2.Data[3],

		m.Data[2]*m2.Data[0] + m.Data[3]*m2.Data[2],
		m.Data[2]*m2.Data[1] + m.Data[3]*m2.Data[3],
	}
}

//Scale m *= x (element wise multiplication)
func (m *Mat2) Scale(x float32) {
	m.Data[0] *= x
	m.Data[1] *= x
	m.Data[2] *= x
	m.Data[3] *= x
}

func (m *Mat2) Eq(m2 *Mat2) bool {
	return m.Data == m2.Data
}

//AddMat2 m3 = m1 + m2
func AddMat2(m1, m2 *Mat2) *Mat2 {
	return &Mat2{
		Data: [4]float32{
			m1.Data[0] + m2.Data[0],
			m1.Data[1] + m2.Data[1],
			m1.Data[2] + m2.Data[2],
			m1.Data[3] + m2.Data[3],
		},
	}
}

//SubMat2 m3 = m1 - m2
func SubMat2(m1, m2 *Mat2) *Mat2 {
	return &Mat2{
		Data: [4]float32{
			m1.Data[0] - m2.Data[0],
			m1.Data[1] - m2.Data[1],
			m1.Data[2] - m2.Data[2],
			m1.Data[3] - m2.Data[3],
		},
	}
}

//MulMat2 m3 = m1 * m2
func MulMat2(m1, m2 *Mat2) *Mat2 {
	return &Mat2{
		Data: [4]float32{
			m1.Data[0]*m2.Data[0] + m1.Data[1]*m2.Data[2],
			m1.Data[0]*m2.Data[1] + m1.Data[1]*m2.Data[3],

			m1.Data[2]*m2.Data[0] + m1.Data[3]*m2.Data[2],
			m1.Data[2]*m2.Data[1] + m1.Data[3]*m2.Data[3],
		},
	}
}

//NewMat2Id returns the 2x2 identity matrix
func NewMat2Id() *Mat2 {
	return &Mat2{
		Data: [4]float32{
			1, 0,
			0, 1,
		},
	}
}
