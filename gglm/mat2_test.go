package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestMulMat2Vec2(t *testing.T) {

	m := &gglm.Mat2{
		Data: [4]float32{
			1, 2,
			3, 4,
		}}
	v := &gglm.Vec2{Data: [2]float32{1, 2}}

	result := gglm.MulMat2Vec2(m, v)
	correctAns := gglm.Vec2{Data: [2]float32{5, 11}}

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}
}
