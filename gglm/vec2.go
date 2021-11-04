package gglm

//Note: We don't use the Swizzle interface for add/sub because the interface doesn't allow inling :(

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas32"
)

var _ Swizzle2 = &Vec2{}
var _ fmt.Stringer = &Vec2{}

type Vec2 struct {
	Arr [2]float32
	*blas32.Vector
}

func (v *Vec2) X() float32 {
	return v.Arr[0]
}

func (v *Vec2) Y() float32 {
	return v.Arr[1]
}

func (v *Vec2) R() float32 {
	return v.Arr[0]
}

func (v *Vec2) G() float32 {
	return v.Arr[1]
}

func (v *Vec2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X(), v.Y())
}

//Scale v *= x (element wise multiplication)
func (v *Vec2) Scale(x float32) {
	v.Arr[0] *= x
	v.Arr[1] *= x
}

//Add v += v2
func (v *Vec2) Add(v2 *Vec2) {
	v.Arr[0] += v2.X()
	v.Arr[1] += v2.Y()
}

//SubVec2 v -= v2
func (v *Vec2) Sub(v2 *Vec2) {
	v.Arr[0] -= v2.X()
	v.Arr[1] -= v2.Y()
}

//Mag returns the magnitude of the vector
func (v *Vec2) Mag() float32 {
	return blas32.Nrm2(*v.Vector)
}

//Mag returns the squared magnitude of the vector
func (v *Vec2) SqrMag() float32 {
	return v.X()*v.X() + v.Y()*v.Y()
}

//AddVec2 v3 = v1 + v2
func AddVec2(v1, v2 *Vec2) *Vec2 {
	return NewVec2([]float32{
		v1.X() + v2.X(),
		v1.Y() + v2.Y(),
	})
}

//SubVec2 v3 = v1 - v2
func SubVec2(v1, v2 *Vec2) *Vec2 {
	return NewVec2([]float32{
		v1.X() - v2.X(),
		v1.Y() - v2.Y(),
	})
}

//NewVec2 copies data if its not nil
func NewVec2(data []float32) *Vec2 {

	arr := [2]float32{}
	if data != nil {

		if len(data) != 2 {
			panic("Data must be nil or size 2")
		}

		arr[0] = data[0]
		arr[1] = data[1]
	}

	return &Vec2{
		Arr: arr,
		Vector: &blas32.Vector{
			N:    2,
			Inc:  1,
			Data: arr[:],
		},
	}
}
