package gglm

import (
	"fmt"
	"math"
)

var _ Swizzle3 = &Vec3{}
var _ fmt.Stringer = &Vec3{}

type Vec3 struct {
	Data [3]float32
}

func (v *Vec3) X() float32 {
	return v.Data[0]
}

func (v *Vec3) Y() float32 {
	return v.Data[1]
}

func (v *Vec3) Z() float32 {
	return v.Data[2]
}

func (v *Vec3) R() float32 {
	return v.Data[0]
}

func (v *Vec3) G() float32 {
	return v.Data[1]
}

func (v *Vec3) B() float32 {
	return v.Data[2]
}

func (v *Vec3) SetX(f float32) {
	v.Data[0] = f
}

func (v *Vec3) SetR(f float32) {
	v.Data[0] = f
}

func (v *Vec3) SetY(f float32) {
	v.Data[1] = f
}

func (v *Vec3) SetG(f float32) {
	v.Data[1] = f
}

func (v *Vec3) SetZ(f float32) {
	v.Data[2] = f
}

func (v *Vec3) SetB(f float32) {
	v.Data[2] = f
}

func (v *Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X(), v.Y(), v.Z())
}

//Scale v *= x (element wise multiplication)
func (v *Vec3) Scale(x float32) *Vec3 {
	v.Data[0] *= x
	v.Data[1] *= x
	v.Data[2] *= x
	return v
}

func (v *Vec3) Add(v2 *Vec3) *Vec3 {

	v.Data[0] += v2.X()
	v.Data[1] += v2.Y()
	v.Data[2] += v2.Z()
	return v
}

//SubVec3 v -= v2
func (v *Vec3) Sub(v2 *Vec3) *Vec3 {
	v.Data[0] -= v2.X()
	v.Data[1] -= v2.Y()
	v.Data[2] -= v2.Z()
	return v
}

//Mag returns the magnitude of the vector
func (v *Vec3) Mag() float32 {
	return float32(math.Sqrt(float64(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())))
}

//Mag returns the squared magnitude of the vector
func (v *Vec3) SqrMag() float32 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

func (v *Vec3) Eq(v2 *Vec3) bool {
	return v.Data == v2.Data
}

func (v *Vec3) Set(x, y, z float32) {
	v.Data[0] = x
	v.Data[1] = y
	v.Data[2] = z
}

//Normalize normalizes this vector and returns it (doesn't copy)
func (v *Vec3) Normalize() *Vec3 {
	mag := float32(math.Sqrt(float64(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())))
	v.Data[0] /= mag
	v.Data[1] /= mag
	v.Data[2] /= mag

	return v
}

func (v *Vec3) Clone() *Vec3 {
	return &Vec3{Data: v.Data}
}

//AsRad returns a new vector with all values converted to Radians (i.e. multiplied by gglm.Deg2Rad)
func (v *Vec3) AsRad() *Vec3 {
	return &Vec3{
		Data: [3]float32{
			v.Data[0] * Deg2Rad,
			v.Data[1] * Deg2Rad,
			v.Data[2] * Deg2Rad,
		},
	}
}

//AddVec3 v3 = v1 + v2
func AddVec3(v1, v2 *Vec3) *Vec3 {
	return &Vec3{
		Data: [3]float32{
			v1.X() + v2.X(),
			v1.Y() + v2.Y(),
			v1.Z() + v2.Z(),
		},
	}
}

//SubVec3 v3 = v1 - v2
func SubVec3(v1, v2 *Vec3) *Vec3 {
	return &Vec3{
		Data: [3]float32{
			v1.X() - v2.X(),
			v1.Y() - v2.Y(),
			v1.Z() - v2.Z(),
		},
	}
}

func NewVec3(x, y, z float32) *Vec3 {
	return &Vec3{
		[3]float32{
			x,
			y,
			z,
		},
	}
}
