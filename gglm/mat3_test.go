package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestMat3GetSet(t *testing.T) {

	m1 := gglm.Mat3{}

	m1.Set(0, 1, -10)
	m1.Set(1, 0, 55)
	m1.Set(2, 2, 99)

	if m1.Get(0, 1) != -10 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(0, 1), -10)
	}

	if m1.Get(1, 0) != 55 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(1, 0), 55)
	}

	if m1.Get(2, 2) != 99 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(2, 2), 99)
	}
}

func TestMat3Id(t *testing.T) {

	correctAns := gglm.Mat3{
		Data: [3][3]float32{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		}}

	m1 := gglm.NewMat3Id()
	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestSubMat3(t *testing.T) {

	correctAns := gglm.Mat3{
		Data: [3][3]float32{
			{-9, -9, -9},
			{-9, -9, -9},
			{-9, -9, -9},
		}}

	m1 := &gglm.Mat3{
		Data: [3][3]float32{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		}}
	m2 := &gglm.Mat3{
		Data: [3][3]float32{
			{10, 13, 16},
			{11, 14, 17},
			{12, 15, 18},
		}}

	result := gglm.SubMat3(m1, m2)
	m1.Sub(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestAddMat3(t *testing.T) {

	correctAns := gglm.Mat3{
		Data: [3][3]float32{
			{11, 17, 23},
			{13, 19, 25},
			{15, 21, 27},
		}}

	m1 := &gglm.Mat3{
		Data: [3][3]float32{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		}}
	m2 := &gglm.Mat3{
		Data: [3][3]float32{
			{10, 13, 16},
			{11, 14, 17},
			{12, 15, 18},
		}}

	result := gglm.AddMat3(m1, m2)
	m1.Add(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestMulMat3(t *testing.T) {

	correctAns := gglm.Mat3{
		Data: [3][3]float32{
			{84, 201, 318},
			{90, 216, 342},
			{96, 231, 366},
		}}

	m1 := &gglm.Mat3{
		Data: [3][3]float32{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		}}
	m2 := &gglm.Mat3{
		Data: [3][3]float32{
			{10, 13, 16},
			{11, 14, 17},
			{12, 15, 18},
		}}

	result := gglm.MulMat3(m1, m2)
	m1.Mul(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestMulMat3Vec3(t *testing.T) {

	m := &gglm.Mat3{
		Data: [3][3]float32{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		}}
	v := &gglm.Vec3{Data: [3]float32{1, 2, 3}}

	result := gglm.MulMat3Vec3(m, v)
	correctAns := gglm.Vec3{Data: [3]float32{14, 32, 50}}

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}
}
