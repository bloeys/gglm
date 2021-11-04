package main

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func BenchmarkVec3Add(b *testing.B) {

	v1 := gglm.NewMat4Id()
	v2 := gglm.NewMat4Id()

	for i := 0; i < b.N; i++ {
		gglm.AddMat4(v1, v2)
	}
}

func BenchmarkVec3Add2(b *testing.B) {

	v1 := gglm.NewMat4Id()
	v2 := gglm.NewMat4Id()

	for i := 0; i < b.N; i++ {
		v1.Add(v2)
	}
}
