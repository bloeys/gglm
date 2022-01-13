package gglm

import (
	"fmt"
	"math"
)

var _ Mat = &TrMat{}
var _ fmt.Stringer = &TrMat{}

//TrMat represents a transformation matrix
type TrMat struct {
	Mat4
}

//Translate adds v to the translation components of the transformation matrix
func (t *TrMat) Translate(v *Vec3) {
	t.Data[3][0] += v.Data[0]
	t.Data[3][1] += v.Data[1]
	t.Data[3][2] += v.Data[2]
}

//Scale multiplies the scale components of the transformation matrix by v
func (t *TrMat) Scale(v *Vec3) {
	t.Data[0][0] *= v.Data[0]
	t.Data[1][1] *= v.Data[1]
	t.Data[2][2] *= v.Data[2]
}

//Rotate takes a *normalized* axis and angles in radians to rotate around the given axis
func (t *TrMat) Rotate(rads float32, axis *Vec3) {

	s := Sin32(rads)
	c := Cos32(rads)

	axis = axis.Normalize()
	temp := axis.Clone().Scale(1 - c)

	rotate := TrMat{}
	rotate.Data[0][0] = c + temp.Data[0]*axis.Data[0]
	rotate.Data[0][1] = temp.Data[0]*axis.Data[1] + s*axis.Data[2]
	rotate.Data[0][2] = temp.Data[0]*axis.Data[2] - s*axis.Data[1]

	rotate.Data[1][0] = temp.Data[1]*axis.Data[0] - s*axis.Data[2]
	rotate.Data[1][1] = c + temp.Data[1]*axis.Data[1]
	rotate.Data[1][2] = temp.Data[1]*axis.Data[2] + s*axis.Data[0]

	rotate.Data[2][0] = temp.Data[2]*axis.Data[0] + s*axis.Data[1]
	rotate.Data[2][1] = temp.Data[2]*axis.Data[1] - s*axis.Data[0]
	rotate.Data[2][2] = c + temp.Data[2]*axis.Data[2]

	result := &Mat4{}
	result.Data[0] = t.Col(0).Scale(rotate.Data[0][0]).
		Add(t.Col(1).Scale(rotate.Data[0][1])).
		Add(t.Col(2).Scale(rotate.Data[0][2])).
		Data

	result.Data[1] = t.Col(0).Scale(rotate.Data[1][0]).
		Add(t.Col(1).Scale(rotate.Data[1][1])).
		Add(t.Col(2).Scale(rotate.Data[1][2])).
		Data

	result.Data[2] = t.Col(0).Scale(rotate.Data[2][0]).
		Add(t.Col(1).Scale(rotate.Data[2][1])).
		Add(t.Col(2).Scale(rotate.Data[2][2])).
		Data

	t.Data[0] = result.Data[0]
	t.Data[1] = result.Data[1]
	t.Data[2] = result.Data[2]
}

func (t *TrMat) Mul(m *TrMat) *TrMat {
	t.Mat4.Mul(&m.Mat4)
	return t
}

func (t *TrMat) Eq(m *TrMat) bool {
	return t.Data == m.Data
}

func (t *TrMat) Clone() *TrMat {
	return &TrMat{
		Mat4: *t.Mat4.Clone(),
	}
}

func NewTranslationMat(v *Vec3) *TrMat {
	return &TrMat{
		Mat4: Mat4{
			Data: [4][4]float32{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{v.Data[0], v.Data[1], v.Data[2], 1},
			},
		},
	}
}

func NewScaleMat(v *Vec3) *TrMat {
	return &TrMat{
		Mat4: Mat4{
			Data: [4][4]float32{
				{v.Data[0], 0, 0, 0},
				{0, v.Data[1], 0, 0},
				{0, 0, v.Data[2], 0},
				{0, 0, 0, 1},
			},
		},
	}
}

func NewRotMat(q *Quat) *TrMat {

	//Based on: https://www.weizmann.ac.il/sci-tea/benari/sites/sci-tea.benari/files/uploads/softwareAndLearningMaterials/quaternion-tutorial-2-0-1.pdf
	//Note: in the reference p0,p1,p2,p3 == w,x,y,z

	xx := q.Data[0] * q.Data[0]
	yy := q.Data[1] * q.Data[1]
	zz := q.Data[2] * q.Data[2]
	ww := q.Data[3] * q.Data[3]

	xy := q.Data[0] * q.Data[1]
	xz := q.Data[0] * q.Data[2]
	xw := q.Data[0] * q.Data[3]

	yz := q.Data[1] * q.Data[2]
	yw := q.Data[1] * q.Data[3]

	zw := q.Data[2] * q.Data[3]

	return &TrMat{
		Mat4: Mat4{
			Data: [4][4]float32{
				{2*(ww+xx) - 1, 2 * (zw + xy), 2 * (xz - yw), 0},
				{2 * (xy - zw), 2*(ww+yy) - 1, 2 * (xw + yz), 0},
				{2 * (yw + xz), 2 * (yz - xw), 2*(ww+zz) - 1, 0},
				{0, 0, 0, 1},
			},
		},
	}
}

func LookAt(pos, targetPos, worldUp *Vec3) *TrMat {

	forward := SubVec3(targetPos, pos).Normalize()
	right := Cross(worldUp, forward).Normalize()
	up := Cross(forward, right)

	return &TrMat{
		Mat4: Mat4{
			Data: [4][4]float32{
				{right.Data[0], up.Data[0], forward.Data[0], 0},
				{right.Data[1], up.Data[1], forward.Data[1], 0},
				{right.Data[2], up.Data[2], forward.Data[2], 0},
				{-DotVec3(pos, right), -DotVec3(pos, up), DotVec3(pos, forward), 1},
			},
		},
	}
}

//Perspective creates a perspective projection matrix
func Perspective(fov, aspectRatio, nearClip, farClip float32) *Mat4 {
	halfFovTan := float32(math.Tan(float64(fov * 0.5)))
	return &Mat4{
		Data: [4][4]float32{
			{1 / (aspectRatio * halfFovTan), 0, 0, 0},
			{0, 1 / halfFovTan, 0, 0},
			{0, 0, -(nearClip + farClip) / (farClip - nearClip), -1},
			{0, 0, -(2 * farClip * nearClip) / (farClip - nearClip), 0},
		},
	}
}

//Perspective creates an orthographic projection matrix
func Ortho(left, right, top, bottom, nearClip, farClip float32) *TrMat {
	return &TrMat{
		Mat4: Mat4{
			Data: [4][4]float32{
				{2 / (right - left), 0, 0, 0},
				{0, 2 / (top - bottom), 0, 0},
				{0, 0, -2 / (farClip - nearClip), 0},
				{-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(farClip + nearClip) / (farClip - nearClip), 1},
			},
		},
	}
}

func NewTrMatId() *TrMat {
	return &TrMat{
		Mat4: *NewMat4Id(),
	}
}
