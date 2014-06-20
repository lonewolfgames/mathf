package mathf

import (
	"fmt"
	"math"
)

// float32 Array representing a Color RGBA
type Color [4]float32

// returns new Color
func NewColor(r, g, b, a float32) *Color {
	this := new(Color)

	this[0], this[1], this[2], this[3] = r, g, b, a

	return this
}

// returns a copy of this
func (this *Color) Clone() *Color {

	return new(Color).Copy(this)
}

// copies other
func (this *Color) Copy(other *Color) *Color {

	this[0], this[1], this[2], this[3] = other[0], other[1], other[2], other[3]

	return this
}

// sets this from values
func (this *Color) Set(r, g, b, a float32) *Color {

	this[0], this[1], this[2], this[3] = r, g, b, a

	return this
}

// adds other to this
func (this *Color) Add(other *Color) *Color {

	this[0] += other[0]
	this[1] += other[1]
	this[2] += other[2]
	this[3] += other[3]

	return this
}

// adds a and b saves in this
func (this *Color) CAdd(a, b *Color) *Color {

	this[0] = a[0] + b[0]
	this[1] = a[1] + b[1]
	this[2] = a[2] + b[2]
	this[3] = a[3] + b[3]

	return this
}

// adds scalar to this
func (this *Color) SAdd(s float32) *Color {

	this[0] += s
	this[1] += s
	this[2] += s
	this[3] += s

	return this
}

// subtracts other from this
func (this *Color) Sub(other *Color) *Color {

	this[0] += other[0]
	this[1] += other[1]
	this[2] += other[2]

	return this
}

// subtracts a and b saves in this
func (this *Color) CSub(a, b *Color) *Color {

	this[0] = a[0] - b[0]
	this[1] = a[1] - b[1]
	this[2] = a[2] - b[2]

	return this
}

// subtracts scalar from this
func (this *Color) SSub(s float32) *Color {

	this[0] -= s
	this[1] -= s
	this[2] -= s
	this[3] -= s

	return this
}

// mutiples this by other
func (this *Color) Mul(other *Color) *Color {

	this[0] *= other[0]
	this[1] *= other[1]
	this[2] *= other[2]
	this[3] *= other[3]

	return this
}

// mutiples a and b saves in this
func (this *Color) CMul(a, b *Color) *Color {

	this[0] = a[0] * b[0]
	this[1] = a[1] * b[1]
	this[2] = a[2] * b[2]
	this[3] = a[3] * b[3]

	return this
}

// mutiples this by scalar
func (this *Color) SMul(s float32) *Color {

	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s

	return this
}

// divides this by other
func (this *Color) Div(other *Color) *Color {
	r, g, b, a := other[0], other[1], other[2], other[3]

	if r != 0 {
		r = 1 / r
	}
	if g != 0 {
		g = 1 / g
	}
	if b != 0 {
		b = 1 / b
	}
	if a != 0 {
		a = 1 / a
	}

	this[0] *= r
	this[1] *= g
	this[2] *= b
	this[3] *= a

	return this
}

// divides a and b saves in this
func (this *Color) CDiv(a, b *Color) *Color {
	br, bg, bb, ba := b[0], b[1], b[2], b[3]

	if br != 0 {
		br = 1 / br
	}
	if bg != 0 {
		bg = 1 / bg
	}
	if bb != 0 {
		bb = 1 / bb
	}
	if ba != 0 {
		ba = 1 / ba
	}

	this[0] = a[0] * br
	this[1] = a[1] * bg
	this[2] = a[2] * bb
	this[3] = a[3] * ba

	return this
}

