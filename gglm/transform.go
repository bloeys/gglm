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

//Translate adds the vector to the translation components of the transformation matrix
func (t *TrMat) Translate(v *Vec3) {
	t.Data[3][0] += v.Data[0]
	t.Data[3][1] += v.Data[1]
	t.Data[3][2] += v.Data[2]
}

//Scale multiplies the vector by the scale components of the transformation matrix
func (t *TrMat) Scale(v *Vec3) {
	t.Data[0][0] *= v.Data[0]
	t.Data[1][1] *= v.Data[1]
	t.Data[2][2] *= v.Data[2]
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

func NewTrMatId() *TrMat {
	return &TrMat{
		Mat4: *NewMat4Id(),
	}
}
