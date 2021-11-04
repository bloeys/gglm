package gglm

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas32"
)

var _ Swizzle3 = &Vec3{}
var _ fmt.Stringer = &Vec3{}

type Vec3 struct {
	Arr [3]float32
	*blas32.Vector
}

func (v *Vec3) X() float32 {
	return v.Arr[0]
}

func (v *Vec3) Y() float32 {
	return v.Arr[1]
}

func (v *Vec3) Z() float32 {
	return v.Arr[2]
}

func (v *Vec3) R() float32 {
	return v.Arr[0]
}

func (v *Vec3) G() float32 {
	return v.Arr[1]
}

func (v *Vec3) B() float32 {
	return v.Arr[2]
}

func (v *Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X(), v.Y(), v.Z())
}

//Scale v *= x (element wise multiplication)
func (v *Vec3) Scale(x float32) {
	v.Arr[0] *= x
	v.Arr[1] *= x
	v.Arr[2] *= x
}

func (v *Vec3) Add(v2 *Vec3) {
	v.Arr[0] += v2.X()
	v.Arr[1] += v2.Y()
	v.Arr[2] += v2.Z()
}

//SubVec3 v -= v2
func (v *Vec3) Sub(v2 *Vec3) {
	v.Arr[0] -= v2.X()
	v.Arr[1] -= v2.Y()
	v.Arr[2] -= v2.Z()
}

//Mag returns the magnitude of the vector
func (v *Vec3) Mag() float32 {
	return blas32.Nrm2(*v.Vector)
}

//Mag returns the squared magnitude of the vector
func (v *Vec3) SqrMag() float32 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

//AddVec3 v3 = v1 + v2
func AddVec3(v1, v2 *Vec3) *Vec3 {
	return NewVec3([]float32{
		v1.X() + v2.X(),
		v1.Y() + v2.Y(),
		v1.Z() + v2.Z(),
	})
}

//SubVec3 v3 = v1 - v2
func SubVec3(v1, v2 *Vec3) *Vec3 {
	return NewVec3([]float32{
		v1.X() - v2.X(),
		v1.Y() - v2.Y(),
		v1.Z() - v2.Z(),
	})
}

//NewVec3 copies data if its not nil
func NewVec3(data []float32) *Vec3 {

	arr := [3]float32{}
	if data != nil {

		if len(data) != 3 {
			panic("Data must be nil or size 3")
		}

		arr[0] = data[0]
		arr[1] = data[1]
		arr[2] = data[2]
	}

	return &Vec3{
		Arr: arr,
		Vector: &blas32.Vector{
			N:    3,
			Inc:  1,
			Data: arr[:],
		},
	}
}
