package main

import "github.com/bloeys/gglm/gglm"

func main() {

	m1 := &gglm.Mat3{
		Data: [9]float32{
			1, 2, 3,
			4, 5, 6,
			7, 8, 9,
		},
	}

	m2 := &gglm.Mat3{
		Data: [9]float32{
			1, 2, 3,
			1, 2, 3,
			1, 2, 3,
		},
	}

	m3 := gglm.MulMat3(m1, m2)
	m1.Mul(m2)
	println(m1.String())
	println(m3.String())
}
