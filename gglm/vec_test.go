package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestVecSwizzleGet(t *testing.T) {

	//Vec2
	v2 := gglm.NewVec2(1, 2)
	var ans2X float32 = 1
	var ans2Y float32 = 2

	if v2.X() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.X(), ans2X)
	}

	if v2.Y() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.Y(), ans2Y)
	}

	if v2.R() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.R(), ans2X)
	}

	if v2.G() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.G(), ans2Y)
	}

	//Vec3
	v3 := gglm.NewVec3(1, 2, 3)
	var ans3X float32 = 1
	var ans3Y float32 = 2
	var ans3Z float32 = 3

	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	if v3.Z() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.Z(), ans3Z)
	}

	if v3.R() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.R(), ans3X)
	}

	if v3.G() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.G(), ans3Y)
	}

	if v3.B() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.B(), ans3Z)
	}

	//Vec4
	v4 := gglm.NewVec4(1, 2, 3, 4)
	var ans4X float32 = 1
	var ans4Y float32 = 2
	var ans4Z float32 = 3
	var ans4W float32 = 4

	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	if v4.W() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.W(), ans4W)
	}

	if v4.R() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.R(), ans4X)
	}

	if v4.G() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.G(), ans4Y)
	}

	if v4.B() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.B(), ans4Z)
	}

	if v4.A() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.A(), ans4W)
	}
}

func TestVecSwizzleSet(t *testing.T) {

	//Vec2
	v2 := gglm.NewVec2(0, 0)
	var ans2X float32 = 1
	var ans2Y float32 = 2

	v2.SetX(1)
	v2.SetY(2)

	if v2.X() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.X(), ans2X)
	}

	if v2.Y() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.Y(), ans2Y)
	}

	ans2X = 11
	ans2Y = 22
	v2.SetR(11)
	v2.SetG(22)

	if v2.X() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.X(), ans2X)
	}

	if v2.Y() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.Y(), ans2Y)
	}

	//Vec3
	v3 := gglm.NewVec3(0, 0, 0)
	var ans3X float32 = 1
	var ans3Y float32 = 2
	var ans3Z float32 = 3

	v3.SetX(1)
	v3.SetY(2)
	v3.SetZ(3)

	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	if v3.Z() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.Z(), ans3Z)
	}

	ans3X = 11
	ans3Y = 22
	ans3Z = 33

	v3.SetR(11)
	v3.SetG(22)
	v3.SetB(33)

	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	if v3.Z() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.Z(), ans3Z)
	}

	//Vec4
	v4 := gglm.NewVec4(0, 0, 0, 0)
	var ans4X float32 = 1
	var ans4Y float32 = 2
	var ans4Z float32 = 3
	var ans4W float32 = 4

	v4.SetX(1)
	v4.SetY(2)
	v4.SetZ(3)
	v4.SetW(4)

	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	if v4.W() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.W(), ans4W)
	}

	ans4X = 11
	ans4Y = 22
	ans4Z = 33
	ans4W = 44

	v4.SetR(11)
	v4.SetG(22)
	v4.SetB(33)
	v4.SetA(44)

	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	if v4.W() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.W(), ans4W)
	}
}

func TestVecSwizzleAdd(t *testing.T) {

	//Vec2
	v2 := gglm.NewVec2(0, 0)
	var ans2X float32 = 1
	var ans2Y float32 = 2
	v2.AddX(1)
	v2.AddY(2)
	if v2.X() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.X(), ans2X)
	}

	if v2.Y() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.Y(), ans2Y)
	}

	v2 = gglm.NewVec2(0, 0)
	v2.AddR(1)
	v2.AddG(2)
	if v2.X() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.X(), ans2X)
	}

	if v2.Y() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.Y(), ans2Y)
	}

	v2 = gglm.NewVec2(0, 0)
	v2.AddXY(1, 2)
	if v2.X() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.X(), ans2X)
	}

	if v2.Y() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.Y(), ans2Y)
	}

	v2 = gglm.NewVec2(0, 0)
	v2.AddRG(1, 2)
	if v2.X() != ans2X {
		t.Errorf("Got: %v; Expected: %v", v2.X(), ans2X)
	}

	if v2.Y() != ans2Y {
		t.Errorf("Got: %v; Expected: %v", v2.Y(), ans2Y)
	}

	//Vec3
	v3 := gglm.NewVec3(0, 0, 0)
	var ans3X float32 = 1
	var ans3Y float32 = 2
	var ans3Z float32 = 3
	v3.AddX(1)
	v3.AddY(2)
	v3.AddZ(3)
	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	if v3.Z() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.Z(), ans3Z)
	}

	v3 = gglm.NewVec3(0, 0, 0)
	v3.AddR(1)
	v3.AddG(2)
	v3.AddB(3)
	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	if v3.Z() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.Z(), ans3Z)
	}

	v3 = gglm.NewVec3(0, 0, 0)
	v3.AddXY(1, 2)
	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	v3 = gglm.NewVec3(0, 0, 0)
	v3.AddRG(1, 2)
	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	v3 = gglm.NewVec3(0, 0, 0)
	v3.AddXYZ(1, 2, 3)
	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	if v3.Z() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.Z(), ans3Z)
	}

	v3 = gglm.NewVec3(0, 0, 0)
	v3.AddRGB(1, 2, 3)
	if v3.X() != ans3X {
		t.Errorf("Got: %v; Expected: %v", v3.X(), ans3X)
	}

	if v3.Y() != ans3Y {
		t.Errorf("Got: %v; Expected: %v", v3.Y(), ans3Y)
	}

	if v3.Z() != ans3Z {
		t.Errorf("Got: %v; Expected: %v", v3.Z(), ans3Z)
	}

	//Vec4
	v4 := gglm.NewVec4(1, 1, 1, 1)
	var ans4X float32 = 2
	var ans4Y float32 = 3
	var ans4Z float32 = 4
	var ans4W float32 = 5
	v4.AddX(1)
	v4.AddY(2)
	v4.AddZ(3)
	v4.AddW(4)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	if v4.W() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.W(), ans4W)
	}

	v4 = gglm.NewVec4(1, 1, 1, 1)
	v4.AddR(1)
	v4.AddG(2)
	v4.AddB(3)
	v4.AddA(4)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	if v4.W() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.W(), ans4W)
	}

	v4 = gglm.NewVec4(1, 1, 1, 1)
	v4.AddXY(1, 2)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	v4 = gglm.NewVec4(1, 1, 1, 1)
	v4.AddRG(1, 2)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	v4 = gglm.NewVec4(1, 1, 1, 1)
	v4.AddXYZ(1, 2, 3)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	v4 = gglm.NewVec4(1, 1, 1, 1)
	v4.AddRGB(1, 2, 3)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	v4 = gglm.NewVec4(1, 1, 1, 1)
	v4.AddXYZW(1, 2, 3, 4)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	if v4.W() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.W(), ans4W)
	}

	v4 = gglm.NewVec4(1, 1, 1, 1)
	v4.AddRGBA(1, 2, 3, 4)
	if v4.X() != ans4X {
		t.Errorf("Got: %v; Expected: %v", v4.X(), ans4X)
	}

	if v4.Y() != ans4Y {
		t.Errorf("Got: %v; Expected: %v", v4.Y(), ans4Y)
	}

	if v4.Z() != ans4Z {
		t.Errorf("Got: %v; Expected: %v", v4.Z(), ans4Z)
	}

	if v4.W() != ans4W {
		t.Errorf("Got: %v; Expected: %v", v4.W(), ans4W)
	}
}
