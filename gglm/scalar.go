package gglm

import "math"

//EqF32 true if abs(f1-f2) <= F32Epsilon
func EqF32(f1, f2 float32) bool {
	return math.Abs(float64(f1-f2)) <= float64(F32Epsilon)
}

//EqF32Epsilon true if abs(f1-f2) <= eps
func EqF32Epsilon(f1, f2, eps float32) bool {
	return math.Abs(float64(f1-f2)) <= float64(eps)
}

func Sin32(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func Asin32(x float32) float32 {
	return float32(math.Asin(float64(x)))
}

func Cos32(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func Acos32(x float32) float32 {
	return float32(math.Acos(float64(x)))
}

func Tan32(x float32) float32 {
	return float32(math.Tan(float64(x)))
}

func Atan32(x float32) float32 {
	return float32(math.Atan(float64(x)))
}

func Atan232(x, y float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func Sincos32(x float32) (sinx, cosx float32) {
	a, b := math.Sincos(float64(x))
	return float32(a), float32(b)
}

func Abs32(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func Sqrt32(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}
