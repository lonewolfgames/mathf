package mathf

import (
	"math"
	"fmt"
)


// float32 Array representing a 3x3 Matrix
type Mat32 [6]float32

// returns new Mat32
func NewMat32( m11, m12, m13, m21, m22, m23 float32 ) *Mat32{
	this := new( Mat32 )
	
	this[0], this[2], this[4] = m11, m12, m13
	this[1], this[3], this[5] = m21, m22, m23
	
	return this
}

// returns a copy of this
func ( this *Mat32 ) Clone() *Mat32{
	
	return new( Mat32 ).Copy( this )
}

// copies other
func ( this *Mat32 ) Copy( other *Mat32 ) *Mat32{
	
	this[0], this[2], this[4] = other[0], other[2], other[4]
	this[1], this[3], this[5] = other[1], other[3], other[5]
	
	return this
}

// sets this from values
func ( this *Mat32 ) Set( m11, m12, m13, m21, m22, m23 float32 ) *Mat32{
	
	this[0], this[2], this[4] = m11, m12, m13
	this[1], this[3], this[5] = m21, m22, m23
	
	return this
}

// mutiples this by other
func ( this *Mat32 ) Mul( other *Mat32 ) *Mat32{
	a11, a12, a13 := this[0], this[2], this[4]
	a21, a22, a23 := this[1], this[3], this[5]
	
	b11, b12, b13 := other[0], other[2], other[4]
	b21, b22, b23 := other[1], other[3], other[5]
	
	this[0] = a11 * b11 + a21 * b12
	this[1] = a12 * b11 + a22 * b12
	
	this[2] = a11 * b21 + a21 * b22
	this[3] = a12 * b21 + a22 * b22
	
	this[4] = a11 * b13 + a12 * b23 + a13
    this[5] = a21 * b13 + a22 * b23 + a23
	
	return this
}

// mutiples a and b saves in this
func ( this *Mat32 ) MMul( a, b *Mat32 ) *Mat32{
	a11, a12, a13 := a[0], a[2], a[4]
	a21, a22, a23 := a[1], a[3], a[5]
	
	b11, b12, b13 := b[0], b[2], b[4]
	b21, b22, b23 := b[1], b[3], b[5]
	
	this[0] = a11 * b11 + a21 * b12
	this[1] = a12 * b11 + a22 * b12
	
	this[2] = a11 * b21 + a21 * b22
	this[3] = a12 * b21 + a22 * b22
	
	this[4] = a11 * b13 + a12 * b23 + a13
    this[5] = a21 * b13 + a22 * b23 + a23
	
	return this
}

// mutiples this by scalar
func ( this *Mat32 ) SMul( s float32 ) *Mat32{
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	this[4] *= s
	this[5] *= s
	
	return this
}

// divides this by scalar
func ( this *Mat32 ) SDiv( s float32 ) *Mat32{
	if s != 0 { s = 1 / s }
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	this[4] *= s
	this[5] *= s
	
	return this
}

// returns inverse of this
func ( this *Mat32 ) Inverse() *Mat32{
	
	m11, m12, m13 := this[0], this[2], this[4]
	m21, m22, m23 := this[1], this[3], this[5]
	
	det := m11 * m22 - m12 * m21

	if det == 0 { return this.Identity() }
	det = 1 / det

	this[0] = m22 * det
	this[1] = -m12 * det
	this[2] = -m21 * det
	this[3] = m11 * det

	this[4] = ( m12 * m23 - m22 * m13 ) * det
	this[5] = ( m21 * m13 - m11 * m23 ) * det
	
	return this
}

// saves inverse of other in this
func ( this *Mat32 ) MInverse( other *Mat32 ) *Mat32{
	
	m11, m12, m13 := other[0], other[2], other[4]
	m21, m22, m23 := other[1], other[3], other[5]
	
	det := m11 * m22 - m12 * m21

	if det == 0 { return this.Identity() }
	det = 1 / det

	this[0] = m22 * det
	this[1] = -m12 * det
	this[2] = -m21 * det
	this[3] = m11 * det

	this[4] = ( m12 * m23 - m22 * m13 ) * det
	this[5] = ( m21 * m13 - m11 * m23 ) * det
	
	return this
}

// returns this as identity
func ( this *Mat32 ) Identity() *Mat32{
	
	this[0] = 1
	this[1] = 0
	this[2] = 0
	this[3] = 1
	this[4] = 0
	this[5] = 0
	
	return this
}

// transposes this across diagonal
func ( this *Mat32 ) Transpose() *Mat32{
	
	tmp := this[1]
	this[1] = this[2]
	this[2] = tmp
	
	return this
}

// returns the determinant
func ( this *Mat32 ) Determinant() float32{

	return this[0] * this[3] - this[2] * this[1]
}

// composes matrix from position, scale and rotation
func ( this *Mat32 ) Compose( position, scale *Vec2, rotation float32 ) *Mat32{
	sx, sy := scale[0], scale[1]
	c, s := float32(math.Cos( float64(rotation) )), float32(math.Sin( float64(rotation) ))
	
	this[0] = c * sx
	this[1] = s * sx
	this[2] = -s * sy
	this[3] = c * sy
	
	this[4] = position[0]
	this[5] = position[1]
	
	return this
}

// decomposes matrix into a position and scale Vec2 and returns rotation angle in radians
func ( this *Mat32 ) Decompose( position, scale *Vec2 ) float32{
	m11, m12 := this[0], this[1]
	
	sx := scale.Set( m11, m12 ).Length()
	sy := scale.Set( this[2], this[3] ).Length()
	
	position[0] = this[4]
	position[1] = this[5]
	
	scale[0] = sx
	scale[1] = sy
	
	return float32(math.Atan2( float64(m12), float64(m11) ))
}

