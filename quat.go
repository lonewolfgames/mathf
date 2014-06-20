package mathf

import (
	"fmt"
	"math"
)

// float32 Array representing a Quaternion
type Quat [4]float32

// returns new Quat
func NewQuat() *Quat {
	this := new(Quat)

	this[0], this[1], this[2], this[3] = 0, 0, 0, 1

	return this
}

// returns a copy of this
func (this *Quat) Clone() *Quat {

	return new(Quat).Copy(this)
}

// copies other
func (this *Quat) Copy(other *Quat) *Quat {

	this[0] = other[0]
	this[1] = other[1]
	this[2] = other[2]
	this[3] = other[3]

	return this
}

// sets this from values
func (this *Quat) Set(x, y, z, w float32) *Quat {

	this[0] = x
	this[1] = y
	this[2] = z
	this[3] = w

	return this
}

// mutiples this by other
func (this *Quat) Mul(other *Quat) *Quat {
	ax, ay, az, aw := this[0], this[1], this[2], this[3]
	bx, by, bz, bw := other[0], other[1], other[2], other[3]

	this[0] = ax*bw + aw*bx + ay*bz - az*by
	this[1] = ay*bw + aw*by + az*bx - ax*bz
	this[2] = az*bw + aw*bz + ax*by - ay*bx
	this[3] = aw*bw - ax*bx - ay*by - az*bz

	return this
}

// mutiples a and b saves in this
func (this *Quat) QMul(a, b *Quat) *Quat {
	ax, ay, az, aw := a[0], a[1], a[2], a[3]
	bx, by, bz, bw := b[0], b[1], b[2], b[3]

	this[0] = ax*bw + aw*bx + ay*bz - az*by
	this[1] = ay*bw + aw*by + az*bx - ax*bz
	this[2] = az*bw + aw*bz + ax*by - ay*bx
	this[3] = aw*bw - ax*bx - ay*by - az*bz

	return this
}

// mutiples this by scalar
func (this *Quat) SMul(s float32) *Quat {

	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s

	return this
}

// divides this by other
func (this *Quat) Div(other *Quat) *Quat {
	ax, ay, az, aw := this[0], this[1], this[2], this[3]
	bx, by, bz, bw := -other[0], -other[1], -other[2], -other[3]

	this[0] = ax*bw + aw*bx + ay*bz - az*by
	this[1] = ay*bw + aw*by + az*bx - ax*bz
	this[2] = az*bw + aw*bz + ax*by - ay*bx
	this[3] = aw*bw - ax*bx - ay*by - az*bz

	return this
}

// divides a and b saves in this
func (this *Quat) QDiv(a, b *Quat) *Quat {
	ax, ay, az, aw := a[0], a[1], a[2], a[3]
	bx, by, bz, bw := -b[0], -b[1], -b[2], -b[3]

	this[0] = ax*bw + aw*bx + ay*bz - az*by
	this[1] = ay*bw + aw*by + az*bx - ax*bz
	this[2] = az*bw + aw*bz + ax*by - ay*bx
	this[3] = aw*bw - ax*bx - ay*by - az*bz

	return this
}

// divides this by scalar
func (this *Quat) SDiv(s float32) *Quat {
	if s != 0 {
		s = 1 / s
	}

	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s

	return this
}

// normalize vector, makes length of 1
func (this *Quat) Normalize() *Quat {
	l := this[0]*this[0] + this[1]*this[1] + this[2]*this[2] + this[3]*this[3]

	if l == 0 {
		return this
	}

	l = 1 / float32(math.Sqrt(float64(l)))
	this[0] *= l
	this[1] *= l
	this[2] *= l
	this[3] *= l

	return this
}

// returns length of this
func (this *Quat) Length() float32 {
	l := this[0]*this[0] + this[1]*this[1] + this[2]*this[2] + this[3]*this[3]

	if l == 0 {
		return 0
	}

	return float32(math.Sqrt(float64(l)))
}

// returns length squared of this
func (this *Quat) LengthSq() float32 {

	return this[0]*this[0] + this[1]*this[1] + this[2]*this[2] + this[3]*this[3]
}

// returns dot product of this and other
func (this *Quat) Dot(other *Quat) float32 {

	return this[0]*other[0] + this[1]*other[1] + this[2]*other[2] + this[3]*other[3]
}

