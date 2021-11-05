package gglm_test

import (
	"testing"

	"github.com/bloeys/gglm/gglm"
)

func TestMat2GetSet(t *testing.T) {

	m1 := gglm.Mat2{}

	m1.Set(0, 1, -10)
	m1.Set(1, 0, 55)

	if m1.Get(0, 1) != -10 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(0, 1), -10)
	}

	if m1.Get(1, 0) != 55 {
		t.Errorf("Got: %v; Expected: %v", m1.Get(1, 0), 55)
	}
}

func TestMat2Id(t *testing.T) {

	correctAns := gglm.Mat2{
		Data: [4]float32{
			1, 0,
			0, 1,
		}}

	m1 := gglm.NewMat2Id()
	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestSubMat2(t *testing.T) {

	correctAns := gglm.Mat2{
		Data: [4]float32{
			-4, -4,
			-4, -4,
		}}

	m1 := &gglm.Mat2{
		Data: [4]float32{
			1, 2,
			3, 4,
		}}
	m2 := &gglm.Mat2{Data: [4]float32{
		5, 6,
		7, 8,
	}}

	result := gglm.SubMat2(m1, m2)
	m1.Sub(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestAddMat2(t *testing.T) {

	correctAns := gglm.Mat2{
		Data: [4]float32{
			6, 8,
			10, 12,
		}}

	m1 := &gglm.Mat2{
		Data: [4]float32{
			1, 2,
			3, 4,
		}}
	m2 := &gglm.Mat2{Data: [4]float32{
		5, 6,
		7, 8,
	}}

	result := gglm.AddMat2(m1, m2)
	m1.Add(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestMulMat2(t *testing.T) {

	correctAns := gglm.Mat2{
		Data: [4]float32{
			19, 22,
			43, 50,
		}}

	m1 := &gglm.Mat2{
		Data: [4]float32{
			1, 2,
			3, 4,
		}}
	m2 := &gglm.Mat2{Data: [4]float32{
		5, 6,
		7, 8,
	}}

	result := gglm.MulMat2(m1, m2)
	m1.Mul(m2)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}

	if !m1.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", m1.String(), correctAns.String())
	}
}

func TestMulMat2Vec2(t *testing.T) {

	correctAns := gglm.Vec2{Data: [2]float32{5, 11}}

	m := &gglm.Mat2{
		Data: [4]float32{
			1, 2,
			3, 4,
		}}
	v := &gglm.Vec2{Data: [2]float32{1, 2}}

	result := gglm.MulMat2Vec2(m, v)

	if !result.Eq(&correctAns) {
		t.Errorf("Got: %v; Expected: %v", result.String(), correctAns.String())
	}
}
