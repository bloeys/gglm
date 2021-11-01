package gglm

func DotVec2(v1 *Vec2, v2 *Vec2) float32 {
	return v1.X()*v2.X() + v2.Y()*v2.Y()
}
