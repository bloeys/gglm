package gglm

//TMat is the 4x4 transformation matrix
// type TMat struct {
// 	Mat4
// }

func GenTranslationMat(v *Vec3) *Mat4 {
	return &Mat4{
		Data: [16]float32{
			1, 0, 0, v.Data[0],
			0, 1, 0, v.Data[1],
			0, 0, 1, v.Data[2],
			0, 0, 0, 1,
		},
	}
}

func GenScaleMat(v *Vec3) *Mat4 {
	return &Mat4{
		Data: [16]float32{
			v.Data[0], 0, 0, 0,
			0, v.Data[1], 0, 0,
			0, 0, v.Data[2], 0,
			0, 0, 0, 1,
		},
	}
}

func GenRotationMat(q *Quat) *Mat4 {

	xy := q.Data[0] * q.Data[1]
	xz := q.Data[0] * q.Data[2]
	xw := q.Data[0] * q.Data[3]

	yz := q.Data[1] * q.Data[2]
	yw := q.Data[1] * q.Data[3]

	zw := q.Data[2] * q.Data[3]

	xx := q.Data[0] * q.Data[0]
	yy := q.Data[1] * q.Data[1]
	zz := q.Data[2] * q.Data[2]
	ww := q.Data[3] * q.Data[3]

	return &Mat4{
		Data: [16]float32{
			xx + yy - zz - ww, 2 * (yz - xw), 2 * (xz + yw), 0,
			2 * (xw + yz), xx - yy + zz - ww, 2 * (zw - xy), 0,
			2 * (yw - xz), 2 * (xy + zw), xx - yy - zz + ww, 0,
			0, 0, 0, 1,
		},
	}
}

// func NewTMatId() *TMat {
// 	return &TMat{
// 		Mat4: Mat4{
// 			Data: [16]float32{
// 				1, 0, 0, 0,
// 				0, 1, 0, 0,
// 				0, 0, 1, 0,
// 				0, 0, 0, 1,
// 			},
// 		},
// 	}
// }
