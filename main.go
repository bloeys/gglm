package main

import (
	"fmt"

	"github.com/bloeys/gglm/gglm"
)

func main() {

	v := gglm.NewVec2([]float32{4, 5})
	println(v.Mag())
	println(v.SqrMag())
	fmt.Println(gglm.NewMat2ID())
	fmt.Println(v)
}
