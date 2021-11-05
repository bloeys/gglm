package gglm

import (
	"fmt"
	"math"
)

var _ Swizzle4 = &Vec4{}
var _ fmt.Stringer = &Vec4{}

type Vec4 struct {
	Data [4]float32
}

func (v *Vec4) X() float32 {
	return v.Data[0]
}

func (v *Vec4) Y() float32 {
	return v.Data[1]
}

func (v *Vec4) Z() float32 {
	return v.Data[2]
}

func (v *Vec4) W() float32 {
	return v.Data[3]
}

func (v *Vec4) R() float32 {
	return v.Data[0]
}

func (v *Vec4) G() float32 {
	return v.Data[1]
}

func (v *Vec4) B() float32 {
	return v.Data[2]
}

func (v *Vec4) A() float32 {
	return v.Data[3]
}

func (v *Vec4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.X(), v.Y(), v.Z(), v.W())
}

//Scale v *= x (element wise multiplication)
func (v *Vec4) Scale(x float32) {
	v.Data[0] *= x
	v.Data[1] *= x
	v.Data[2] *= x
	v.Data[3] *= x
}

func (v *Vec4) Add(v2 *Vec4) {

	v.Data[0] += v2.X()
	v.Data[1] += v2.Y()
	v.Data[2] += v2.Z()
	v.Data[3] += v2.W()
}

//SubVec4 v -= v2
func (v *Vec4) Sub(v2 *Vec4) {
	v.Data[0] -= v2.X()
	v.Data[1] -= v2.Y()
	v.Data[2] -= v2.Z()
	v.Data[3] -= v2.W()
}

//Mag returns the magnitude of the vector
func (v *Vec4) Mag() float32 {
	return float32(math.Sqrt(float64(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z() + v.W()*v.W())))
}

//Mag returns the squared magnitude of the vector
func (v *Vec4) SqrMag() float32 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z() + v.Z()*v.Z()
}

func (v *Vec4) Eq(v2 *Vec4) bool {
	return v.Data == v2.Data
}

//AddVec4 v3 = v1 + v2
func AddVec4(v1, v2 *Vec4) *Vec4 {
	return &Vec4{
		Data: [4]float32{
			v1.X() + v2.X(),
			v1.Y() + v2.Y(),
			v1.Z() + v2.Z(),
			v1.W() + v2.W(),
		},
	}
}

//SubVec4 v3 = v1 - v2
func SubVec4(v1, v2 *Vec4) *Vec4 {
	return &Vec4{
		Data: [4]float32{
			v1.X() - v2.X(),
			v1.Y() - v2.Y(),
			v1.Z() - v2.Z(),
			v1.W() - v2.W(),
		},
	}
}
