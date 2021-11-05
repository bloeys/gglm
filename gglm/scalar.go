package gglm

import "math"

//F32Epsilon = 0.0000005
const F32Epsilon float32 = 1e-6

//EqF32 true if abs(f1-f2) <= F32Epsilon
func EqF32(f1, f2 float32) bool {
	return math.Abs(float64(f1-f2)) <= float64(F32Epsilon)
}

//EqF32Epsilon true if abs(f1-f2) <= eps
func EqF32Epsilon(f1, f2, eps float32) bool {
	return math.Abs(float64(f1-f2)) <= float64(eps)
}
