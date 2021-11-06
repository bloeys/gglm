package gglm

import "math"

func DotVec2(v1, v2 *Vec2) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y()
}

func DotVec3(v1, v2 *Vec3) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

func DotVec4(v1, v2 *Vec4) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z() + v1.W()*v2.W()
}

func DotQuat(q1, q2 *Quat) float32 {
	return q1.X()*q2.X() + q1.Y()*q2.Y() + q1.Z()*q2.Z() + q1.W()*q2.W()
}

func Cross(v1, v2 *Vec3) *Vec3 {
	return &Vec3{
		Data: [3]float32{
			v1.Data[1]*v2.Data[2] - v1.Data[2]*v2.Data[1],
			v1.Data[2]*v2.Data[0] - v1.Data[0]*v2.Data[2],
			v1.Data[0]*v2.Data[1] - v1.Data[1]*v2.Data[0],
		},
	}
}

//DistVec2 returns euclidean distance between v1 and v2
func DistVec2(v1, v2 *Vec2) float32 {
	x := v1.X() - v2.X()
	y := v1.Y() - v2.Y()
	return float32(math.Sqrt(float64(x*x + y*y)))
}

//DistVec3 returns euclidean distance between v1 and v2
func DistVec3(v1, v2 *Vec3) float32 {
	x := v1.X() - v2.X()
	y := v1.Y() - v2.Y()
	z := v1.Z() - v2.Z()
	return float32(math.Sqrt(float64(x*x + y*y + z*z)))
}

//DistVec4 returns euclidean distance between v1 and v2
func DistVec4(v1, v2 *Vec4) float32 {

	//Using X() etc won't let the function inline
	x := v1.Data[0] - v2.Data[0]
	y := v1.Data[1] - v2.Data[1]
	z := v1.Data[2] - v2.Data[2]
	w := v1.Data[3] - v2.Data[3]
	return float32(math.Sqrt(float64(x*x + y*y + z*z + w*w)))
}

//DistVec2 returns the squared euclidean distance between v1 and v2 (avoids a sqrt)
func SqrDistVec2(v1, v2 *Vec2) float32 {
	x := v1.X() - v2.X()
	y := v1.Y() - v2.Y()
	return x*x + y*y
}

//DistVec3 returns the squared euclidean distance between v1 and v2 (avoids a sqrt)
func SqrDistVec3(v1, v2 *Vec3) float32 {
	x := v1.X() - v2.X()
	y := v1.Y() - v2.Y()
	z := v1.Z() - v2.Z()
	return x*x + y*y + z*z
}

//DistVec4 returns the squared euclidean distance between v1 and v2 (avoids a sqrt)
func SqrDistVec4(v1, v2 *Vec4) float32 {
	x := v1.Data[0] - v2.Data[0]
	y := v1.Data[1] - v2.Data[1]
	z := v1.Data[2] - v2.Data[2]
	w := v1.Data[3] - v2.Data[3]
	return x*x + y*y + z*z + w*w
}

//ReflectVec2 returns the reflected vector of the incoming vector 'v', and the surface normal 'n'.
//
//Note: n must be normalized or you will get wrong results
func ReflectVec2(v, n *Vec2) *Vec2 {

	//reflectedVec = v − 2*dot(v, norm)*norm
	d := 2 * (v.Data[0]*n.Data[0] + v.Data[1]*n.Data[1])

	return &Vec2{
		Data: [2]float32{
			v.Data[0] - d*n.Data[0],
			v.Data[1] - d*n.Data[1],
		},
	}
}

//ReflectVec3 returns the reflected vector of the incoming vector 'v', and the surface normal 'n'.
//
//Note: n must be normalized or you will get wrong results
func ReflectVec3(v, n *Vec3) *Vec3 {

	//reflectedVec = v − 2*dot(v, norm)*norm
	d := 2 * (v.Data[0]*n.Data[0] + v.Data[1]*n.Data[1] + v.Data[2]*n.Data[2])

	return &Vec3{
		Data: [3]float32{
			v.Data[0] - d*n.Data[0],
			v.Data[1] - d*n.Data[1],
			v.Data[2] - d*n.Data[2],
		},
	}
}

//AngleQuat returns the angle between the two quaternions in radians
func AngleQuat(q1, q2 *Quat) float32 {
	return Acos32(DotQuat(q1, q2))
}
