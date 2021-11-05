package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestMat4GetSet(t *testing.T) {

	m1 := gglm.Mat4{}

	m1.Set(0, 1, -10)
	m1.Set(1, 0, 55)
	m1.Set(2, 2, 99)
	m1.Set(3, 3, 513)

	if m1.Get(0, 1) != -10 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(0, 1), -10)
	}

	if m1.Get(1, 0) != 55 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(1, 0), 55)
	}

	if m1.Get(2, 2) != 99 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(2, 2), 99)
	}

	if m1.Get(3, 3) != 513 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(3, 3), 513)
	}
}

func TestMat4Id(t *testing.T) {

	correctAns := gglm.Mat4{
		Data: [16]float32{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}}

	m1 := gglm.NewMat4Id()
	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestSubMat4(t *testing.T) {

	correctAns := gglm.Mat4{
		Data: [16]float32{
			-16, -16, -16, -16,
			-16, -16, -16, -16,
			-16, -16, -16, -16,
			-16, -16, -16, -16,
		}}

	m1 := &gglm.Mat4{
		Data: [16]float32{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		}}
	m2 := &gglm.Mat4{
		Data: [16]float32{
			17, 18, 19, 20,
			21, 22, 23, 24,
			25, 26, 27, 28,
			29, 30, 31, 32,
		}}

	result := gglm.SubMat4(m1, m2)
	m1.Sub(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestAddMat4(t *testing.T) {

	correctAns := gglm.Mat4{
		Data: [16]float32{
			18, 20, 22, 24,
			26, 28, 30, 32,
			34, 36, 38, 40,
			42, 44, 46, 48,
		}}

	m1 := &gglm.Mat4{
		Data: [16]float32{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		}}
	m2 := &gglm.Mat4{
		Data: [16]float32{
			17, 18, 19, 20,
			21, 22, 23, 24,
			25, 26, 27, 28,
			29, 30, 31, 32,
		}}

	result := gglm.AddMat4(m1, m2)
	m1.Add(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestMulMat4(t *testing.T) {

	correctAns := gglm.Mat4{
		Data: [16]float32{
			250, 260, 270, 280,
			618, 644, 670, 696,
			986, 1028, 1070, 1112,
			1354, 1412, 1470, 1528,
		}}

	m1 := &gglm.Mat4{
		Data: [16]float32{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
		}}
	m2 := &gglm.Mat4{
		Data: [16]float32{
			17, 18, 19, 20,
			21, 22, 23, 24,
			25, 26, 27, 28,
			29, 30, 31, 32,
		}}

	result := gglm.MulMat4(m1, m2)
	m1.Mul(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

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
