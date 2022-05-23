package gglm

type Swizzle1 interface {
	X() float32
	SetX(float32)
	AddX(float32)

	R() float32
	SetR(float32)
	AddR(float32)
}

type Swizzle2 interface {
	Swizzle1

	Y() float32
	SetY(float32)
	AddY(float32)

	SetXY(float32, float32)
	AddXY(float32, float32)

	G() float32
	SetG(float32)
	AddG(float32)

	SetRG(float32, float32)
	AddRG(float32, float32)
}

type Swizzle3 interface {
	Swizzle2

	Z() float32
	SetZ(float32)
	AddZ(float32)

	SetXYZ(float32, float32, float32)
	AddXYZ(float32, float32, float32)

	B() float32
	SetB(float32)
	AddB(float32)

	SetRGB(float32, float32, float32)
	AddRGB(float32, float32, float32)
}

type Swizzle4 interface {
	Swizzle3

	W() float32
	SetW(float32)
	AddW(float32)

	SetXYZW(float32, float32, float32, float32)
	AddXYZW(float32, float32, float32, float32)

	A() float32
	SetA(float32)
	AddA(float32)

	SetRGBA(float32, float32, float32, float32)
	AddRGBA(float32, float32, float32, float32)
}
