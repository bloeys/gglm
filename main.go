package main

import (
	"fmt"

	"github.com/bloeys/gglm/gglm"
)

func main() {

	//Mat3
	m1 := &gglm.Mat3{
		Data: [3][3]float32{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
	}

	m2 := &gglm.Mat3{
		Data: [3][3]float32{
			{1, 1, 1},
			{2, 2, 2},
			{3, 3, 3},
		},
	}

	m3 := gglm.MulMat3(m1, m2)
	m1.Mul(m2)
	println(m1.String())
	println(m3.String())

	//Mat4
	m4 := &gglm.Mat4{
		Data: [4][4]float32{
			{1, 5, 9, 13},
			{2, 6, 10, 14},
			{3, 7, 11, 15},
			{4, 8, 12, 16},
		},
	}

	m5 := &gglm.Mat4{
		Data: [4][4]float32{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
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

	v1.AddXY(v2.X(), v2.Y())
	println(v1.String())

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

	v5.Add(v6)
	v5.AddXYZW(v6.X(), v6.Y(), v6.Z(), v6.W())
	println(v6.String())

	println("V6: " + v6.String())
	v6.Normalize()
	println("V6 Normal: " + v6.String())

	//Mat2Vec2
	mat2A := gglm.Mat2{
		Data: [2][2]float32{
			{1, 3},
			{2, 4},
		},
	}

	vec2A := gglm.Vec2{Data: [2]float32{1, 2}}
	println(gglm.MulMat2Vec2(&mat2A, &vec2A).String())

	//Mat3Vec3
	mat3A := gglm.Mat3{
		Data: [3][3]float32{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
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

	//Transform
	translationMat := gglm.NewTranslationMat(&gglm.Vec3{Data: [3]float32{1, 2, 3}})
	rotMat := gglm.NewRotMat(gglm.NewQuatEuler(gglm.NewVec3(60, 30, 20).AsRad()))
	scaleMat := gglm.NewScaleMat(gglm.NewVec3(1, 1, 1))

	modelMat := gglm.NewTrMatId()
	modelMat.Mul(translationMat.Mul(rotMat.Mul(scaleMat)))

	println("\n\n\n", modelMat.String())

	//Clone Vec2
	v2Orig := gglm.Vec2{Data: [2]float32{1, 2}}
	v2Clone := v2Orig.Clone()
	v2Clone.SetX(99)
	println("\n\n", v2Orig.String(), "; ", v2Clone.String())

	//Clone TrMat
	trMatOrig := gglm.NewTranslationMat(gglm.NewVec3(1, 2, 3))
	trMatClone := trMatOrig.Clone()
	trMatClone.Scale(gglm.NewVec3(2, 2, 2))
	trMatClone.Translate(gglm.NewVec3(9, 0, 0))
	println("\n\n", trMatOrig.String(), "; ", trMatClone.String())

	//Quat geo
	q1 := gglm.NewQuatEuler(gglm.NewVec3(180, 0, 0).AsRad())
	q2 := gglm.NewQuatEuler(gglm.NewVec3(0, 180, 0).AsRad())
	println(gglm.AngleQuat(q1, q2) * gglm.Rad2Deg)

	//LookAt
	camPos := gglm.NewVec3(0, 0, 3)
	worldUp := gglm.NewVec3(0, 1, 0)
	targetPos := gglm.NewVec3(0, 0, 0)
	viewMat := gglm.LookAtRH(camPos, targetPos, worldUp)
	println(viewMat.String())

	//Mat2Col
	mc := gglm.NewMat2Id()
	println("===============================")
	println(mc.String())

	mc.Data = [2][2]float32{
		{1, 3},
		{2, 4},
	}
	println(mc.String())
	fmt.Printf("Arr: %v", mc.Data)

	mc2 := gglm.Mat2{Data: [2][2]float32{
		{1, 3},
		{2, 4},
	}}

	println(mc2.Mul(mc).String())
}
