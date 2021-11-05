package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestMulMat4Vec4(t *testing.T) {

	m := &gglm.Mat4{
		Data: [16]float32{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		}}
	v := &gglm.Vec4{Data: [4]float32{1, 2, 3, 4}}

	result := gglm.MulMat4Vec4(m, v)
	correctAns := gglm.Vec4{Data: [4]float32{30, 70, 110, 150}}

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}
}
