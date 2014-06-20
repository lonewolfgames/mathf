package mathf

import (
	"fmt"
	"math"
)

// float32 Array representing a 3x3 Matrix
type Mat3 [9]float32

// returns new Mat3
func NewMat3() *Mat3 {
	this := new(Mat3)

	this[0], this[3], this[6] = 1, 0, 0
	this[1], this[4], this[7] = 0, 1, 0
	this[2], this[5], this[8] = 0, 0, 1

	return this
}

// returns a copy of this
func (this *Mat3) Clone() *Mat3 {

	return new(Mat3).Copy(this)
}

// copies other
func (this *Mat3) Copy(other *Mat3) *Mat3 {

	this[0], this[3], this[6] = other[0], other[3], other[6]
	this[1], this[4], this[7] = other[1], other[4], other[7]
	this[2], this[5], this[8] = other[2], other[5], other[8]

	return this
}

// sets this from values
func (this *Mat3) Set(m11, m12, m13, m21, m22, m23, m31, m32, m33 float32) *Mat3 {

	this[0], this[3], this[6] = m11, m12, m13
	this[1], this[4], this[7] = m21, m22, m23
	this[2], this[5], this[8] = m31, m23, m33

	return this
}

// mutiples this by other
func (this *Mat3) Mul(other *Mat3) *Mat3 {
	a11, a12, a13 := this[0], this[3], this[6]
	a21, a22, a23 := this[1], this[4], this[7]
	a31, a32, a33 := this[2], this[5], this[8]

	b11, b12, b13 := other[0], other[3], other[6]
	b21, b22, b23 := other[1], other[4], other[7]
	b31, b32, b33 := other[2], other[5], other[8]

	this[0] = a11*b11 + a21*b12 + a31*b13
	this[3] = a12*b11 + a22*b12 + a32*b13
	this[6] = a13*b11 + a23*b12 + a33*b13

	this[1] = a11*b21 + a21*b22 + a31*b23
	this[4] = a12*b21 + a22*b22 + a32*b23
	this[7] = a13*b21 + a23*b22 + a33*b23

	this[2] = a11*b31 + a21*b32 + a31*b33
	this[5] = a12*b31 + a22*b32 + a32*b33
	this[8] = a13*b31 + a23*b32 + a33*b33

	return this
}

// mutiples a and b saves in this
func (this *Mat3) MMul(a, b *Mat3) *Mat3 {
	a11, a12, a13 := a[0], a[3], a[6]
	a21, a22, a23 := a[1], a[4], a[7]
	a31, a32, a33 := a[2], a[5], a[8]

	b11, b12, b13 := b[0], b[3], b[6]
	b21, b22, b23 := b[1], b[4], b[7]
	b31, b32, b33 := b[2], b[5], b[8]

	this[0] = a11*b11 + a21*b12 + a31*b13
	this[3] = a12*b11 + a22*b12 + a32*b13
	this[6] = a13*b11 + a23*b12 + a33*b13

	this[1] = a11*b21 + a21*b22 + a31*b23
	this[4] = a12*b21 + a22*b22 + a32*b23
	this[7] = a13*b21 + a23*b22 + a33*b23

	this[2] = a11*b31 + a21*b32 + a31*b33
	this[5] = a12*b31 + a22*b32 + a32*b33
	this[8] = a13*b31 + a23*b32 + a33*b33

	return this
}

// mutiples this by scalar
func (this *Mat3) SMul(s float32) *Mat3 {

	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	this[4] *= s
	this[5] *= s
	this[6] *= s
	this[7] *= s
	this[8] *= s

	return this
}

// divides this by scalar
func (this *Mat3) SDiv(s float32) *Mat3 {
	if s != 0 {
		s = 1 / s
	}

	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	this[4] *= s
	this[5] *= s
	this[6] *= s
	this[7] *= s
	this[8] *= s

	return this
}

// returns inverse of this
func (this *Mat3) Inverse() *Mat3 {

	m11, m12, m13 := this[0], this[3], this[6]
	m21, m22, m23 := this[1], this[4], this[7]
	m31, m32, m33 := this[2], this[5], this[8]

	te0 := m22*m33 - m23*m32
	te3 := m13*m32 - m12*m33
	te6 := m12*m23 - m13*m22

	det := m11*te0 + m21*te3 + m31*te6

	if det == 0 {
		return this.Identity()
	}
	det = 1 / det

	this[0] = te0 * det
	this[3] = te3 * det
	this[6] = te6 * det

	this[1] = (m23*m31 - m21*m33) * det
	this[4] = (m11*m33 - m13*m31) * det
	this[7] = (m13*m21 - m11*m23) * det

	this[2] = (m21*m32 - m22*m31) * det
	this[5] = (m12*m31 - m11*m32) * det
	this[8] = (m11*m22 - m12*m21) * det

	return this
}

// saves inverse of other in this
func (this *Mat3) MInverse(other *Mat3) *Mat3 {

	m11, m12, m13 := other[0], other[3], other[6]
	m21, m22, m23 := other[1], other[4], other[7]
	m31, m32, m33 := other[2], other[5], other[8]

	te0 := m22*m33 - m23*m32
	te3 := m13*m32 - m12*m33
	te6 := m12*m23 - m13*m22

	det := m11*te0 + m21*te3 + m31*te6

	if det == 0 {
		return this.Identity()
	}
	det = 1 / det

	this[0] = te0 * det
	this[3] = te3 * det
	this[6] = te6 * det

	this[1] = (m23*m31 - m21*m33) * det
	this[4] = (m11*m33 - m13*m31) * det
	this[7] = (m13*m21 - m11*m23) * det

	this[2] = (m21*m32 - m22*m31) * det
	this[5] = (m12*m31 - m11*m32) * det
	this[8] = (m11*m22 - m12*m21) * det

	return this
}

