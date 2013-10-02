package mathf

import (
	"math"
	"fmt"
)


// float32 Array representing a 3x3 Matrix
type Mat4 [16]float32

// returns new Mat4
func NewMat4( m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44 float32 ) *Mat4{
	this := new( Mat4 )
	
	this[0], this[4], this[8], this[12] = m11, m12, m13, m14
	this[1], this[5], this[9], this[13] = m21, m22, m23, m24
	this[2], this[6], this[10], this[14] = m31, m32, m33, m34
	this[3], this[7], this[11], this[15] = m41, m42, m43, m44
	
	return this
}

// returns a copy of this
func ( this *Mat4 ) Clone() *Mat4{
	
	return new( Mat4 ).Copy( this )
}

// copies other
func ( this *Mat4 ) Copy( other *Mat4 ) *Mat4{
	
	this[0], this[4], this[8], this[12] = other[0], other[4], other[8], other[12]
	this[1], this[5], this[9], this[13] = other[1], other[5], other[9], other[13]
	this[2], this[6], this[10], this[14] = other[2], other[6], other[10], other[14]
	this[3], this[7], this[11], this[15] = other[3], other[7], other[11], other[15]
	
	return this
}

// sets this from values
func ( this *Mat4 ) Set( m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44 float32 ) *Mat4{
	
	this[0], this[4], this[8], this[12] = m11, m12, m13, m14
	this[1], this[5], this[9], this[13] = m21, m22, m23, m24
	this[2], this[6], this[10], this[14] = m31, m32, m33, m34
	this[3], this[7], this[11], this[15] = m41, m42, m43, m44
	
	return this
}

// mutiples this by other
func ( this *Mat4 ) Mul( other *Mat4 ) *Mat4{
	a11, a12, a13, a14 := this[0], this[4], this[8], this[12]
	a21, a22, a23, a24 := this[1], this[5], this[9], this[13]
	a31, a32, a33, a34 := this[2], this[6], this[10], this[14]
	a41, a42, a43, a44 := this[3], this[7], this[11], this[15]
	
	b11, b12, b13, b14 := other[0], other[4], other[8], other[12]
	b21, b22, b23, b24 := other[1], other[5], other[9], other[13]
	b31, b32, b33, b34 := other[2], other[6], other[10], other[14]
	b41, b42, b43, b44 := other[3], other[7], other[11], other[15]
	
	this[0] = a11 * b11 + a12 * b21 + a13 * b31 + a14 * b41
	this[4] = a11 * b12 + a12 * b22 + a13 * b32 + a14 * b42
	this[8] = a11 * b13 + a12 * b23 + a13 * b33 + a14 * b43
	this[12] = a11 * b14 + a12 * b24 + a13 * b34 + a14 * b44

	this[1] = a21 * b11 + a22 * b21 + a23 * b31 + a24 * b41
	this[5] = a21 * b12 + a22 * b22 + a23 * b32 + a24 * b42
	this[9] = a21 * b13 + a22 * b23 + a23 * b33 + a24 * b43
	this[13] = a21 * b14 + a22 * b24 + a23 * b34 + a24 * b44

	this[2] = a31 * b11 + a32 * b21 + a33 * b31 + a34 * b41
	this[6] = a31 * b12 + a32 * b22 + a33 * b32 + a34 * b42
	this[10] = a31 * b13 + a32 * b23 + a33 * b33 + a34 * b43
	this[14] = a31 * b14 + a32 * b24 + a33 * b34 + a34 * b44

	this[3] = a41 * b11 + a42 * b21 + a43 * b31 + a44 * b41
	this[7] = a41 * b12 + a42 * b22 + a43 * b32 + a44 * b42
	this[11] = a41 * b13 + a42 * b23 + a43 * b33 + a44 * b43
	this[15] = a41 * b14 + a42 * b24 + a43 * b34 + a44 * b44
	
	return this
}

