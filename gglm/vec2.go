package gglm

//Note: We don't use the Swizzle interface for add/sub because the interface doesn't allow inling :(

import (
	"fmt"
	"math"
)

var _ Swizzle2 = &Vec2{}
var _ fmt.Stringer = &Vec2{}

type Vec2 struct {
	Data [2]float32
}

func (v *Vec2) X() float32 {
	return v.Data[0]
}

func (v *Vec2) Y() float32 {
	return v.Data[1]
}

func (v *Vec2) R() float32 {
	return v.Data[0]
}

func (v *Vec2) G() float32 {
	return v.Data[1]
}

func (v *Vec2) SetX(x float32) {
	v.Data[0] = x
}

func (v *Vec2) SetR(r float32) {
	v.Data[0] = r
}

func (v *Vec2) SetY(y float32) {
	v.Data[1] = y
}

func (v *Vec2) SetG(g float32) {
	v.Data[1] = g
}

func (v *Vec2) AddX(x float32) {
	v.Data[0] += x
}

func (v *Vec2) AddY(y float32) {
	v.Data[1] += y
}

func (v *Vec2) AddR(r float32) {
	v.Data[0] += r
}

func (v *Vec2) AddG(g float32) {
	v.Data[1] += g
}

func (v *Vec2) SetXY(x, y float32) {
	v.Data[0] = x
	v.Data[1] = y
}

func (v *Vec2) AddXY(x, y float32) {
	v.Data[0] += x
	v.Data[1] += y
}

func (v *Vec2) SetRG(r, g float32) {
	v.Data[0] = r
	v.Data[1] = g
}

func (v *Vec2) AddRG(r, g float32) {
	v.Data[0] += r
	v.Data[1] += g
}

func (v *Vec2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X(), v.Y())
}

//Scale v *= x (element wise multiplication)
func (v *Vec2) Scale(x float32) *Vec2 {
	v.Data[0] *= x
	v.Data[1] *= x
	return v
}

//Add v += v2
func (v *Vec2) Add(v2 *Vec2) *Vec2 {
	v.Data[0] += v2.X()
	v.Data[1] += v2.Y()
	return v
}

//SubVec2 v -= v2
func (v *Vec2) Sub(v2 *Vec2) *Vec2 {
	v.Data[0] -= v2.X()
	v.Data[1] -= v2.Y()
	return v
}

//Mag returns the magnitude of the vector
func (v *Vec2) Mag() float32 {
	return float32(math.Sqrt(float64(v.X()*v.X() + v.Y()*v.Y())))
}

//Mag returns the squared magnitude of the vector
func (v *Vec2) SqrMag() float32 {
	return v.X()*v.X() + v.Y()*v.Y()
}

func (v *Vec2) Eq(v2 *Vec2) bool {
	return v.Data == v2.Data
}

func (v *Vec2) Set(x, y float32) {
	v.Data[0] = x
	v.Data[1] = y
}

func (v *Vec2) Normalize() {
	mag := v.Mag()
	v.Data[0] /= mag
	v.Data[1] /= mag
}

func (v *Vec2) Clone() *Vec2 {
	return &Vec2{Data: v.Data}
}

//AddVec2 v3 = v1 + v2
func AddVec2(v1, v2 *Vec2) *Vec2 {
	return &Vec2{
		Data: [2]float32{
			v1.X() + v2.X(),
			v1.Y() + v2.Y(),
		},
	}
}

//SubVec2 v3 = v1 - v2
func SubVec2(v1, v2 *Vec2) *Vec2 {
	return &Vec2{
		Data: [2]float32{
			v1.X() - v2.X(),
			v1.Y() - v2.Y(),
		},
	}
}

func NewVec2(x, y float32) *Vec2 {
	return &Vec2{
		[2]float32{
			x,
			y,
		},
	}
}