// returns dot product of a and b
func (this *Quat) QDot(a, b *Quat) float32 {

	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// returns inverse of this
func (this *Quat) Inverse() *Quat {
	d := this[0]*this[0] + this[1]*this[1] + this[2]*this[2] + this[3]*this[3]
	if d == 0 {
		return this.Identity()
	}
	d = 1 / float32(math.Sqrt(float64(d)))

	this[0] *= -d
	this[1] *= -d
	this[2] *= -d
	this[3] *= d

	return this
}

// saves inverse of other in this
func (this *Quat) QInverse(other *Quat) *Quat {
	x, y, z, w := other[0], other[1], other[2], other[3]
	d := x*x + y*y + z*z + w*w
	if d == 0 {
		return this.Identity()
	}
	d = 1 / float32(math.Sqrt(float64(d)))

	this[0] = x * -d
	this[1] = y * -d
	this[2] = z * -d
	this[3] = w * d

	return this
}

// if this is normalized this is faster than Inverse
func (this *Quat) Conjugate() *Quat {

	this[0] *= -1
	this[1] *= -1
	this[2] *= -1

	return this
}

// returns this as identity
func (this *Quat) Identity() *Quat {

	this[0] = 0
	this[1] = 0
	this[2] = 0
	this[3] = 1

	return this
}

// lerps this and other by x
func (this *Quat) Lerp(other *Quat, x float32) *Quat {

	this[0] += (other[0] - this[0]) * x
	this[1] += (other[1] - this[1]) * x
	this[2] += (other[2] - this[2]) * x
	this[3] += (other[3] - this[3]) * x

	return this
}

// lerps a and b by x, saves in this
func (this *Quat) QLerp(a, b *Quat, x float32) *Quat {

	this[0] = a[0] + (b[0]-a[0])*x
	this[1] = a[1] + (b[1]-a[1])*x
	this[2] = a[2] + (b[2]-a[2])*x
	this[3] = a[3] + (b[3]-a[3])*x

	return this
}

// lerps this and other by x and normalizes
func (this *Quat) Nlerp(other *Quat, x float32) *Quat {

	this[0] += (other[0] - this[0]) * x
	this[1] += (other[1] - this[1]) * x
	this[2] += (other[2] - this[2]) * x
	this[3] += (other[3] - this[3]) * x

	return this.Normalize()
}

// lerps a and b by x, saves in this
func (this *Quat) QNlerp(a, b *Quat, x float32) *Quat {

	this[0] = a[0] + (b[0]-a[0])*x
	this[1] = a[1] + (b[1]-a[1])*x
	this[2] = a[2] + (b[2]-a[2])*x
	this[3] = a[3] + (b[3]-a[3])*x

	return this.Normalize()
}

func (this *Quat) Slerp(other *Quat, x float32) *Quat {
	ax, ay, az, aw := this[0], this[1], this[2], this[3]
	bx, by, bz, bw := other[0], other[1], other[2], other[3]

	cosom := ax*bx + ay*by + az*bz + aw*bw
	var omega, sinom, scale0, scale1 float32

	if cosom < 0 {
		cosom *= -1
		bx *= -1
		by *= -1
		bz *= -1
		bw *= -1
	}

	if 1-cosom > 0 {
		omega = float32(math.Acos(float64(cosom)))
		sinom = 1 / float32(math.Sin(float64(omega)))
		scale0 = float32(math.Sin(float64((1-x)*omega))) * sinom
		scale1 = float32(math.Sin(float64(x*omega))) * sinom
	} else {
		scale0 = 1 - x
		scale1 = x
	}

	this[0] = scale0*ax + scale1*bx
	this[1] = scale0*ay + scale1*by
	this[2] = scale0*az + scale1*bz
	this[3] = scale0*aw + scale1*bw

	return this
}

func (this *Quat) QSlerp(a, b *Quat, x float32) *Quat {
	ax, ay, az, aw := a[0], a[1], a[2], a[3]
	bx, by, bz, bw := b[0], b[1], b[2], b[3]

	cosom := ax*bx + ay*by + az*bz + aw*bw
	var omega, sinom, scale0, scale1 float32

	if cosom < 0 {
		cosom *= -1
		bx *= -1
		by *= -1
		bz *= -1
		bw *= -1
	}

	if 1-cosom > 0 {
		omega = float32(math.Acos(float64(cosom)))
		sinom = 1 / float32(math.Sin(float64(omega)))
		scale0 = float32(math.Sin(float64((1-x)*omega))) * sinom
		scale1 = float32(math.Sin(float64(x*omega))) * sinom
	} else {
		scale0 = 1 - x
		scale1 = x
	}

	this[0] = scale0*ax + scale1*bx
	this[1] = scale0*ay + scale1*by
	this[2] = scale0*az + scale1*bz
	this[3] = scale0*aw + scale1*bw

	return this
}

func (this *Quat) RotateX(angle float32) *Quat {
	halfAngle := angle * 0.5
	x, y, z, w := this[0], this[1], this[2], this[3]
	s, c := float32(math.Sin(float64(halfAngle))), float32(math.Cos(float64(halfAngle)))

	this[0] = x*c + w*s
	this[1] = y*c + z*s
	this[2] = z*c - y*s
	this[3] = w*c - x*s

	return this
}

func (this *Quat) RotateY(angle float32) *Quat {
	halfAngle := angle * 0.5
	x, y, z, w := this[0], this[1], this[2], this[3]
	s, c := float32(math.Sin(float64(halfAngle))), float32(math.Cos(float64(halfAngle)))

	this[0] = x*c - z*s
	this[1] = y*c + w*s
	this[2] = z*c + x*s
	this[3] = w*c - y*s

	return this
}

func (this *Quat) RotateZ(angle float32) *Quat {
	halfAngle := angle * 0.5
	x, y, z, w := this[0], this[1], this[2], this[3]
	s, c := float32(math.Sin(float64(halfAngle))), float32(math.Cos(float64(halfAngle)))

	this[0] = x*c + y*s
	this[1] = y*c - x*s
	this[2] = z*c + w*s
	this[3] = w*c - z*s

	return this
}

func (this *Quat) Rotate(x, y, z float32) *Quat {

	this.RotateZ(z)
	this.RotateX(x)
	this.RotateY(y)

	return this
}

// sets values from axis and angle
func (this *Quat) FromAxisAngle(axis *Vec3, angle float32) *Quat {
	halfAngle := angle * 0.5
	s := float32(math.Sin(float64(halfAngle)))

	this[0] = axis[0] * s
	this[1] = axis[1] * s
	this[2] = axis[2] * s
	this[3] = float32(math.Cos(float64(halfAngle)))

	return this
}

// sets values from Mat3
func (this *Quat) FromMat3(m *Mat3) *Quat {
	m11, m12, m13 := m[0], m[3], m[6]
	m21, m22, m23 := m[1], m[4], m[7]
	m31, m32, m33 := m[2], m[5], m[8]
	trace := m11 + m22 + m33
	var s, invS float32

	if trace > 0 {
		s = 0.5 / float32(math.Sqrt(float64(trace+1)))

		this[3] = 0.25 / s
		this[0] = (m32 - m23) * s
		this[1] = (m13 - m31) * s
		this[2] = (m21 - m12) * s

	} else if m11 > m22 && m11 > m33 {
		s = 2 * float32(math.Sqrt(float64(1+m11-m22-m33)))
		invS = 1 / s

		this[3] = (m32 - m23) * invS
		this[0] = 0.25 * s
		this[1] = (m12 + m21) * invS
		this[2] = (m13 + m31) * invS

	} else if m22 > m33 {
		s = 2 * float32(math.Sqrt(float64(1+m22-m11-m33)))
		invS = 1 / s

		this[3] = (m13 - m31) * invS
		this[0] = (m12 + m21) * invS
		this[1] = 0.25 * s
		this[2] = (m23 + m32) * invS

	} else {
		s = 2 * float32(math.Sqrt(float64(1+m33-m11-m22)))
		invS = 1 / s

		this[3] = (m21 - m12) * invS
		this[0] = (m13 + m31) * invS
		this[1] = (m23 + m32) * invS
		this[2] = 0.25 * s
	}

	return this
}

// sets values from Mat4
func (this *Quat) FromMat4(m *Mat4) *Quat {
	m11, m12, m13 := m[0], m[4], m[8]
	m21, m22, m23 := m[1], m[5], m[9]
	m31, m32, m33 := m[2], m[6], m[10]
	trace := m11 + m22 + m33
	var s, invS float32

	if trace > 0 {
		s = 0.5 / float32(math.Sqrt(float64(trace+1)))

		this[3] = 0.25 / s
		this[0] = (m32 - m23) * s
		this[1] = (m13 - m31) * s
		this[2] = (m21 - m12) * s

	} else if m11 > m22 && m11 > m33 {
		s = 2 * float32(math.Sqrt(float64(1+m11-m22-m33)))
		invS = 1 / s

		this[3] = (m32 - m23) * invS
		this[0] = 0.25 * s
		this[1] = (m12 + m21) * invS
		this[2] = (m13 + m31) * invS

	} else if m22 > m33 {
		s = 2 * float32(math.Sqrt(float64(1+m22-m11-m33)))
		invS = 1 / s

		this[3] = (m13 - m31) * invS
		this[0] = (m12 + m21) * invS
		this[1] = 0.25 * s
		this[2] = (m23 + m32) * invS

	} else {
		s = 2 * float32(math.Sqrt(float64(1+m33-m11-m22)))
		invS = 1 / s

		this[3] = (m21 - m12) * invS
		this[0] = (m13 + m31) * invS
		this[1] = (m23 + m32) * invS
		this[2] = 0.25 * s
	}

	return this
}

// checks if this equals other
func (this *Quat) Equals(other *Quat) bool {

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

	return true
}

// returns this as string type
func (this *Quat) String() string {

	return fmt.Sprintf("Quat[ %f, %f, %f, %f ]", this[0], this[1], this[2], this[3])
}
