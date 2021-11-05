package main

import "github.com/bloeys/gglm/gglm"

func main() {

	//Mat3
	m1 := &gglm.Mat3{
		Data: [9]float32{
			1, 2, 3,
			4, 5, 6,
			7, 8, 9,
		},
	}

	m2 := &gglm.Mat3{
		Data: [9]float32{
			1, 2, 3,
			1, 2, 3,
			1, 2, 3,
		},
	}

	m3 := gglm.MulMat3(m1, m2)
	m1.Mul(m2)
	println(m1.String())
	println(m3.String())

	//Mat4
	m4 := &gglm.Mat4{
		Data: [16]float32{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		},
	}

	m5 := &gglm.Mat4{
		Data: [16]float32{
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
			1, 2, 3, 4,
		},
	}

	m6 := gglm.MulMat4(m4, m5)
	m4.Mul(m5)
	println(m4.String())
	println(m6.String())
	println(m4.Eq(m6))

	//Vec2
	v1 := &gglm.Vec2{Data: [2]float32{1, 2}}
	v2 := &gglm.Vec2{Data: [2]float32{3, 4}}
	println(gglm.DistVec2(v1, v2))
	println(gglm.SqrDistVec2(v2, v1))
	println(v1.Eq(v2))
	v2.Set(1, 2)
	println(v1.Eq(v2))

	//Vec3
	v3 := &gglm.Vec3{Data: [3]float32{1, 2, 3}}
	v4 := &gglm.Vec3{Data: [3]float32{4, 5, 6}}
	println(gglm.DistVec3(v3, v4))
	println(gglm.SqrDistVec3(v4, v3))

	println(v3.Eq(v4))
	v4.Set(1, 2, 3)
	println(v3.Eq(v4))

	println(gglm.DotVec3(v3, v4))

	//Vec4
	v5 := &gglm.Vec4{Data: [4]float32{1, 2, 3, 4}}
	v6 := &gglm.Vec4{Data: [4]float32{5, 6, 7, 8}}
	println(gglm.DistVec4(v5, v6))
	println(gglm.SqrDistVec4(v5, v6))

	println(v5.Eq(v6))
	v6.Set(1, 2, 3, 4)
	println(v5.Eq(v6))

	println(gglm.DotVec4(v5, v6))
}
