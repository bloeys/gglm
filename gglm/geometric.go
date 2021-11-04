package gglm

import "gonum.org/v1/gonum/blas/blas32"

func DotVec2(v1, v2 *Vec2) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y()
}

func DotVec3(v1, v2 *Vec3) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

func Cross(v1, v2 *Vec3) *Vec3 {

	// x = a[1]*b[2] - a[2]*b[1]
	// y = a[2]*b[0] - a[0]*b[2]
	// z = a[0]*b[1] - a[1]*b[0]

	//Note: It's done this way to get inlining. Any small changes
	//break inlining so check before any changes.
	f := [3]float32{
		v1.Arr[1]*v2.Arr[2] - v1.Arr[2]*v2.Arr[1],
		v1.Arr[2]*v2.Arr[0] - v1.Arr[0]*v2.Arr[2],
		v1.Arr[0]*v2.Arr[1] - v1.Arr[1]*v2.Arr[0],
	}

	return &Vec3{
		Arr: f,
		Vector: &blas32.Vector{
			N:    3,
			Inc:  1,
			Data: f[:],
		},
	}
}
