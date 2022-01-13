package gglm

import (
	"fmt"
)

var _ Mat = &Mat4{}
var _ fmt.Stringer = &Mat4{}

type Mat4 struct {
	Data [4][4]float32
}

func (m *Mat4) Get(row, col int) float32 {
	return m.Data[col][row]
}

func (m *Mat4) Set(row, col int, val float32) {
	m.Data[col][row] = val
}

func (m *Mat4) Size() MatSize {
	return MatSize4x4
}

func (m *Mat4) String() string {
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f %+-9.3f %+-9.3f |\n",
		m.Data[0][0], m.Data[0][1], m.Data[0][2], m.Data[0][3],
		m.Data[1][0], m.Data[1][1], m.Data[1][2], m.Data[1][3],
		m.Data[2][0], m.Data[2][1], m.Data[2][2], m.Data[2][3],
		m.Data[3][0], m.Data[3][1], m.Data[3][2], m.Data[3][3],
	)
}

func (m *Mat4) Col(c int) *Vec4 {
	return &Vec4{Data: m.Data[c]}
}

//Add m += m2
func (m *Mat4) Add(m2 *Mat4) *Mat4 {

	m.Data[0][0] += m2.Data[0][0]
	m.Data[0][1] += m2.Data[0][1]
	m.Data[0][2] += m2.Data[0][2]
	m.Data[0][3] += m2.Data[0][3]

	m.Data[1][0] += m2.Data[1][0]
	m.Data[1][1] += m2.Data[1][1]
	m.Data[1][2] += m2.Data[1][2]
	m.Data[1][3] += m2.Data[1][3]

	m.Data[2][0] += m2.Data[2][0]
	m.Data[2][1] += m2.Data[2][1]
	m.Data[2][2] += m2.Data[2][2]
	m.Data[2][3] += m2.Data[2][3]

	m.Data[3][0] += m2.Data[3][0]
	m.Data[3][1] += m2.Data[3][1]
	m.Data[3][2] += m2.Data[3][2]
	m.Data[3][3] += m2.Data[3][3]

	return m
}

//Add m -= m2
func (m *Mat4) Sub(m2 *Mat4) *Mat4 {

	m.Data[0][0] -= m2.Data[0][0]
	m.Data[0][1] -= m2.Data[0][1]
	m.Data[0][2] -= m2.Data[0][2]
	m.Data[0][3] -= m2.Data[0][3]

	m.Data[1][0] -= m2.Data[1][0]
	m.Data[1][1] -= m2.Data[1][1]
	m.Data[1][2] -= m2.Data[1][2]
	m.Data[1][3] -= m2.Data[1][3]

	m.Data[2][0] -= m2.Data[2][0]
	m.Data[2][1] -= m2.Data[2][1]
	m.Data[2][2] -= m2.Data[2][2]
	m.Data[2][3] -= m2.Data[2][3]

	m.Data[3][0] -= m2.Data[3][0]
	m.Data[3][1] -= m2.Data[3][1]
	m.Data[3][2] -= m2.Data[3][2]
	m.Data[3][3] -= m2.Data[3][3]
	return m
}