// extracts others position into this
func ( this *Mat32 ) ExtractPosition( other *Mat32 ) *Mat32{
	
	this[4] = other[4]
	this[5] = other[5]
	
	return this
}

// extracts others rotation into this
func ( this *Mat32 ) ExtractRotation( other *Mat32 ) *Mat32{
	
	m11, m12 := other[0], other[2]
	m21, m22 := other[1], other[3]
	
	x := m11 * m11 + m21 * m21
	y := m12 * m12 + m22 * m22
	
	var sx, sy float32 = 0, 0
	
	if x != 0 { sx = 1 / float32(math.Sqrt( float64(x) )) }
	if y != 0 { sy = 1 / float32(math.Sqrt( float64(y) )) }
	
	this[0] = m11 * sx
	this[1] = m21 * sx
	
	this[2] = m12 * sy
	this[3] = m22 * sy
	
	return this
}

// makes this look from eye to target
func ( this *Mat32 ) LookAt( eye, target *Vec2 ) *Mat32{
	x, y := target[0] - eye[0], target[1] - eye[1]
	a := float32(math.Atan2( float64(y), float64(x) )) - HALF_PI
	c, s := float32(math.Cos( float64(a) )), float32(math.Sin( float64(a) ))
	
	this[0] = c
	this[1] = s
	this[2] = -s
	this[3] = c
	
	return this
}

// sets position from Vec2
func ( this *Mat32 ) SetPosition( position *Vec2 ) *Mat32{
	
	this[4], this[5] = position[0], position[1]
	
	return this
}

// sets rotation from angle
func ( this *Mat32 ) SetRotation( angle float32 ) *Mat32{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	
	this[0], this[2] = c, -s
	this[1], this[3] = s, c
	
	return this
}

// returns rotation in radians
func ( this *Mat32 ) GetRotation() float32{
	
	return float32(math.Atan2( float64(this[1]), float64(this[0]) ))
}

// translates this by Vec2
func ( this *Mat32 ) Translate( v *Vec2 ) *Mat32{
	x, y := v[0], v[1]
	
	this[4] = this[0] * x + this[2] * y + this[4]
    this[5] = this[1] * x + this[3] * y + this[5]
	
	return this
}

// rotates this by angle
func ( this *Mat32 ) Rotate( angle float32 ) *Mat32{
	
	m11, m12 := this[0], this[2]
	m21, m22 := this[1], this[3]
	
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	
	this[0] = m11 * c + m12 * s
	this[1] = m11 * -s + m12 * c
	this[2] = m21 * c + m22 * s
	this[3] = m21 * -s + m22 * c
	
	return this
}

// scales this by Vec2
func ( this *Mat32 ) Scale( v *Vec2 ) *Mat32{
	x, y := v[0], v[1]
	
	this[0] *= x
	this[1] *= x
	this[4] *= x

	this[2] *= y
	this[3] *= y
	this[5] *= y
	
	return this
}

// returns translation matrix
func ( this *Mat32 ) MakeTranslation( x, y float32 ) *Mat32{
	
	this[0], this[2], this[4] = 1, 0, x
	this[1], this[3], this[5] = 0, 1, y
	
	return this
}

// returns scale matrix
func ( this *Mat32 ) MakeScale( x, y float32 ) *Mat32{
	
	this[0], this[2], this[4] = x, 0, 0
	this[1], this[3], this[5] = 0, y, 0
	
	return this
}

// returns rotation matrix
func ( this *Mat32 ) MakeRotation( angle float32 ) *Mat32{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	
	this[0], this[2], this[4] = c, -s, 0
	this[1], this[3], this[5] = s, c, 0
	
	return this
}

// returns orthographic matrix
func ( this *Mat32 ) MakeOrthographic( left, right, bottom, top float32 ) *Mat32{
	w := 1 / ( right - left )
	h := 1 / ( top - bottom )
	x := ( right + left ) * w
	y := ( top + bottom ) * h
	
	this[0] = 2 * w
	this[1] = 0
	this[2] = 0
	this[3] = 2 * h
	
	this[4] = -x
	this[5] = -y
	
	return this
}

// sets values from Mat2
func ( this *Mat32 ) FromMat2( m *Mat2 ) *Mat32{
	
	this[0] = m[0]
	this[1] = m[1]
	this[2] = m[2]
	this[3] = m[3]
	this[4] = 0
	this[5] = 0
	
	return this
}

// sets values from Mat3
func ( this *Mat32 ) FromMat3( m *Mat3 ) *Mat32{
	
	this[0] = m[0]
	this[1] = m[1]
	this[2] = m[3]
	this[3] = m[4]
	this[4] = 0
	this[5] = 0
	
	return this
}

// sets values from Mat4
func ( this *Mat32 ) FromMat4( m *Mat4 ) *Mat32{
	
	this[0] = m[0]
	this[1] = m[1]
	this[2] = m[4]
	this[3] = m[5]
	this[4] = m[12]
	this[5] = m[13]
	
	return this
}

// checks if this equals other
func ( this *Mat32 ) Equals( other *Mat32 ) bool{
	
	if this[0] != other[0] { return false }
	if this[1] != other[1] { return false }
	if this[2] != other[2] { return false }
	if this[3] != other[3] { return false }
	
	return true
}

// returns this as string type
func ( this *Mat32 ) String() string{
	
	return fmt.Sprintf("Mat32[\n %f, %f, %f,\n %f, %f, %f\n ]", this[0], this[2], this[4], this[1], this[3], this[5] )
}