// mutiples a and b saves in this
func ( this *Mat4 ) MMul( a, b *Mat4 ) *Mat4{
	a11, a12, a13, a14 := a[0], a[4], a[8], a[12]
	a21, a22, a23, a24 := a[1], a[5], a[9], a[13]
	a31, a32, a33, a34 := a[2], a[6], a[10], a[14]
	a41, a42, a43, a44 := a[3], a[7], a[11], a[15]
	
	b11, b12, b13, b14 := b[0], b[4], b[8], b[12]
	b21, b22, b23, b24 := b[1], b[5], b[9], b[13]
	b31, b32, b33, b34 := b[2], b[6], b[10], b[14]
	b41, b42, b43, b44 := b[3], b[7], b[11], b[15]
	
	this[0] = a11 * b11 + a12 * b21 + a13 * b31 + a14 * b41
	this[4] = a11 * b12 + a12 * b22 + a13 * b32 + a14 * b42
	this[8] = a11 * b13 + a12 * b23 + a13 * b33 + a14 * b43
	this[12] = a11 * b14 + a12 * b24 + a13 * b34 + a14 * b44

	this[1] = a21 * b11 + a22 * b21 + a23 * b31 + a24 * b41
	this[5] = a21 * b12 + a22 * b22 + a23 * b32 + a24 * b42
	this[9] = a21 * b13 + a22 * b23 + a23 * b33 + a24 * b43
	this[13] = a21 * b14 + a22 * b24 + a23 * b34 + a24 * b44

	this[2] = a31 * b11 + a32 * b21 + a33 * b31 + a34 * b41
	this[6] = a31 * b12 + a32 * b22 + a33 * b32 + a34 * b42
	this[10] = a31 * b13 + a32 * b23 + a33 * b33 + a34 * b43
	this[14] = a31 * b14 + a32 * b24 + a33 * b34 + a34 * b44

	this[3] = a41 * b11 + a42 * b21 + a43 * b31 + a44 * b41
	this[7] = a41 * b12 + a42 * b22 + a43 * b32 + a44 * b42
	this[11] = a41 * b13 + a42 * b23 + a43 * b33 + a44 * b43
	this[15] = a41 * b14 + a42 * b24 + a43 * b34 + a44 * b44
	
	return this
}

// mutiples this by scalar
func ( this *Mat4 ) SMul( s float32 ) *Mat4{
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	this[4] *= s
	this[5] *= s
	this[6] *= s
	this[7] *= s
	this[8] *= s
	this[9] *= s
	this[10] *= s
	this[11] *= s
	this[12] *= s
	this[13] *= s
	this[14] *= s
	this[15] *= s
	
	return this
}

// divides this by scalar
func ( this *Mat4 ) SDiv( s float32 ) *Mat4{
	if s != 0 { s = 1 / s }
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	this[4] *= s
	this[5] *= s
	this[6] *= s
	this[7] *= s
	this[8] *= s
	this[9] *= s
	this[10] *= s
	this[11] *= s
	this[12] *= s
	this[13] *= s
	this[14] *= s
	this[15] *= s
	
	return this
}