//Mul m *= m2
func (m *Mat4) Mul(m2 *Mat4) *Mat4 {

	//Array indices:
	// 00, 10, 20, 30,
	// 01, 11, 21, 31,
	// 02, 12, 22, 32,
	// 03, 13, 23, 33,

	//Improves performance by ~8%
	m00 := m.Data[0][0]
	m01 := m.Data[0][1]
	m02 := m.Data[0][2]
	m03 := m.Data[0][3]

	m10 := m.Data[1][0]
	m11 := m.Data[1][1]
	m12 := m.Data[1][2]
	m13 := m.Data[1][3]

	m20 := m.Data[2][0]
	m21 := m.Data[2][1]
	m22 := m.Data[2][2]
	m23 := m.Data[2][3]

	m30 := m.Data[3][0]
	m31 := m.Data[3][1]
	m32 := m.Data[3][2]
	m33 := m.Data[3][3]

	m.Data = [4][4]float32{
		{
			m00*m2.Data[0][0] + m10*m2.Data[0][1] + m20*m2.Data[0][2] + m30*m2.Data[0][3],
			m01*m2.Data[0][0] + m11*m2.Data[0][1] + m21*m2.Data[0][2] + m31*m2.Data[0][3],
			m02*m2.Data[0][0] + m12*m2.Data[0][1] + m22*m2.Data[0][2] + m32*m2.Data[0][3],
			m03*m2.Data[0][0] + m13*m2.Data[0][1] + m23*m2.Data[0][2] + m33*m2.Data[0][3],
		},
		{
			m00*m2.Data[1][0] + m10*m2.Data[1][1] + m20*m2.Data[1][2] + m30*m2.Data[1][3],
			m01*m2.Data[1][0] + m11*m2.Data[1][1] + m21*m2.Data[1][2] + m31*m2.Data[1][3],
			m02*m2.Data[1][0] + m12*m2.Data[1][1] + m22*m2.Data[1][2] + m32*m2.Data[1][3],
			m03*m2.Data[1][0] + m13*m2.Data[1][1] + m23*m2.Data[1][2] + m33*m2.Data[1][3],
		},
		{
			m00*m2.Data[2][0] + m10*m2.Data[2][1] + m20*m2.Data[2][2] + m30*m2.Data[2][3],
			m01*m2.Data[2][0] + m11*m2.Data[2][1] + m21*m2.Data[2][2] + m31*m2.Data[2][3],
			m02*m2.Data[2][0] + m12*m2.Data[2][1] + m22*m2.Data[2][2] + m32*m2.Data[2][3],
			m03*m2.Data[2][0] + m13*m2.Data[2][1] + m23*m2.Data[2][2] + m33*m2.Data[2][3],
		},
		{
			m00*m2.Data[3][0] + m10*m2.Data[3][1] + m20*m2.Data[3][2] + m30*m2.Data[3][3],
			m01*m2.Data[3][0] + m11*m2.Data[3][1] + m21*m2.Data[3][2] + m31*m2.Data[3][3],
			m02*m2.Data[3][0] + m12*m2.Data[3][1] + m22*m2.Data[3][2] + m32*m2.Data[3][3],
			m03*m2.Data[3][0] + m13*m2.Data[3][1] + m23*m2.Data[3][2] + m33*m2.Data[3][3],
		},
	}

	return m
}

//Scale m *= x (element wise multiplication)
func (m *Mat4) Scale(x float32) *Mat4 {

	m.Data[0][0] *= x
	m.Data[0][1] *= x
	m.Data[0][2] *= x
	m.Data[0][3] *= x

	m.Data[1][0] *= x
	m.Data[1][1] *= x
	m.Data[1][2] *= x
	m.Data[1][3] *= x

	m.Data[2][0] *= x
	m.Data[2][1] *= x
	m.Data[2][2] *= x
	m.Data[2][3] *= x

	m.Data[3][0] *= x
	m.Data[3][1] *= x
	m.Data[3][2] *= x
	m.Data[3][3] *= x
	return m
}

func (v *Mat4) Clone() *Mat4 {
	return &Mat4{Data: v.Data}
}

func (m *Mat4) Eq(m2 *Mat4) bool {
	return m.Data == m2.Data
}

//AddMat4 m3 = m1 + m2
func AddMat4(m1, m2 *Mat4) *Mat4 {
	return &Mat4{
		Data: [4][4]float32{
			{
				m1.Data[0][0] + m2.Data[0][0],
				m1.Data[0][1] + m2.Data[0][1],
				m1.Data[0][2] + m2.Data[0][2],
				m1.Data[0][3] + m2.Data[0][3],
			},
			{
				m1.Data[1][0] + m2.Data[1][0],
				m1.Data[1][1] + m2.Data[1][1],
				m1.Data[1][2] + m2.Data[1][2],
				m1.Data[1][3] + m2.Data[1][3],
			},
			{
				m1.Data[2][0] + m2.Data[2][0],
				m1.Data[2][1] + m2.Data[2][1],
				m1.Data[2][2] + m2.Data[2][2],
				m1.Data[2][3] + m2.Data[2][3],
			},
			{
				m1.Data[3][0] + m2.Data[3][0],
				m1.Data[3][1] + m2.Data[3][1],
				m1.Data[3][2] + m2.Data[3][2],
				m1.Data[3][3] + m2.Data[3][3],
			},
		},
	}
}

//SubMat4 m3 = m1 - m2
func SubMat4(m1, m2 *Mat4) *Mat4 {
	return &Mat4{
		Data: [4][4]float32{
			{
				m1.Data[0][0] - m2.Data[0][0],
				m1.Data[0][1] - m2.Data[0][1],
				m1.Data[0][2] - m2.Data[0][2],
				m1.Data[0][3] - m2.Data[0][3],
			},
			{
				m1.Data[1][0] - m2.Data[1][0],
				m1.Data[1][1] - m2.Data[1][1],
				m1.Data[1][2] - m2.Data[1][2],
				m1.Data[1][3] - m2.Data[1][3],
			},
			{
				m1.Data[2][0] - m2.Data[2][0],
				m1.Data[2][1] - m2.Data[2][1],
				m1.Data[2][2] - m2.Data[2][2],
				m1.Data[2][3] - m2.Data[2][3],
			},
			{
				m1.Data[3][0] - m2.Data[3][0],
				m1.Data[3][1] - m2.Data[3][1],
				m1.Data[3][2] - m2.Data[3][2],
				m1.Data[3][3] - m2.Data[3][3],
			},
		},
	}
}

