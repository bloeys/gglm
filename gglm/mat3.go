package gglm

import (
	"fmt"
)

var _ Mat = &Mat3{}
var _ fmt.Stringer = &Mat3{}

type Mat3 struct {
	Data [3][3]float32
}

func (m *Mat3) Get(row, col int) float32 {
	return m.Data[col][row]
}

func (m *Mat3) Set(row, col int, val float32) {
	m.Data[col][row] = val
}

func (m *Mat3) Size() MatSize {
	return MatSize3x3
}

func (m *Mat3) String() string {
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f |\n",
		m.Data[0][0], m.Data[0][1], m.Data[0][2],
		m.Data[1][0], m.Data[1][1], m.Data[1][2],
		m.Data[2][0], m.Data[2][1], m.Data[2][2],
	)
}

//Add m += m2
func (m *Mat3) Add(m2 *Mat3) *Mat3 {

	m.Data[0][0] += m2.Data[0][0]
	m.Data[0][1] += m2.Data[0][1]
	m.Data[0][2] += m2.Data[0][2]

	m.Data[1][0] += m2.Data[1][0]
	m.Data[1][1] += m2.Data[1][1]
	m.Data[1][2] += m2.Data[1][2]

	m.Data[2][0] += m2.Data[2][0]
	m.Data[2][1] += m2.Data[2][1]
	m.Data[2][2] += m2.Data[2][2]
	return m
}

//Add m -= m2
func (m *Mat3) Sub(m2 *Mat3) *Mat3 {

	m.Data[0][0] -= m2.Data[0][0]
	m.Data[0][1] -= m2.Data[0][1]
	m.Data[0][2] -= m2.Data[0][2]

	m.Data[1][0] -= m2.Data[1][0]
	m.Data[1][1] -= m2.Data[1][1]
	m.Data[1][2] -= m2.Data[1][2]

	m.Data[2][0] -= m2.Data[2][0]
	m.Data[2][1] -= m2.Data[2][1]
	m.Data[2][2] -= m2.Data[2][2]
	return m
}

//Mul m *= m2
func (m *Mat3) Mul(m2 *Mat3) *Mat3 {

	//Array indices:
	// 00, 10, 20,
	// 01, 11, 21,
	// 02, 12, 22,

	m.Data = [3][3]float32{
		{
			m.Data[0][0]*m2.Data[0][0] + m.Data[1][0]*m2.Data[0][1] + m.Data[2][0]*m2.Data[0][2],
			m.Data[0][1]*m2.Data[0][0] + m.Data[1][1]*m2.Data[0][1] + m.Data[2][1]*m2.Data[0][2],
			m.Data[0][2]*m2.Data[0][0] + m.Data[1][2]*m2.Data[0][1] + m.Data[2][2]*m2.Data[0][2],
		},
		{
			m.Data[0][0]*m2.Data[1][0] + m.Data[1][0]*m2.Data[1][1] + m.Data[2][0]*m2.Data[1][2],
			m.Data[0][1]*m2.Data[1][0] + m.Data[1][1]*m2.Data[1][1] + m.Data[2][1]*m2.Data[1][2],
			m.Data[0][2]*m2.Data[1][0] + m.Data[1][2]*m2.Data[1][1] + m.Data[2][2]*m2.Data[1][2],
		},
		{
			m.Data[0][0]*m2.Data[2][0] + m.Data[1][0]*m2.Data[2][1] + m.Data[2][0]*m2.Data[2][2],
			m.Data[0][1]*m2.Data[2][0] + m.Data[1][1]*m2.Data[2][1] + m.Data[2][1]*m2.Data[2][2],
			m.Data[0][2]*m2.Data[2][0] + m.Data[1][2]*m2.Data[2][1] + m.Data[2][2]*m2.Data[2][2],
		},
	}
	return m
}

//Scale m *= x (element wise multiplication)
func (m *Mat3) Scale(x float32) *Mat3 {

	m.Data[0][0] *= x
	m.Data[0][1] *= x
	m.Data[0][2] *= x

	m.Data[1][0] *= x
	m.Data[1][1] *= x
	m.Data[1][2] *= x

	m.Data[2][0] *= x
	m.Data[2][1] *= x
	m.Data[2][2] *= x
	return m
}