// returns inverse of this
func ( this *Mat4 ) Inverse() *Mat4{
	
	m11, m12, m13, m14 := this[0], this[4], this[8], this[12]
	m21, m22, m23, m24 := this[1], this[5], this[9], this[13]
	m31, m32, m33, m34 := this[2], this[6], this[10], this[14]
	m41, m42, m43, m44 := this[3], this[7], this[11], this[15]
	
	te0 := m23 * m34 * m42 - m24 * m33 * m42 + m24 * m32 * m43 - m22 * m34 * m43 - m23 * m32 * m44 + m22 * m33 * m44
	te4 := m14 * m33 * m42 - m13 * m34 * m42 - m14 * m32 * m43 + m12 * m34 * m43 + m13 * m32 * m44 - m12 * m33 * m44
	te8 := m13 * m24 * m42 - m14 * m23 * m42 + m14 * m22 * m43 - m12 * m24 * m43 - m13 * m22 * m44 + m12 * m23 * m44
	te12 := m14 * m23 * m32 - m13 * m24 * m32 - m14 * m22 * m33 + m12 * m24 * m33 + m13 * m22 * m34 - m12 * m23 * m34
	
	det := m11 * te0 + m21 * te4 + m31 * te8 + m41 * te12
	
	if det == 0 { return this.Identity() }
	det = 1 / det
	
	this[0] = te0 * det
	this[4] = te4 * det
	this[8] = te8 * det
	this[12] = te12 * det
	this[1] = ( m24 * m33 * m41 - m23 * m34 * m41 - m24 * m31 * m43 + m21 * m34 * m43 + m23 * m31 * m44 - m21 * m33 * m44 ) * det
	this[5] = ( m13 * m34 * m41 - m14 * m33 * m41 + m14 * m31 * m43 - m11 * m34 * m43 - m13 * m31 * m44 + m11 * m33 * m44 ) * det
	this[9] = ( m14 * m23 * m41 - m13 * m24 * m41 - m14 * m21 * m43 + m11 * m24 * m43 + m13 * m21 * m44 - m11 * m23 * m44 ) * det
	this[13] = ( m13 * m24 * m31 - m14 * m23 * m31 + m14 * m21 * m33 - m11 * m24 * m33 - m13 * m21 * m34 + m11 * m23 * m34 ) * det
	this[2] = ( m22 * m34 * m41 - m24 * m32 * m41 + m24 * m31 * m42 - m21 * m34 * m42 - m22 * m31 * m44 + m21 * m32 * m44 ) * det
	this[6] = ( m14 * m32 * m41 - m12 * m34 * m41 - m14 * m31 * m42 + m11 * m34 * m42 + m12 * m31 * m44 - m11 * m32 * m44 ) * det
	this[10] = ( m12 * m24 * m41 - m14 * m22 * m41 + m14 * m21 * m42 - m11 * m24 * m42 - m12 * m21 * m44 + m11 * m22 * m44 ) * det
	this[14] = ( m14 * m22 * m31 - m12 * m24 * m31 - m14 * m21 * m32 + m11 * m24 * m32 + m12 * m21 * m34 - m11 * m22 * m34 ) * det
	this[3] = ( m23 * m32 * m41 - m22 * m33 * m41 - m23 * m31 * m42 + m21 * m33 * m42 + m22 * m31 * m43 - m21 * m32 * m43 ) * det
	this[7] = ( m12 * m33 * m41 - m13 * m32 * m41 + m13 * m31 * m42 - m11 * m33 * m42 - m12 * m31 * m43 + m11 * m32 * m43 ) * det
	this[11] = ( m13 * m22 * m41 - m12 * m23 * m41 - m13 * m21 * m42 + m11 * m23 * m42 + m12 * m21 * m43 - m11 * m22 * m43 ) * det
	this[15] = ( m12 * m23 * m31 - m13 * m22 * m31 + m13 * m21 * m32 - m11 * m23 * m32 - m12 * m21 * m33 + m11 * m22 * m33 ) * det
	
	return this
}

