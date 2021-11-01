package main

import (
	"fmt"

	"github.com/bloeys/gglm/gglm"
)

func main() {

	v1 := gglm.NewVec2([]float32{4, 5})
	v2 := gglm.NewVec2([]float32{1, 1})
	println(v1.Mag())
	println(v1.SqrMag())

	v1.Add(v2)
	v1.Sub(v2)

	fmt.Println(gglm.DotVec2(v1, v2))
	fmt.Println(v1, v1.X(), v1.Y())

	m := gglm.NewMat2(nil)
	fmt.Println(m.String(), m.At(0, 1), m.At(1, 1))
	m.Set(0, 1, 55)
	m.Set(1, 0, 77)
	fmt.Println("\n"+m.String(), m.At(0, 1), m.At(1, 1))
}
