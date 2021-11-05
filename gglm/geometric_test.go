package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestDotVec2(t *testing.T) {

	v1 := gglm.Vec2{Data: [2]float32{1, 2}}
	v2 := gglm.Vec2{Data: [2]float32{3, 4}}
	ans := float32(11)

	res := gglm.DotVec2(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestDotVec3(t *testing.T) {

	v1 := gglm.Vec3{Data: [3]float32{1, 2, 3}}
	v2 := gglm.Vec3{Data: [3]float32{4, 5, 6}}
	ans := float32(32)

	res := gglm.DotVec3(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestDotVec4(t *testing.T) {

	v1 := gglm.Vec4{Data: [4]float32{1, 2, 3, 4}}
	v2 := gglm.Vec4{Data: [4]float32{5, 6, 7, 8}}
	ans := float32(70)

	res := gglm.DotVec4(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestDistVec2(t *testing.T) {

	v1 := gglm.Vec2{Data: [2]float32{1, 2}}
	v2 := gglm.Vec2{Data: [2]float32{3, 4}}
	ans := float32(2.828427)

	res := gglm.DistVec2(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestDistVec3(t *testing.T) {

	v1 := gglm.Vec3{Data: [3]float32{1, 2, 3}}
	v2 := gglm.Vec3{Data: [3]float32{4, 5, 6}}
	ans := float32(5.196152)

	res := gglm.DistVec3(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestDistVec4(t *testing.T) {

	v1 := gglm.Vec4{Data: [4]float32{1, 2, 3, 4}}
	v2 := gglm.Vec4{Data: [4]float32{5, 6, 7, 8}}
	ans := float32(8)

	res := gglm.DistVec4(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestSqrDistVec2(t *testing.T) {

	v1 := gglm.Vec2{Data: [2]float32{1, 2}}
	v2 := gglm.Vec2{Data: [2]float32{3, 4}}
	ans := float32(8)

	res := gglm.SqrDistVec2(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestSqrDistVec3(t *testing.T) {

	v1 := gglm.Vec3{Data: [3]float32{1, 2, 3}}
	v2 := gglm.Vec3{Data: [3]float32{4, 5, 6}}
	ans := float32(27)

	res := gglm.SqrDistVec3(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestSqrDistVec4(t *testing.T) {

	v1 := gglm.Vec4{Data: [4]float32{1, 2, 3, 4}}
	v2 := gglm.Vec4{Data: [4]float32{5, 6, 7, 8}}
	ans := float32(64)

	res := gglm.SqrDistVec4(&v1, &v2)
	if res != ans {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestReflectVec2(t *testing.T) {

	v1 := gglm.Vec2{Data: [2]float32{1, 2}}
	n := gglm.Vec2{Data: [2]float32{0, 1}}
	ans := gglm.Vec2{Data: [2]float32{1, -2}}

	res := gglm.ReflectVec2(&v1, &n)
	if !res.Eq(&ans) {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}

func TestReflectVec3(t *testing.T) {

	v1 := gglm.Vec3{Data: [3]float32{1, 2, 3}}
	n := gglm.Vec3{Data: [3]float32{0, 1, 0}}
	ans := gglm.Vec3{Data: [3]float32{1, -2, 3}}

	res := gglm.ReflectVec3(&v1, &n)
	if !res.Eq(&ans) {
		t.Errorf("Got: %v; Expected: %v", res, ans)
	}
}