// saves inverse of other in this
func ( this *Mat4 ) MInverse( other *Mat4 ) *Mat4{
	
	m11, m12, m13, m14 := other[0], other[4], other[8], other[12]
	m21, m22, m23, m24 := other[1], other[5], other[9], other[13]
	m31, m32, m33, m34 := other[2], other[6], other[10], other[14]
	m41, m42, m43, m44 := other[3], other[7], other[11], other[15]
	
	te0 := m23 * m34 * m42 - m24 * m33 * m42 + m24 * m32 * m43 - m22 * m34 * m43 - m23 * m32 * m44 + m22 * m33 * m44
	te4 := m14 * m33 * m42 - m13 * m34 * m42 - m14 * m32 * m43 + m12 * m34 * m43 + m13 * m32 * m44 - m12 * m33 * m44
	te8 := m13 * m24 * m42 - m14 * m23 * m42 + m14 * m22 * m43 - m12 * m24 * m43 - m13 * m22 * m44 + m12 * m23 * m44
	te12 := m14 * m23 * m32 - m13 * m24 * m32 - m14 * m22 * m33 + m12 * m24 * m33 + m13 * m22 * m34 - m12 * m23 * m34
	
	det := m11 * te0 + m21 * te4 + m31 * te8 + m41 * te12
	
	if det == 0 { return this.Identity() }
	det = 1 / det
	
	this[0] = te0 * det
	this[4] = te4 * det
	this[8] = te8 * det
	this[12] = te12 * det
	this[1] = ( m24 * m33 * m41 - m23 * m34 * m41 - m24 * m31 * m43 + m21 * m34 * m43 + m23 * m31 * m44 - m21 * m33 * m44 ) * det
	this[5] = ( m13 * m34 * m41 - m14 * m33 * m41 + m14 * m31 * m43 - m11 * m34 * m43 - m13 * m31 * m44 + m11 * m33 * m44 ) * det
	this[9] = ( m14 * m23 * m41 - m13 * m24 * m41 - m14 * m21 * m43 + m11 * m24 * m43 + m13 * m21 * m44 - m11 * m23 * m44 ) * det
	this[13] = ( m13 * m24 * m31 - m14 * m23 * m31 + m14 * m21 * m33 - m11 * m24 * m33 - m13 * m21 * m34 + m11 * m23 * m34 ) * det
	this[2] = ( m22 * m34 * m41 - m24 * m32 * m41 + m24 * m31 * m42 - m21 * m34 * m42 - m22 * m31 * m44 + m21 * m32 * m44 ) * det
	this[6] = ( m14 * m32 * m41 - m12 * m34 * m41 - m14 * m31 * m42 + m11 * m34 * m42 + m12 * m31 * m44 - m11 * m32 * m44 ) * det
	this[10] = ( m12 * m24 * m41 - m14 * m22 * m41 + m14 * m21 * m42 - m11 * m24 * m42 - m12 * m21 * m44 + m11 * m22 * m44 ) * det
	this[14] = ( m14 * m22 * m31 - m12 * m24 * m31 - m14 * m21 * m32 + m11 * m24 * m32 + m12 * m21 * m34 - m11 * m22 * m34 ) * det
	this[3] = ( m23 * m32 * m41 - m22 * m33 * m41 - m23 * m31 * m42 + m21 * m33 * m42 + m22 * m31 * m43 - m21 * m32 * m43 ) * det
	this[7] = ( m12 * m33 * m41 - m13 * m32 * m41 + m13 * m31 * m42 - m11 * m33 * m42 - m12 * m31 * m43 + m11 * m32 * m43 ) * det
	this[11] = ( m13 * m22 * m41 - m12 * m23 * m41 - m13 * m21 * m42 + m11 * m23 * m42 + m12 * m21 * m43 - m11 * m22 * m43 ) * det
	this[15] = ( m12 * m23 * m31 - m13 * m22 * m31 + m13 * m21 * m32 - m11 * m23 * m32 - m12 * m21 * m33 + m11 * m22 * m33 ) * det
	
	return this
}

// returns this as identity
func ( this *Mat4 ) Identity() *Mat4{
	
	this[0] = 1
	this[1] = 0
	this[2] = 0
	this[3] = 0
	this[4] = 0
	this[5] = 1
	this[6] = 0
	this[7] = 0
	this[8] = 0
	this[9] = 0
	this[10] = 1
	this[11] = 0
	this[12] = 0
	this[13] = 0
	this[14] = 0
	this[15] = 1
	
	return this
}

// transposes this across diagonal
func ( this *Mat4 ) Transpose() *Mat4{
	
	tmp := this[1]
	this[1] = this[4]
	this[4] = tmp
	
	tmp = this[2]
	this[2] = this[8]
	this[8] = tmp
	
	tmp = this[6]
	this[6] = this[9]
	this[9] = tmp
	
	tmp = this[3]
	this[3] = this[12]
	this[12] = tmp
	
	tmp = this[7]
	this[7] = this[13]
	this[13] = tmp
	
	tmp = this[11]
	this[11] = this[14]
	this[14] = tmp
	
	return this
}

// returns the determinant
func ( this *Mat4 ) Determinant() float32{
	
	m11, m12, m13, m14 := this[0], this[4], this[8], this[12]
	m21, m22, m23, m24 := this[1], this[5], this[9], this[13]
	m31, m32, m33, m34 := this[2], this[6], this[10], this[14]
	m41, m42, m43, m44 := this[3], this[7], this[11], this[15]
	
	return m41 * ( m14 * m23 * m32 - m13 * m24 * m32 - m14 * m22 * m33 + m12 * m24 * m33 + m13 * m22 * m34 - m12 * m23 * m34 ) + m42 * ( m11 * m23 * m34 - m11 * m24 * m33 + m14 * m21 * m33 - m13 * m21 * m34 + m13 * m24 * m31 - m14 * m23 * m31 ) + m43 * ( m11 * m24 * m32 - m11 * m22 * m34 - m14 * m21 * m32 + m12 * m21 * m34 + m14 * m22 * m31 - m12 * m24 * m31 ) + m44 * (-m13 * m22 * m31 - m11 * m23 * m32 + m11 * m22 * m33 + m13 * m21 * m32 - m12 * m21 * m33 + m12 * m23 * m31 )
}

