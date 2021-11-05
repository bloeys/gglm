package main

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func BenchmarkVec3Add(b *testing.B) {

	v1 := gglm.NewMat3Id()
	v2 := gglm.NewMat3Id()

	for i := 0; i < b.N; i++ {
		v1.Mul(v2)
	}
}

func BenchmarkVec3Add2(b *testing.B) {

	v1 := gglm.NewMat3Id()
	v2 := gglm.NewMat3Id()

	for i := 0; i < b.N; i++ {
		v1.Mul(v2)
	}
}
