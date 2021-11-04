package main

import (
	"github.com/bloeys/gglm/gglm"
)

func main() {

	// v1 := gglm.NewVec2([]float32{4, 5})
	// v2 := gglm.NewVec2([]float32{1, 1})
	// println(v1.Mag())
	// println(v1.SqrMag())
	// v1 := gglm.NewVec2(nil)
	// v2 := gglm.NewVec2(nil)
	// v1.Add(v2)

	m1 := gglm.NewMat4(nil)

	// f := []float32{
	// 	0, 0, 0, 0,
	// 	0, 0, 0, 0,
	// 	0, 0, 0, 0,
	// 	0, 0, 0, 0,
	// }
	m2 := gglm.NewMat4(nil)

	println("m1:", m1.String())
	println("m2:", m2.String())

	m3 := gglm.AddMat4(m1, m2)
	println("m1:", m1.String())
	println("m2:", m2.String())
	println("m3:", m3.String())
}