// composes matrix from position, scale and rotation
func ( this *Mat4 ) Compose( position, scale *Vec3, rotation *Quat ) *Mat4{
	x, y, z, w := rotation[0], rotation[1], rotation[2], rotation[3]
	x2, y2, z2 := x + x, y + y, z + z
	xx, xy, xz := x * x2, x * y2, x * z2
	yy, yz, zz := y * y2, y * z2, z * z2
	wx, wy, wz := w * x2, w * y2, w * z2
	
	sx, sy, sz := scale[0], scale[1], scale[2]
	
	this[0] = ( 1 - ( yy + zz ) ) * sx
	this[4] = ( xy - wz ) * sy
	this[8] = ( xz + wy ) * sz
	
	this[1] = ( xy + wz ) * sx
	this[5] = ( 1 - ( xx + zz ) ) * sy
	this[9] = ( yz - wx ) * sz
	
	this[2] = ( xz - wy ) * sx
	this[6] = ( yz + wx ) * sy
	this[10] = ( 1 - ( xx + yy ) ) * sz
	
	this[3] = 0
	this[7] = 0
	this[11] = 0
	
	this[12] = position[0]
	this[13] = position[1]
	this[14] = position[2]
	this[15] = 1
	
	return this
}

// decomposes matrix into a position and scale Vec2 and a rotation angle in radians
func ( this *Mat4 ) Decompose( position, scale *Vec3, rotation *Quat ) *Mat4{
	
	m11, m12, m13 := this[0], this[4], this[8]
	m21, m22, m23 := this[1], this[5], this[9]
	m31, m32, m33 := this[2], this[6], this[10]
	
	var trace, x, y, z, w, s float32
	
	sx := scale.Set( m11, m21, m31 ).Length()
	sy := scale.Set( m12, m22, m32 ).Length()
	sz := scale.Set( m13, m23, m33 ).Length()
	
	invSx := 1 / sx
	invSy := 1 / sy
	invSz := 1 / sz
	
	scale[0] = sx
	scale[1] = sy
	scale[2] = sz
	
	position[0] = this[12]
	position[1] = this[13]
	position[2] = this[14]
	
	m11 *= invSx
	m12 *= invSy
	m13 *= invSz
	m21 *= invSx
	m22 *= invSy
	m23 *= invSz
	m31 *= invSx
	m32 *= invSy
	m33 *= invSz
	
	trace = m11 + m22 + m33
	
	if( trace > 0 ){
		s = 0.5 / float32(math.Sqrt( float64(trace + 1) ))
		
		w = 0.25 / s
		x = ( m32 - m23 ) * s
		y = ( m13 - m31 ) * s
		z = ( m21 - m12 ) * s
		
	}else if( m11 > m22 && m11 > m33 ){
		s = 2 * float32(math.Sqrt( float64(1 + m11 - m22 - m33) ))
		
		w = ( m32 - m23 ) / s
		x = 0.25 * s
		y = ( m12 + m21 ) / s
		z = ( m13 + m31 ) / s
		
	}else if( m22 > m33 ){
		s = 2 * float32(math.Sqrt( float64(1 + m22 - m11 - m33) ))
		
		w = ( m13 - m31 ) / s
		x = ( m12 + m21 ) / s
		y = 0.25 * s
		z = ( m23 + m32 ) / s
		
	}else{
		s = 2 * float32(math.Sqrt( float64(1 + m33 - m11 - m22) ))
		
		w = ( m21 - m12 ) / s
		x = ( m13 + m31 ) / s
		y = ( m23 + m32 ) / s
		z = 0.25 * s
	}
	
	rotation[0] = x
	rotation[1] = y
	rotation[2] = w
	rotation[3] = z
	
	return this
}

// extracts others position and saves it in this
func ( this *Mat4 ) ExtractPosition( other *Mat4 ) *Mat4{
	
	this[12] = other[12]
	this[13] = other[13] 
	this[14] = other[14]
	
	return this
}

