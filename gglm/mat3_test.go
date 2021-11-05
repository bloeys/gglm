package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestMulMat3Vec3(t *testing.T) {

	m := &gglm.Mat3{
		Data: [9]float32{
			1, 2, 3,
			4, 5, 6,
			7, 8, 9,
		}}
	v := &gglm.Vec3{Data: [3]float32{1, 2, 3}}

	result := gglm.MulMat3Vec3(m, v)
	correctAns := gglm.Vec3{Data: [3]float32{14, 32, 50}}

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}
}
