package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestNewTrMatId(t *testing.T) {

	m := gglm.NewTrMatId()
	ans := &gglm.TrMat{
		Mat4: gglm.Mat4{
			Data: [4][4]float32{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
		},
	}

	if !m.Eq(ans) {
		t.Errorf("Got: %v; Expected: %v", m.String(), ans.String())
	}
}

func TestNewTranslationMat(t *testing.T) {

	m := gglm.NewTranslationMat(gglm.NewVec3(1, 2, 3))
	ans := &gglm.TrMat{
		Mat4: gglm.Mat4{
			Data: [4][4]float32{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{1, 2, 3, 1},
			},
		},
	}

	if !m.Eq(ans) {
		t.Errorf("Got: %v; Expected: %v", m.String(), ans.String())
	}
}

func TestNewScaleMat(t *testing.T) {

	m := gglm.NewScaleMat(gglm.NewVec3(1, 2, 3))
	ans := &gglm.TrMat{
		Mat4: gglm.Mat4{
			Data: [4][4]float32{
				{1, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 3, 0},
				{0, 0, 0, 1},
			},
		},
	}

	if !m.Eq(ans) {
		t.Errorf("Got: %v; Expected: %v", m.String(), ans.String())
	}
}

func TestNewRotMat(t *testing.T) {

	m := gglm.NewRotMat(gglm.NewQuatId())
	ans := &gglm.TrMat{
		Mat4: gglm.Mat4{
			Data: [4][4]float32{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
		},
	}

	if !m.Eq(ans) {
		t.Errorf("Got: %v; Expected: %v", m.String(), ans.String())
	}
}