// saves inverse of mat4 in this
func (this *Mat3) Mat4Inverse(other *Mat4) *Mat3 {

	m11, m12, m13 := other[0], other[4], other[8]
	m21, m22, m23 := other[1], other[5], other[9]
	m31, m32, m33 := other[2], other[6], other[10]

	this0 := m33*m22 - m32*m23
	this3 := -m33*m12 + m32*m13
	this6 := m23*m12 - m22*m13

	det := m11*this0 + m21*this3 + m31*this6

	if det == 0 {
		return this.Identity()
	}
	det = 1 / det

	this[0] = this0 * det
	this[3] = this3 * det
	this[6] = this6 * det

	this[1] = (-m33*m21 + m31*m23) * det
	this[4] = (m33*m11 - m31*m13) * det
	this[7] = (-m23*m11 + m21*m13) * det

	this[2] = (m32*m21 - m31*m22) * det
	this[5] = (-m32*m11 + m31*m12) * det
	this[8] = (m22*m11 - m21*m12) * det

	return this
}

// returns this as identity
func (this *Mat3) Identity() *Mat3 {

	this[0] = 1
	this[1] = 0
	this[2] = 0
	this[3] = 0
	this[4] = 1
	this[5] = 0
	this[6] = 0
	this[7] = 0
	this[8] = 1

	return this
}

// transposes this across diagonal
func (this *Mat3) Transpose() *Mat3 {

	this[1], this[3] = this[3], this[1]
	this[2], this[6] = this[6], this[2]
	this[5], this[7] = this[7], this[5]

	return this
}

// returns the determinant
func (this *Mat3) Determinant() float32 {

	a, b, c := this[0], this[1], this[2]
	d, e, f := this[3], this[4], this[5]
	g, h, i := this[6], this[7], this[8]

	return a*e*i - a*f*h - b*d*i + b*f*g + c*d*h - c*e*g
}

// returns scale matrix
func (this *Mat3) MakeScale(x, y, z float32) *Mat3 {

	this[0], this[3], this[6] = x, 0, 0
	this[1], this[4], this[7] = 0, y, 0
	this[2], this[5], this[8] = 0, 0, z

	return this
}

// returns rotation matrix around the x axis
func (this *Mat3) MakeRotationX(angle float32) *Mat3 {
	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))

	this[0], this[3], this[6] = 0, 0, 0
	this[1], this[4], this[7] = 0, c, -s
	this[2], this[5], this[8] = 0, s, c

	return this
}

// returns rotation matrix around the y axis
func (this *Mat3) MakeRotationY(angle float32) *Mat3 {
	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))

	this[0], this[3], this[6] = c, 0, s
	this[1], this[4], this[7] = 0, 1, 0
	this[2], this[5], this[8] = -s, 0, c

	return this
}

// returns rotation matrix around the z axis
func (this *Mat3) MakeRotationZ(angle float32) *Mat3 {
	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))

	this[0], this[3], this[6] = c, -s, 0
	this[1], this[4], this[7] = s, c, 0
	this[2], this[5], this[8] = 0, 0, 1

	return this
}

// sets values from Quat
func (this *Mat3) FromQuat(q *Quat) *Mat3 {
	x, y, z, w := q[0], q[1], q[2], q[3]
	x2, y2, z2 := x+x, y+y, z+z
	xx, xy, xz := x*x2, x*y2, x*z2
	yy, yz, zz := y*y2, y*z2, z*z2
	wx, wy, wz := w*x2, w*y2, w*z2

	this[0] = 1 - (yy + zz)
	this[3] = xy - wz
	this[6] = xz + wy

	this[1] = xy + wz
	this[4] = 1 - (xx + zz)
	this[7] = yz - wx

	this[2] = xz - wy
	this[5] = yz + wx
	this[8] = 1 - (xx + yy)

	return this
}

// sets values from Mat2
func (this *Mat3) FromMat2(m *Mat2) *Mat3 {

	this[0] = m[0]
	this[1] = m[1]
	this[2] = 0
	this[3] = m[2]
	this[4] = m[3]
	this[5] = 0
	this[6] = 0
	this[7] = 0
	this[8] = 1

	return this
}

// sets values from Mat4
func (this *Mat3) FromMat4(m *Mat4) *Mat3 {

	this[0] = m[0]
	this[1] = m[1]
	this[2] = m[2]
	this[3] = m[4]
	this[4] = m[5]
	this[5] = m[6]
	this[6] = m[8]
	this[7] = m[9]
	this[8] = m[10]

	return this
}

// checks if this equals other
func (this *Mat3) Equals(other *Mat3) bool {

	if this[0] != other[0] {
		return false
	}
	if this[1] != other[1] {
		return false
	}
	if this[2] != other[2] {
		return false
	}
	if this[3] != other[3] {
		return false
	}
	if this[4] != other[4] {
		return false
	}
	if this[5] != other[5] {
		return false
	}
	if this[6] != other[6] {
		return false
	}
	if this[7] != other[7] {
		return false
	}
	if this[8] != other[8] {
		return false
	}

	return true
}

// returns this as string
func (this *Mat3) String() string {

	return fmt.Sprintf("Mat3{\n %f, %f, %f,\n %f, %f, %f,\n %f, %f, %f\n }", this[0], this[3], this[6], this[1], this[4], this[7], this[2], this[5], this[8])
}
