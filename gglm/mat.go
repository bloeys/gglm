package gglm

type MatSize int

const (
	MatSizeUnknown MatSize = iota
	MatSize2x2
	MatSize3x3
	MatSize4x4
)

type Mat interface {
	At(row, col int) float32
	Set(row, col int, val float32)
	Size() MatSize
}
