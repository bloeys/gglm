package gglm

type Mat interface {
	At(row, col int) float32
	Set(row, col int, val float32)
	Size() int
}