func (v *Mat3) Clone() *Mat3 {
	return &Mat3{Data: v.Data}
}

func (m *Mat3) Eq(m2 *Mat3) bool {
	return m.Data == m2.Data
}

//AddMat3 m3 = m1 + m2
func AddMat3(m1, m2 *Mat3) *Mat3 {
	return &Mat3{
		Data: [3][3]float32{
			{
				m1.Data[0][0] + m2.Data[0][0],
				m1.Data[0][1] + m2.Data[0][1],
				m1.Data[0][2] + m2.Data[0][2],
			},
			{
				m1.Data[1][0] + m2.Data[1][0],
				m1.Data[1][1] + m2.Data[1][1],
				m1.Data[1][2] + m2.Data[1][2],
			},
			{
				m1.Data[2][0] + m2.Data[2][0],
				m1.Data[2][1] + m2.Data[2][1],
				m1.Data[2][2] + m2.Data[2][2],
			},
		},
	}
}

//SubMat3 m3 = m1 - m2
func SubMat3(m1, m2 *Mat3) *Mat3 {
	return &Mat3{
		Data: [3][3]float32{
			{
				m1.Data[0][0] - m2.Data[0][0],
				m1.Data[0][1] - m2.Data[0][1],
				m1.Data[0][2] - m2.Data[0][2],
			},
			{
				m1.Data[1][0] - m2.Data[1][0],
				m1.Data[1][1] - m2.Data[1][1],
				m1.Data[1][2] - m2.Data[1][2],
			},
			{
				m1.Data[2][0] - m2.Data[2][0],
				m1.Data[2][1] - m2.Data[2][1],
				m1.Data[2][2] - m2.Data[2][2],
			},
		},
	}
}

//MulMat3 m3 = m1 * m2
func MulMat3(m1, m2 *Mat3) *Mat3 {
	return &Mat3{
		Data: [3][3]float32{
			{
				m1.Data[0][0]*m2.Data[0][0] + m1.Data[1][0]*m2.Data[0][1] + m1.Data[2][0]*m2.Data[0][2],
				m1.Data[0][1]*m2.Data[0][0] + m1.Data[1][1]*m2.Data[0][1] + m1.Data[2][1]*m2.Data[0][2],
				m1.Data[0][2]*m2.Data[0][0] + m1.Data[1][2]*m2.Data[0][1] + m1.Data[2][2]*m2.Data[0][2],
			},
			{
				m1.Data[0][0]*m2.Data[1][0] + m1.Data[1][0]*m2.Data[1][1] + m1.Data[2][0]*m2.Data[1][2],
				m1.Data[0][1]*m2.Data[1][0] + m1.Data[1][1]*m2.Data[1][1] + m1.Data[2][1]*m2.Data[1][2],
				m1.Data[0][2]*m2.Data[1][0] + m1.Data[1][2]*m2.Data[1][1] + m1.Data[2][2]*m2.Data[1][2],
			},
			{
				m1.Data[0][0]*m2.Data[2][0] + m1.Data[1][0]*m2.Data[2][1] + m1.Data[2][0]*m2.Data[2][2],
				m1.Data[0][1]*m2.Data[2][0] + m1.Data[1][1]*m2.Data[2][1] + m1.Data[2][1]*m2.Data[2][2],
				m1.Data[0][2]*m2.Data[2][0] + m1.Data[1][2]*m2.Data[2][1] + m1.Data[2][2]*m2.Data[2][2],
			},
		},
	}
}

//MulMat3Vec3 v2 = m1 * v1
func MulMat3Vec3(m1 *Mat3, v1 *Vec3) *Vec3 {
	return &Vec3{
		Data: [3]float32{
			m1.Data[0][0]*v1.Data[0] + m1.Data[1][0]*v1.Data[1] + m1.Data[2][0]*v1.Data[2],
			m1.Data[0][1]*v1.Data[0] + m1.Data[1][1]*v1.Data[1] + m1.Data[2][1]*v1.Data[2],
			m1.Data[0][2]*v1.Data[0] + m1.Data[1][2]*v1.Data[1] + m1.Data[2][2]*v1.Data[2],
		},
	}
}

//NewMat3Id returns the 3x3 identity matrix
func NewMat3Id() *Mat3 {
	return &Mat3{
		Data: [3][3]float32{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}
