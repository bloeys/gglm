package gglm

type Swizzle1 interface {
	X() float32
	R() float32
}

type Swizzle2 interface {
	Swizzle1
	Y() float32
	G() float32
}

type Swizzle3 interface {
	Swizzle2
	Z() float32
	B() float32
}
