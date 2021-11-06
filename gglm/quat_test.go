package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestNewQuatEuler(t *testing.T) {

	q := gglm.NewQuatEuler(gglm.NewVec3(180, 180, 180).AsRad())
	ans := &gglm.Quat{Vec4: *gglm.NewVec4(0, 0, 0, 1)}

	if !gglm.EqF32(q.X(), ans.X()) || !gglm.EqF32(q.Y(), ans.Y()) || !gglm.EqF32(q.Z(), ans.Z()) || !gglm.EqF32(q.W(), ans.W()) {
		t.Errorf("Got: %v; Expected: %v", q.String(), ans.String())
	}
}

func TestNewQuatAngleAxis(t *testing.T) {

	q := gglm.NewQuatAngleAxis(180*gglm.Deg2Rad, gglm.NewVec3(0, 1, 0))
	ans := &gglm.Quat{Vec4: *gglm.NewVec4(0, 1, 0, 0)}

	if !gglm.EqF32(q.X(), ans.X()) || !gglm.EqF32(q.Y(), ans.Y()) || !gglm.EqF32(q.Z(), ans.Z()) || !gglm.EqF32(q.W(), ans.W()) {
		t.Errorf("Got: %v; Expected: %v", q.String(), ans.String())
	}
}