// extracts others rotation and saves it in this
func ( this *Mat4 ) ExtractRotation( other *Mat4 ) *Mat4{
	m11, m12, m13 := this[0], this[4], this[8]
	m21, m22, m23 := this[1], this[5], this[9]
	m31, m32, m33 := this[2], this[6], this[10]
	
	x := m11 * m11 + m21 * m21 + m31 * m31
	y := m12 * m12 + m22 * m22 + m32 * m32
	z := m13 * m13 + m23 * m23 + m33 * m33
	
	var sx, sy, sz float32 = 0, 0, 0
	
	if x != 0 { sx = 1 / float32(math.Sqrt( float64(x) )) }
	if y != 0 { sy = 1 / float32(math.Sqrt( float64(y) )) }
	if z != 0 { sz = 1 / float32(math.Sqrt( float64(z) )) }
	
	this[0] = m11 * sx
	this[1] = m21 * sx
	this[2] = m31 * sx
	
	this[4] = m12 * sy
	this[5] = m22 * sy
	this[6] = m32 * sy
	
	this[8] = m13 * sz
	this[9] = m23 * sz
	this[10] = m33 * sz
	
	return this
}

// makes this look from eye to target
func ( this *Mat4 ) LookAt( eye, target, up *Vec3 ) *Mat4{
	zx, zy, zz := target[0] - eye[0], target[1] - eye[1], target[2] - eye[2]
	l := zx * zx + zy * zy + zz * zz
	if l != 0 { l = 1 / float32(math.Sqrt( float64(l) )) }
	
	zx *= l
	zy *= l
	zz *= l
	
	xx := up[1] * zz - up[2] * zy
	xy := up[2] * zx - up[0] * zz
	xz := up[0] * zy - up[1] * zx
	
	l = xx * xx + xy * xy + xz * xz
	if l != 0 { l = 1 / float32(math.Sqrt( float64(l) )) }
	
	xx *= l
	xy *= l
	xz *= l
	
	yx := zy * xz - zz * xy
	yy := zz * xx - zx * xz
	yz := zx * xy - zy * xx

	this[0] = xx
	this[4] = yx
	this[8] = zx
	this[1] = xy
	this[5] = yy
	this[9] = zy
	this[2] = xz
	this[6] = yz
	this[10] = zz
	
	return this
}

// sets position of this
func ( this *Mat4 ) SetPosition( position *Vec3 ) *Mat4{
	
	this[12], this[13], this[14] = position[0], position[1], position[2]
	
	return this
}

// translates this by Vec3
func ( this *Mat4 ) Translate( v *Vec3 ) *Mat4{
	x, y, z := v[0], v[1], v[2]
	
	this[12] = this[0] * x + this[4] * y + this[8] * z + this[12]
	this[13] = this[1] * x + this[5] * y + this[9] * z + this[13]
	this[14] = this[2] * x + this[6] * y + this[10] * z + this[14]
	this[15] = this[3] * x + this[7] * y + this[11] * z + this[15]
	
	return this
}

// rotates this by angle along the x axis
func ( this *Mat4 ) RotateX( angle float32 ) *Mat4{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	m12 := this[4]
	m22 := this[5]
	m32 := this[6]
	m42 := this[7]
	m13 := this[8]
	m23 := this[9]
	m33 := this[10]
	m43 := this[11]
	
	this[4] = c * m12 + s * m13
	this[5] = c * m22 + s * m23
	this[6] = c * m32 + s * m33
	this[7] = c * m42 + s * m43
	
	this[8] = c * m13 - s * m12
	this[9] = c * m23 - s * m22
	this[10] = c * m33 - s * m32
	this[11] = c * m43 - s * m42
	
	return this
}

// rotates this by angle along the y axis
func ( this *Mat4 ) RotateY( angle float32 ) *Mat4{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	m11 := this[0]
	m21 := this[1]
	m31 := this[2]
	m41 := this[3]
	m13 := this[8]
	m23 := this[9]
	m33 := this[10]
	m43 := this[11]
	
	this[0] = c * m11 - s * m13
	this[1] = c * m21 - s * m23
	this[2] = c * m31 - s * m33
	this[3] = c * m41 - s * m43
	
	this[8] = c * m13 + s * m11
	this[9] = c * m23 + s * m21
	this[10] = c * m33 + s * m31
	this[11] = c * m43 + s * m41
	
	return this
}