// divides this by scalar
func (this *Color) SDiv(s float32) *Color {
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
func (this *Color) Normalize() *Color {
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
func (this *Color) Length() float32 {
	l := this[0]*this[0] + this[1]*this[1] + this[2]*this[2] + this[3]*this[3]

	if l == 0 {
		return 0
	}

	return float32(math.Sqrt(float64(l)))
}

// returns length squared of this
func (this *Color) LengthSq() float32 {

	return this[0]*this[0] + this[1]*this[1] + this[2]*this[2] + this[3]*this[3]
}

// sets length of vector, if length is zero returns zero vector
func (this *Color) SetLength(length float32) *Color {
	l := this[0]*this[0] + this[1]*this[1] + this[2]*this[2] + this[3]*this[3]

	if l == 0 {
		return this
	}

	l = 1 / float32(math.Sqrt(float64(l)))
	this[0] *= length * l
	this[1] *= length * l
	this[2] *= length * l
	this[3] *= length * l

	return this
}

// returns dot product of this and other
func (this *Color) Dot(other *Color) float32 {

	return this[0]*other[0] + this[1]*other[1] + this[2]*other[2] + this[3]*other[3]
}

// returns dot product of a and b
func (this *Color) CDot(a, b *Color) float32 {

	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// lerps this and other by x
func (this *Color) Lerp(other *Color, x float32) *Color {

	this[0] += (other[0] - this[0]) * x
	this[1] += (other[1] - this[1]) * x
	this[2] += (other[2] - this[2]) * x
	this[3] += (other[3] - this[3]) * x

	return this
}

// lerps a and b by x, saves in this
func (this *Color) CLerp(a, b *Color, x float32) *Color {

	this[0] = a[0] + (b[0]-a[0])*x
	this[1] = a[1] + (b[1]-a[1])*x
	this[2] = a[2] + (b[2]-a[2])*x
	this[3] = a[3] + (b[3]-a[3])*x

	return this
}

// sets this values from  min values of this and other
func (this *Color) Min(other *Color) *Color {

	if this[0] > other[0] {
		this[0] = other[0]
	}
	if this[1] > other[1] {
		this[1] = other[1]
	}
	if this[2] > other[2] {
		this[2] = other[2]
	}
	if this[3] > other[3] {
		this[3] = other[3]
	}

	return this
}

// sets this values from max values of this and other
func (this *Color) Max(other *Color) *Color {

	if this[0] < other[0] {
		this[0] = other[0]
	}
	if this[1] < other[1] {
		this[1] = other[1]
	}
	if this[2] < other[2] {
		this[2] = other[2]
	}
	if this[3] < other[3] {
		this[3] = other[3]
	}

	return this
}

// clamps this between min and max
func (this *Color) Clamp(min, max *Color) *Color {

	if this[0] < min[0] {
		this[0] = min[0]
	}
	if this[1] < min[1] {
		this[1] = min[1]
	}
	if this[2] < min[2] {
		this[2] = min[2]
	}
	if this[3] < min[3] {
		this[3] = min[3]
	}

	if this[0] > max[0] {
		this[0] = max[0]
	}
	if this[1] > max[1] {
		this[1] = max[1]
	}
	if this[2] > max[2] {
		this[2] = max[2]
	}
	if this[3] > max[3] {
		this[3] = max[3]
	}

	return this
}

// clamps each element between 0 and 1
func (this *Color) Clamp01() *Color {

	this[0] = Clamp01(this[0])
	this[1] = Clamp01(this[1])
	this[2] = Clamp01(this[2])
	this[3] = Clamp01(this[3])

	return this
}

// sets values from Vec2
func (this *Color) FromVec2(v *Vec2) *Color {

	this[0] = v[0]
	this[1] = v[1]
	this[2] = 0
	this[3] = 1

	return this
}

// sets values from Vec3
func (this *Color) FromVec3(v *Vec3) *Color {

	this[0] = v[0]
	this[1] = v[1]
	this[2] = v[2]
	this[3] = 1

	return this
}

// sets values from Vec4
func (this *Color) FromVec4(v *Vec4) *Color {

	this[0] = v[0]
	this[1] = v[1]
	this[2] = v[2]
	this[3] = v[3]

	return this
}

// checks if this equals other
func (this *Color) Equals(other *Color) bool {

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
func (this *Color) String() string {

	return fmt.Sprintf("Color[ %f, %f, %f, %f ]", this[0], this[1], this[2], this[3])
}
