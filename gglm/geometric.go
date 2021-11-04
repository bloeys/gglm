package gglm

func DotVec2(v1, v2 *Vec2) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y()
}

func DotVec3(v1, v2 *Vec3) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
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