// rotates this by angle along the z axis
func ( this *Mat4 ) RotateZ( angle float32 ) *Mat4{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	m11 := this[0]
	m21 := this[1]
	m31 := this[2]
	m41 := this[3]
	m12 := this[4]
	m22 := this[5]
	m32 := this[6]
	m42 := this[7]
	
	this[0] = c * m11 + s * m12
	this[1] = c * m21 + s * m22
	this[2] = c * m31 + s * m32
	this[3] = c * m41 + s * m42
	
	this[4] = c * m12 - s * m11
	this[5] = c * m22 - s * m21
	this[6] = c * m32 - s * m31
	this[7] = c * m42 - s * m41
	
	return this
}

// scales this by Vec3
func ( this *Mat4 ) Scale( s *Vec3 ) *Mat4{
	x, y, z := s[0], s[1], s[2]
	
	this[0] *= x
	this[4] *= y
	this[8] *= z
	
	this[1] *= x
	this[5] *= y
	this[9] *= z
	
	this[2] *= x
	this[6] *= y
	this[10] *= z
	
	this[3] *= x
	this[7] *= y
	this[11] *= z
	
	return this
}

// returns scale matrix
func ( this *Mat4 ) MakeScale( x, y, z float32 ) *Mat4{
	
	this[0], this[3], this[6] = x, 0, 0
	this[1], this[4], this[7] = 0, y, 0
	this[2], this[5], this[8] = 0, 0, z
	
	return this
}

// returns rotation matrix around the x axis
func ( this *Mat4 ) MakeRotationX( angle float32 ) *Mat4{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	
	this[0], this[3], this[6] = 0, 0, 0
	this[1], this[4], this[7] = 0, c, -s
	this[2], this[5], this[8] = 0, s, c
	
	return this
}

// returns rotation matrix around the y axis
func ( this *Mat4 ) MakeRotationY( angle float32 ) *Mat4{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	
	this[0], this[3], this[6] = c, 0, s
	this[1], this[4], this[7] = 0, 1, 0
	this[2], this[5], this[8] = -s, 0, c
	
	return this
}

// returns rotation matrix around the z axis
func ( this *Mat4 ) MakeRotationZ( angle float32 ) *Mat4{
	c, s := float32(math.Cos( float64(angle) )), float32(math.Sin( float64(angle) ))
	
	this[0], this[3], this[6] = c, -s, 0
	this[1], this[4], this[7] = s, c, 0
	this[2], this[5], this[8] = 0, 0, 1
	
	return this
}

// sets values from Quat
func ( this *Mat4 ) FromQuat( q *Quat ) *Mat4{
	x, y, z, w := q[0], q[1], q[2], q[3]
	x2, y2, z2 := x + x, y + y, z + z
	xx, xy, xz := x * x2, x * y2, x * z2
	yy, yz, zz := y * y2, y * z2, z * z2
	wx, wy, wz := w * x2, w * y2, w * z2
	
	this[0] = 1 - ( yy + zz )
	this[3] = xy - wz
	this[6] = xz + wy
	
	this[1] = xy + wz
	this[4] = 1 - ( xx + zz )
	this[7] = yz - wx
	
	this[2] = xz - wy
	this[5] = yz + wx
	this[8] = 1 - ( xx + yy )
	
	return this
}

// sets values from Mat2
func ( this *Mat4 ) FromMat2( m *Mat2 ) *Mat4{
	
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
func ( this *Mat4 ) FromMat4( m *Mat4 ) *Mat4{
	
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
func ( this *Mat4 ) Equals( other *Mat4 ) bool{
	
	if this[0] != other[0] { return false }
	if this[1] != other[1] { return false }
	if this[2] != other[2] { return false }
	if this[3] != other[3] { return false }
	if this[4] != other[4] { return false }
	if this[5] != other[5] { return false }
	if this[6] != other[6] { return false }
	if this[7] != other[7] { return false }
	if this[8] != other[8] { return false }
	
	return true
}

// returns this as string
func ( this *Mat4 ) String() string{
	
	return fmt.Sprintf("Mat4[\n %f, %f, %f, %f,\n %f, %f, %f, %f,\n %f, %f, %f, %f,\n %f, %f, %f, %f\n ]", this[0], this[4], this[8], this[12], this[1], this[5], this[9], this[13], this[2], this[6], this[10], this[14], this[3], this[7], this[11], this[15] )
}