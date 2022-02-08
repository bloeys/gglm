# gglm

Fast OpenGL/Graphics focused Mathematics library in  Go inspired by the c++ library [glm](https://github.com/g-truc/glm).

gglm currently has the following:

- `Vec2`, `Vec3` and `Vec4` structs that implement vector (x,y,z,w) operations
- `Mat2`, `Mat3`, `Mat4` structs that implement square matrix operations
- `Quat` struct that implements quaternion operations
- `TrMat` struct that implements 3D transformation matrix operations
- Many useful geometric functions (e.g. dot product, cross product, vector reflection etc)
- 32-bit scalar operations (e.g. sin32, cos32, equality using epsilon, sqrt32 etc)
- Useful 32-bit constants (e.g. pi, Deg2Rad, Rad2Deg, float32 epsilon etc)
- Simple 'siwzzle' interfaces that allow you to do things like `.X()` or `.R()` etc.
- Very easy to use with graphics/native APIs as everything is implemented using arrays
- `.String()` functions on all types for pretty pritning

## Installation

`go get github.com/bloeys/gglm`

## Usage

```go

import "github.com/bloeys/gglm/gglm"


func main() {

	//LookAt
	camPos := gglm.NewVec3(0, 0, 3)
	worldUp := gglm.NewVec3(0, 1, 0)
	targetPos := gglm.NewVec3(0, 0, 0)
	viewMat := gglm.LookAt(camPos, targetPos, worldUp)
	println(viewMat.String())
	
	//Vec2
	v1 := &gglm.Vec2{Data: [2]float32{1, 2}}
	v2 := &gglm.Vec2{Data: [2]float32{3, 4}}
	println(gglm.DistVec2(v1, v2))
	println(gglm.SqrDistVec2(v2, v1))
	println(v1.Eq(v2))
	v2.Set(1, 2)
	println(v1.Eq(v2))
}
```

## Notes
You can check compiler inlining decisions using `go run -gcflags "-m" .`. Some functions look a bit weird compared to similar ones
because we are trying to reduce function complexity so the compiler inlines.

