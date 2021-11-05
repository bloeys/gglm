package main

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

var (
	dotVec2Result  float32 = 0
	dotVec3Result  float32 = 0
	crossResult    *gglm.Vec3
	mulMat4Vec4Res *gglm.Vec4
)

func BenchmarkDotVec2(b *testing.B) {

	v1 := &gglm.Vec2{}
	v2 := &gglm.Vec2{}

	for i := 0; i < b.N; i++ {
		dotVec2Result = gglm.DotVec2(v1, v2)
	}
}

func BenchmarkDotVec3(b *testing.B) {

	v1 := &gglm.Vec3{}
	v2 := &gglm.Vec3{}

	for i := 0; i < b.N; i++ {
		dotVec3Result = gglm.DotVec3(v1, v2)
	}
}

func BenchmarkCross(b *testing.B) {

	v1 := &gglm.Vec3{}
	v2 := &gglm.Vec3{}

	for i := 0; i < b.N; i++ {
		crossResult = gglm.Cross(v1, v2)
	}
}

func BenchmarkMulMat2(b *testing.B) {

	m1 := gglm.NewMat2Id()
	m2 := gglm.NewMat2Id()

	for i := 0; i < b.N; i++ {
		m1.Mul(m2)
	}
}

func BenchmarkMulMat3(b *testing.B) {

	m1 := gglm.NewMat3Id()
	m2 := gglm.NewMat3Id()

	for i := 0; i < b.N; i++ {
		m1.Mul(m2)
	}
}

func BenchmarkMulMat4(b *testing.B) {

	m1 := gglm.NewMat4Id()
	m2 := gglm.NewMat4Id()

	for i := 0; i < b.N; i++ {
		m1.Mul(m2)
	}
}

func BenchmarkMulMat4Vec4(b *testing.B) {

	m1 := gglm.NewMat4Id()
	v1 := gglm.Vec4{}

	for i := 0; i < b.N; i++ {
		mulMat4Vec4Res = gglm.MulMat4Vec4(m1, &v1)
	}
}
