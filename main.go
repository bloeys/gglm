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

	println("V1: " + v1.String())
	v1.Normalize()
	println("V1 Normal: " + v1.String())

	//Vec3
	v3 := &gglm.Vec3{Data: [3]float32{1, 2, 3}}
	v4 := &gglm.Vec3{Data: [3]float32{4, 5, 6}}
	println(gglm.DistVec3(v3, v4))
	println(gglm.SqrDistVec3(v4, v3))

	println(v3.Eq(v4))
	v4.Set(1, 2, 3)
	println(v3.Eq(v4))

	println(gglm.DotVec3(v3, v4))
	println(gglm.Cross(v3, v4).String())

	println("V3: " + v3.String())
	v3.Normalize()
	println("V3 Normal: " + v3.String())

	//Vec4
	v5 := &gglm.Vec4{Data: [4]float32{1, 2, 3, 4}}
	v6 := &gglm.Vec4{Data: [4]float32{5, 6, 7, 8}}
	println(gglm.DistVec4(v5, v6))
	println(gglm.SqrDistVec4(v5, v6))

	println(v5.Eq(v6))
	v6.Set(1, 2, 3, 4)
	println(v5.Eq(v6))

	println(gglm.DotVec4(v5, v6))

	println("V6: " + v6.String())
	v6.Normalize()
	println("V6 Normal: " + v6.String())

	//Mat2Vec2
	mat2A := gglm.Mat2{
		Data: [4]float32{
			1, 2,
			3, 4,
		},
	}

	vec2A := gglm.Vec2{Data: [2]float32{1, 2}}
	println(gglm.MulMat2Vec2(&mat2A, &vec2A).String())

	//Mat3Vec3
	mat3A := gglm.Mat3{
		Data: [9]float32{
			1, 2, 3,
			4, 5, 6,
			7, 8, 9,
		},
	}

	vec3A := gglm.Vec3{Data: [3]float32{1, 2, 3}}
	mm3v3 := gglm.MulMat3Vec3(&mat3A, &vec3A)
	println(mm3v3.String())

	//ReflectVec2
	vec2B := &gglm.Vec2{Data: [2]float32{4, 5}}
	normA := &gglm.Vec2{Data: [2]float32{0, 1}}
	rVec2A := gglm.ReflectVec2(vec2B, normA)
	println(rVec2A.String())

	//Quaternion
	vRot := &gglm.Vec3{Data: [3]float32{60, 30, 20}}
	q := gglm.NewQuatEuler(vRot.AsRad())
	println("\n" + vRot.AsRad().String())
	println(q.String(), "\n", q.Mag())

	q = gglm.NewQuatAngleAxis(60*gglm.Deg2Rad, vRot.Normalize())
	println("\n" + vRot.Normalize().String())
	println(q.String())
}
