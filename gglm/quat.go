package gglm

import (
	"fmt"
)

var _ Swizzle4 = &Quat{}
var _ fmt.Stringer = &Quat{}

type Quat struct {
	Vec4
}

//Eq checks for exact equality
func (q *Quat) Eq(q2 *Quat) bool {
	return q.Data == q2.Data
}

//Angle returns the angle represented by this quaternion in radians
func (q *Quat) Angle() float32 {

	if Abs32(q.Data[3]) > CosHalf {

		a := Asin32(Sqrt32(q.Data[0]*q.Data[0]+q.Data[1]*q.Data[1]+q.Data[2]*q.Data[2])) * 2
		if q.Data[3] < 0 {
			return Pi*2 - a
		}
		return a
	}

	return Acos32(q.Data[3]) * 2
}

//Axis returns the rotation axis represented by this quaternion
func (q *Quat) Axis() *Vec3 {

	var t float32 = 1 - q.Data[3]*q.Data[3]
	if t <= 0 {
		return &Vec3{Data: [3]float32{0, 0, 1}}
	}

	t = 1 / Sqrt32(t)
	return &Vec3{Data: [3]float32{
		q.Data[0] * t,
		q.Data[1] * t,
		q.Data[2] * t,
	}}
}

//Euler takes rotations in radians and produces a rotation that
//rotates around the z-axis, y-axis and lastly x-axis.
func NewQuatEuler(v *Vec3) *Quat {

	//Some other common terminology: x=roll, y=pitch, z=yaw
	sinX, cosX := Sincos32(v.Data[0] * 0.5)
	sinY, cosY := Sincos32(v.Data[1] * 0.5)
	sinZ, cosZ := Sincos32(v.Data[2] * 0.5)

	//This produces a z->y->x multiply order, but its written as XYZ.
	//This is due to XYZ meaning independent rotation matrices, so Z is applied
	//first, then Y matrix and lastly X.
	//See this for more info: https://github.com/godotengine/godot/issues/6816#issuecomment-254592170
	//
	//Note: On most conversion tools putting the multiply order (e.g. ZYX for us) is required.
	return &Quat{
		Vec4: Vec4{
			Data: [4]float32{
				sinX*cosY*cosZ - cosX*sinY*sinZ,
				cosX*sinY*cosZ + sinX*cosY*sinZ,
				cosX*cosY*sinZ - sinX*sinY*cosZ,
				cosX*cosY*cosZ + sinX*sinY*sinZ,
			},
		},
	}
}

//Euler takes rotations in radians and produces a rotation that
//rotates around the z-axis, y-axis and lastly x-axis.
func NewQuatEulerXYZ(x, y, z float32) *Quat {

	//Some other common terminology: x=roll, y=pitch, z=yaw
	sinX, cosX := Sincos32(x * 0.5)
	sinY, cosY := Sincos32(y * 0.5)
	sinZ, cosZ := Sincos32(z * 0.5)

	//This produces a z->y->x multiply order, but its written as XYZ.
	//This is due to XYZ meaning independent rotation matrices, so Z is applied
	//first, then Y matrix and lastly X.
	//See this for more info: https://github.com/godotengine/godot/issues/6816#issuecomment-254592170
	//
	//Note: On most conversion tools putting the multiply order (e.g. ZYX for us) is required.
	return &Quat{
		Vec4: Vec4{
			Data: [4]float32{
				sinX*cosY*cosZ - cosX*sinY*sinZ,
				cosX*sinY*cosZ + sinX*cosY*sinZ,
				cosX*cosY*sinZ - sinX*sinY*cosZ,
				cosX*cosY*cosZ + sinX*sinY*sinZ,
			},
		},
	}
}

//NewQuatAngleAxis produces a quaternion thats rotates rotRad radians around the *normalized* vector rotAxisNorm
func NewQuatAngleAxis(rotRad float32, rotAxisNorm *Vec3) *Quat {

	s, c := Sincos32(rotRad * 0.5)
	return &Quat{
		Vec4: Vec4{
			Data: [4]float32{
				rotAxisNorm.Data[0] * s,
				rotAxisNorm.Data[1] * s,
				rotAxisNorm.Data[2] * s,
				c,
			},
		},
	}
}

func NewQuatId() *Quat {
	return &Quat{
		Vec4: Vec4{
			Data: [4]float32{0, 0, 0, 1},
		},
	}
}