//MulMat4 m3 = m1 * m2
func MulMat4(m1, m2 *Mat4) *Mat4 {

	m00 := m1.Data[0][0]
	m01 := m1.Data[0][1]
	m02 := m1.Data[0][2]
	m03 := m1.Data[0][3]

	m10 := m1.Data[1][0]
	m11 := m1.Data[1][1]
	m12 := m1.Data[1][2]
	m13 := m1.Data[1][3]

	m20 := m1.Data[2][0]
	m21 := m1.Data[2][1]
	m22 := m1.Data[2][2]
	m23 := m1.Data[2][3]

	m30 := m1.Data[3][0]
	m31 := m1.Data[3][1]
	m32 := m1.Data[3][2]
	m33 := m1.Data[3][3]

	return &Mat4{
		Data: [4][4]float32{
			{
				m00*m2.Data[0][0] + m10*m2.Data[0][1] + m20*m2.Data[0][2] + m30*m2.Data[0][3],
				m01*m2.Data[0][0] + m11*m2.Data[0][1] + m21*m2.Data[0][2] + m31*m2.Data[0][3],
				m02*m2.Data[0][0] + m12*m2.Data[0][1] + m22*m2.Data[0][2] + m32*m2.Data[0][3],
				m03*m2.Data[0][0] + m13*m2.Data[0][1] + m23*m2.Data[0][2] + m33*m2.Data[0][3],
			},
			{
				m00*m2.Data[1][0] + m10*m2.Data[1][1] + m20*m2.Data[1][2] + m30*m2.Data[1][3],
				m01*m2.Data[1][0] + m11*m2.Data[1][1] + m21*m2.Data[1][2] + m31*m2.Data[1][3],
				m02*m2.Data[1][0] + m12*m2.Data[1][1] + m22*m2.Data[1][2] + m32*m2.Data[1][3],
				m03*m2.Data[1][0] + m13*m2.Data[1][1] + m23*m2.Data[1][2] + m33*m2.Data[1][3],
			},
			{
				m00*m2.Data[2][0] + m10*m2.Data[2][1] + m20*m2.Data[2][2] + m30*m2.Data[2][3],
				m01*m2.Data[2][0] + m11*m2.Data[2][1] + m21*m2.Data[2][2] + m31*m2.Data[2][3],
				m02*m2.Data[2][0] + m12*m2.Data[2][1] + m22*m2.Data[2][2] + m32*m2.Data[2][3],
				m03*m2.Data[2][0] + m13*m2.Data[2][1] + m23*m2.Data[2][2] + m33*m2.Data[2][3],
			},
			{
				m00*m2.Data[3][0] + m10*m2.Data[3][1] + m20*m2.Data[3][2] + m30*m2.Data[3][3],
				m01*m2.Data[3][0] + m11*m2.Data[3][1] + m21*m2.Data[3][2] + m31*m2.Data[3][3],
				m02*m2.Data[3][0] + m12*m2.Data[3][1] + m22*m2.Data[3][2] + m32*m2.Data[3][3],
				m03*m2.Data[3][0] + m13*m2.Data[3][1] + m23*m2.Data[3][2] + m33*m2.Data[3][3],
			},
		},
	}
}

//MulMat4Vec4 v2 = m1 * v1
func MulMat4Vec4(m1 *Mat4, v1 *Vec4) *Vec4 {
	return &Vec4{
		Data: [4]float32{
			m1.Data[0][0]*v1.Data[0] + m1.Data[1][0]*v1.Data[1] + m1.Data[2][0]*v1.Data[2] + m1.Data[3][0]*v1.Data[3],
			m1.Data[0][1]*v1.Data[0] + m1.Data[1][1]*v1.Data[1] + m1.Data[2][1]*v1.Data[2] + m1.Data[3][1]*v1.Data[3],
			m1.Data[0][2]*v1.Data[0] + m1.Data[1][2]*v1.Data[1] + m1.Data[2][2]*v1.Data[2] + m1.Data[3][2]*v1.Data[3],
			m1.Data[0][3]*v1.Data[0] + m1.Data[1][3]*v1.Data[1] + m1.Data[2][3]*v1.Data[2] + m1.Data[3][3]*v1.Data[3],
		},
	}
}

//NewMat4Id returns the 4x4 identity matrix
func NewMat4Id() *Mat4 {
	return &Mat4{
		Data: [4][4]float32{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}
