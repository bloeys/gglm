package gglm

import "fmt"

var _ fmt.Stringer = MatSize2x2

type MatSize int

const (
	MatSizeUnknown MatSize = iota
	MatSize2x2
	MatSize3x3
	MatSize4x4
)

//String panics if the MatSize is not known
func (ms MatSize) String() string {

	switch ms {
	case MatSize2x2:
		return "2x2"
	case MatSize3x3:
		return "3x3"
	case MatSize4x4:
		return "4x4"
	default:
		panic("Unkown MatSize '" + fmt.Sprintf("%d", int(ms)) + "'")
	}
}

type Mat interface {
	Get(row, col int) float32
	Set(row, col int, val float32)
	Size() MatSize
}
