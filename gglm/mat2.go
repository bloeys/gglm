package gglm

import (
	"fmt"
)

var _ Mat = &Mat2{}
var _ fmt.Stringer = &Mat2{}

type Mat2 struct {
	Data [2][2]float32
}

func (m *Mat2) Get(row, col int) float32 {
	return m.Data[col][row]
}

func (m *Mat2) Set(row, col int, val float32) {
	m.Data[col][row] = val
}

func (m *Mat2) Size() MatSize {
	return MatSize2x2
}

func (m *Mat2) String() string {
	//+ always shows +/- sign; - means pad to the right; 9 means total of 9 digits (or padding if less); .3 means 3 decimals
	return fmt.Sprintf("\n| %+-9.3f %+-9.3f |\n| %+-9.3f %+-9.3f |\n", m.Data[0][0], m.Data[0][1], m.Data[1][0], m.Data[1][1])
}

func (m *Mat2) Col(c int) *Vec2 {
	return &Vec2{Data: m.Data[c]}
}

//Add m += m2
func (m *Mat2) Add(m2 *Mat2) *Mat2 {
	m.Data[0][0] += m2.Data[0][0]
	m.Data[0][1] += m2.Data[0][1]
	m.Data[1][0] += m2.Data[1][0]
	m.Data[1][1] += m2.Data[1][1]
	return m
}

//Add m -= m2
func (m *Mat2) Sub(m2 *Mat2) *Mat2 {
	m.Data[0][0] -= m2.Data[0][0]
	m.Data[0][1] -= m2.Data[0][1]
	m.Data[1][0] -= m2.Data[1][0]
	m.Data[1][1] -= m2.Data[1][1]
	return m
}

//Mul m *= m2
func (m1 *Mat2) Mul(m2 *Mat2) *Mat2 {
	m1.Data = [2][2]float32{
		{
			m1.Data[0][0]*m2.Data[0][0] + m1.Data[1][0]*m2.Data[0][1],
			m1.Data[0][1]*m2.Data[0][0] + m1.Data[1][1]*m2.Data[0][1],
		},
		{
			m1.Data[0][0]*m2.Data[1][0] + m1.Data[1][0]*m2.Data[1][1],
			m1.Data[0][1]*m2.Data[1][0] + m1.Data[1][1]*m2.Data[1][1],
		},
	}

	return m1
}

//Scale m *= x (element wise multiplication)
func (m *Mat2) Scale(x float32) *Mat2 {
	m.Data[0][0] *= x
	m.Data[0][1] *= x
	m.Data[1][0] *= x
	m.Data[1][1] *= x
	return m
}

func (v *Mat2) Clone() *Mat2 {
	return &Mat2{Data: v.Data}
}

func (m *Mat2) Eq(m2 *Mat2) bool {
	return m.Data == m2.Data
}

//AddMat2 m3 = m1 + m2
func AddMat2(m1, m2 *Mat2) *Mat2 {
	return &Mat2{
		Data: [2][2]float32{
			{
				m1.Data[0][0] + m2.Data[0][0],
				m1.Data[0][1] + m2.Data[0][1],
			},
			{
				m1.Data[1][0] + m2.Data[1][0],
				m1.Data[1][1] + m2.Data[1][1],
			},
		},
	}
}

//SubMat2 m3 = m1 - m2
func SubMat2(m1, m2 *Mat2) *Mat2 {
	return &Mat2{
		Data: [2][2]float32{
			{
				m1.Data[0][0] - m2.Data[0][0],
				m1.Data[0][1] - m2.Data[0][1],
			},
			{
				m1.Data[1][0] - m2.Data[1][0],
				m1.Data[1][1] - m2.Data[1][1],
			},
		},
	}
}

//MulMat2 m3 = m1 * m2
func MulMat2(m1, m2 *Mat2) *Mat2 {
	return &Mat2{
		Data: [2][2]float32{
			{
				m1.Data[0][0]*m2.Data[0][0] + m1.Data[1][0]*m2.Data[0][1],
				m1.Data[0][1]*m2.Data[0][0] + m1.Data[1][1]*m2.Data[0][1],
			},
			{
				m1.Data[0][0]*m2.Data[1][0] + m1.Data[1][0]*m2.Data[1][1],
				m1.Data[0][1]*m2.Data[1][0] + m1.Data[1][1]*m2.Data[1][1],
			},
		},
	}
}

//MulMat2Vec2 v2 = m1 * v1
func MulMat2Vec2(m1 *Mat2, v1 *Vec2) *Vec2 {
	return &Vec2{
		Data: [2]float32{
			m1.Data[0][0]*v1.Data[0] + m1.Data[1][0]*v1.Data[1],
			m1.Data[0][1]*v1.Data[0] + m1.Data[1][1]*v1.Data[1],
		},
	}
}

//NewMat2Id returns the 2x2 identity matrix
func NewMat2Id() *Mat2 {
	return &Mat2{
		Data: [2][2]float32{
			{1, 0},
			{0, 1},
		},
	}
}
