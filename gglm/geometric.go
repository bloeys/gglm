package gglm

func DotVec2(v1, v2 *Vec2) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y()
}

func DotVec3(v1, v2 *Vec3) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

func DotVec4(v1, v2 *Vec4) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z() + v1.W()*v2.W()
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

	diff := Vec2{
		Data: [2]float32{
			v1.X() - v2.X(),
			v1.Y() - v2.Y(),
		},
	}

	return diff.Mag()
}

//DistVec3 returns euclidean distance between v1 and v2
func DistVec3(v1, v2 *Vec3) float32 {

	diff := Vec3{
		Data: [3]float32{
			v1.X() - v2.X(),
			v1.Y() - v2.Y(),
			v1.Z() - v2.Z(),
		},
	}

	return diff.Mag()
}

//DistVec2 returns the squared euclidean distance between v1 and v2 (avoids a sqrt)
func SqrDistVec2(v1, v2 *Vec2) float32 {

	diff := Vec2{
		Data: [2]float32{
			v1.X() - v2.X(),
			v1.Y() - v2.Y(),
		},
	}

	return diff.SqrMag()
}

//DistVec3 returns the squared euclidean distance between v1 and v2 (avoids a sqrt)
func SqrDistVec3(v1, v2 *Vec3) float32 {

	diff := Vec3{
		Data: [3]float32{
			v1.X() - v2.X(),
			v1.Y() - v2.Y(),
			v1.Z() - v2.Z(),
		},
	}

	return diff.SqrMag()
}
