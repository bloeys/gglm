package gglm

import (
	"gonum.org/v1/gonum/blas/blas32"
)

var _ Swizzle2 = &Vec2{}

type Vec2 struct {
	*blas32.Vector
}

func (v *Vec2) AddVec2(v2 *Vec2) {
	v.Data[0] += v2.Data[0]
	v.Data[1] += v2.Data[1]
}

func (v *Vec2) Scale(x float32) {
	v.Data[0] *= x
	v.Data[1] *= x
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

//Mag returns the magnitude of the vector
func (v *Vec2) Mag() float32 {
	return blas32.Nrm2(*v.Vector)
}

//Mag returns the squared magnitude of the vector
func (v *Vec2) SqrMag() float32 {
	return v.X()*v.X() + v.Y()*v.Y()
}

func NewVec2(data []float32) *Vec2 {

	if data == nil {
		data = make([]float32, 2)
	} else if len(data) != 2 {
		panic("Data must be nil or size 2")
	}

	return &Vec2{
		Vector: &blas32.Vector{
			N:    2,
			Inc:  1,
			Data: data,
		},
	}
}
