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

func TestQuatAngle(t *testing.T) {

	a := gglm.NewQuatAngleAxis(180*gglm.Deg2Rad, gglm.NewVec3(0, 1, 0)).Angle()
	var ans float32 = 180.0 * gglm.Deg2Rad

	if !gglm.EqF32(a, ans) {
		t.Errorf("Got: %v; Expected: %v", a, ans)
	}

	a = gglm.NewQuatAngleAxis(90*gglm.Deg2Rad, gglm.NewVec3(1, 1, 0).Normalize()).Angle()
	ans = 90 * gglm.Deg2Rad

	if !gglm.EqF32(a, ans) {
		t.Errorf("Got: %v; Expected: %v", a, ans)
	}

	a = gglm.NewQuatAngleAxis(125*gglm.Deg2Rad, gglm.NewVec3(1, 1, 0).Normalize()).Angle()
	ans = 125 * gglm.Deg2Rad

	if !gglm.EqF32(a, ans) {
		t.Errorf("Got: %v; Expected: %v", a, ans)
	}
}

func TestQuatAxis(t *testing.T) {

	a := gglm.NewQuatAngleAxis(1, gglm.NewVec3(0, 1, 0)).Axis()
	ans := gglm.NewVec3(0, 1, 0)

	if !gglm.EqF32(a.X(), ans.X()) || !gglm.EqF32(a.Y(), ans.Y()) || !gglm.EqF32(a.Z(), ans.Z()) {
		t.Errorf("Got: %v; Expected: %v", a.String(), ans.String())
	}

	a = gglm.NewQuatAngleAxis(1, gglm.NewVec3(1, 1, 0).Normalize()).Axis()
	ans = gglm.NewVec3(1, 1, 0).Normalize()

	if !gglm.EqF32(a.X(), ans.X()) || !gglm.EqF32(a.Y(), ans.Y()) || !gglm.EqF32(a.Z(), ans.Z()) {
		t.Errorf("Got: %v; Expected: %v", a.String(), ans.String())
	}

	a = gglm.NewQuatAngleAxis(1, gglm.NewVec3(67, 46, 32).Normalize()).Axis()
	ans = gglm.NewVec3(67, 46, 32).Normalize()

	if !gglm.EqF32(a.X(), ans.X()) || !gglm.EqF32(a.Y(), ans.Y()) || !gglm.EqF32(a.Z(), ans.Z()) {
		t.Errorf("Got: %v; Expected: %v", a.String(), ans.String())
	}
}
