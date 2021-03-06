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

func (v *Vec4) SetX(f float32) {
	v.Data[0] = f
}

func (v *Vec4) SetR(f float32) {
	v.Data[0] = f
}

func (v *Vec4) SetY(f float32) {
	v.Data[1] = f
}

func (v *Vec4) SetG(f float32) {
	v.Data[1] = f
}

func (v *Vec4) SetZ(f float32) {
	v.Data[2] = f
}

func (v *Vec4) SetB(f float32) {
	v.Data[2] = f
}

func (v *Vec4) SetW(f float32) {
	v.Data[3] = f
}

func (v *Vec4) SetA(f float32) {
	v.Data[3] = f
}

func (v *Vec4) AddX(x float32) {
	v.Data[0] += x
}

func (v *Vec4) AddY(y float32) {
	v.Data[1] += y
}

func (v *Vec4) AddZ(z float32) {
	v.Data[2] += z
}

func (v *Vec4) AddW(w float32) {
	v.Data[3] += w
}

func (v *Vec4) AddR(r float32) {
	v.Data[0] += r
}

func (v *Vec4) AddG(g float32) {
	v.Data[1] += g
}

func (v *Vec4) AddB(b float32) {
	v.Data[2] += b
}

func (v *Vec4) AddA(a float32) {
	v.Data[3] += a
}

func (v *Vec4) SetXY(x, y float32) {
	v.Data[0] = x
	v.Data[1] = y
}

func (v *Vec4) AddXY(x, y float32) {
	v.Data[0] += x
	v.Data[1] += y
}

func (v *Vec4) SetRG(r, g float32) {
	v.Data[0] = r
	v.Data[1] = g
}

func (v *Vec4) AddRG(r, g float32) {
	v.Data[0] += r
	v.Data[1] += g
}

func (v *Vec4) SetXYZ(x, y, z float32) {
	v.Data[0] = x
	v.Data[1] = y
	v.Data[2] = z
}

func (v *Vec4) AddXYZ(x, y, z float32) {
	v.Data[0] += x
	v.Data[1] += y
	v.Data[2] += z
}

func (v *Vec4) SetRGB(r, g, b float32) {
	v.Data[0] = r
	v.Data[1] = g
	v.Data[2] = b
}

func (v *Vec4) AddRGB(r, g, b float32) {
	v.Data[0] += r
	v.Data[1] += g
	v.Data[2] += b
}

func (v *Vec4) SetXYZW(x, y, z, w float32) {
	v.Data[0] = x
	v.Data[1] = y
	v.Data[2] = z
	v.Data[3] = w
}

func (v *Vec4) AddXYZW(x, y, z, w float32) {
	v.Data[0] += x
	v.Data[1] += y
	v.Data[2] += z
	v.Data[3] += w
}

func (v *Vec4) SetRGBA(r, g, b, a float32) {
	v.Data[0] = r
	v.Data[1] = g
	v.Data[2] = b
	v.Data[3] = a
}

func (v *Vec4) AddRGBA(r, g, b, a float32) {
	v.Data[0] += r
	v.Data[1] += g
	v.Data[2] += b
	v.Data[3] += a
}

func (v *Vec4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.X(), v.Y(), v.Z(), v.W())
}

//Scale v *= x (element wise multiplication)
func (v *Vec4) Scale(x float32) *Vec4 {
	v.Data[0] *= x
	v.Data[1] *= x
	v.Data[2] *= x
	v.Data[3] *= x
	return v
}

func (v *Vec4) Add(v2 *Vec4) *Vec4 {
	v.Data[0] += v2.X()
	v.Data[1] += v2.Y()
	v.Data[2] += v2.Z()
	v.Data[3] += v2.W()
	return v
}

//SubVec4 v -= v2
func (v *Vec4) Sub(v2 *Vec4) *Vec4 {
	v.Data[0] -= v2.X()
	v.Data[1] -= v2.Y()
	v.Data[2] -= v2.Z()
	v.Data[3] -= v2.W()
	return v
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

func (v *Vec4) Set(x, y, z, w float32) {
	v.Data[0] = x
	v.Data[1] = y
	v.Data[2] = z
	v.Data[3] = w
}

func (v *Vec4) Normalize() {
	mag := float32(math.Sqrt(float64(v.Data[0]*v.Data[0] + v.Data[1]*v.Data[1] + v.Data[2]*v.Data[2] + v.Data[3]*v.Data[3])))
	v.Data[0] /= mag
	v.Data[1] /= mag
	v.Data[2] /= mag
	v.Data[3] /= mag
}

func (v *Vec4) Clone() *Vec4 {
	return &Vec4{Data: v.Data}
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

func NewVec4(x, y, z, w float32) *Vec4 {
	return &Vec4{
		[4]float32{
			x,
			y,
			z,
			w,
		},
	}
}
