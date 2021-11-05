package gglm

type Swizzle1 interface {
	X() float32
	R() float32

	SetX(float32)
	SetR(float32)
}

type Swizzle2 interface {
	Swizzle1
	Y() float32
	G() float32

	SetY(float32)
	SetG(float32)
}

type Swizzle3 interface {
	Swizzle2
	Z() float32
	B() float32

	SetZ(float32)
	SetB(float32)
}

type Swizzle4 interface {
	Swizzle3
	W() float32
	A() float32

	SetW(float32)
	SetA(float32)
}
