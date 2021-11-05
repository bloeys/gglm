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

func (v *Vec2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X(), v.Y())
}

//Scale v *= x (element wise multiplication)
func (v *Vec2) Scale(x float32) {
	v.Data[0] *= x
	v.Data[1] *= x
}

//Add v += v2
func (v *Vec2) Add(v2 *Vec2) {
	v.Data[0] += v2.X()
	v.Data[1] += v2.Y()
}

//SubVec2 v -= v2
func (v *Vec2) Sub(v2 *Vec2) {
	v.Data[0] -= v2.X()
	v.Data[1] -= v2.Y()
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
