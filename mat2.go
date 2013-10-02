package mathf

import (
	"math"
	"fmt"
)


// float32 Array representing a 2x2 Matrix
type Mat2 [4]float32

// returns new Mat2
func NewMat2() *Mat2{
	this := new( Mat2 )
	
	this[0], this[2] = 1, 0
	this[1], this[3] = 0, 1
	
	return this
}

// returns a copy of this
func ( this *Mat2 ) Clone() *Mat2{
	
	return new( Mat2 ).Copy( this )
}

// copies other
func ( this *Mat2 ) Copy( other *Mat2 ) *Mat2{
	
	this[0], this[2] = other[0], other[2]
	this[1], this[3] = other[1], other[3]
	
	return this
}

// sets this from values
func ( this *Mat2 ) Set( m11, m12, m21, m22 float32 ) *Mat2{
	
	this[0], this[2] = m11, m12
	this[1], this[3] = m21, m22
	
	return this
}

// mutiples this by other
func ( this *Mat2 ) Mul( other *Mat2 ) *Mat2{
	a11, a12 := this[0], this[2]
	a21, a22 := this[1], this[3]
	
	b11, b12 := other[0], other[2]
	b21, b22 := other[1], other[3]
	
	this[0] = a11 * b11 + a21 * b12
	this[1] = a12 * b11 + a22 * b12
	
	this[2] = a11 * b21 + a21 * b22
	this[3] = a12 * b21 + a22 * b22
	
	return this
}

// mutiples a and b saves in this
func ( this *Mat2 ) MMul( a, b *Mat2 ) *Mat2{
	a11, a12 := a[0], a[2]
	a21, a22 := a[1], a[3]
	
	b11, b12 := b[0], b[2]
	b21, b22 := b[1], b[3]
	
	this[0] = a11 * b11 + a21 * b12
	this[1] = a12 * b11 + a22 * b12
	
	this[2] = a11 * b21 + a21 * b22
	this[3] = a12 * b21 + a22 * b22
	
	return this
}

// mutiples this by scalar
func ( this *Mat2 ) SMul( s float32 ) *Mat2{
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	
	return this
}

// divides this by scalar
func ( this *Mat2 ) SDiv( s float32 ) *Mat2{
	if s != 0 { s = 1 / s }
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	
	return this
}

// returns inverse of this
func ( this *Mat2 ) Inverse() *Mat2{
	
	m11, m12 := this[0], this[2]
	m21, m22 := this[1], this[3]
	
	det := m11 * m22 - m12 * m22
	if det == 0 { return this.Identity() }
	
	this[0] = m22 * det
	this[1] = -m12 * det
	this[2] = -m21 * det
	this[3] = m11 * det
	
	return this
}

// saves inverse of other in this
func ( this *Mat2 ) MInverse( other *Mat2 ) *Mat2{
	
	m11, m12 := other[0], other[2]
	m21, m22 := other[1], other[3]
	
	det := m11 * m22 - m12 * m22
	if det == 0 { return this.Identity() }
	
	this[0] = m22 * det
	this[1] = -m12 * det
	this[2] = -m21 * det
	this[3] = m11 * det
	
	return this
}

// returns this as identity
func ( this *Mat2 ) Identity() *Mat2{
	
	this[0] = 1
	this[1] = 0
	this[2] = 0
	this[3] = 1
	
	return this
}

// transposes this across diagonal
func ( this *Mat2 ) Transpose() *Mat2{
	
	this[1], this[2] = this[2], this[1]
	
	return this
}

// returns the determinant
func ( this *Mat2 ) Determinant() float32{

	return this[0] * this[3] - this[2] * this[1]
}

// returns scale matrix
func ( this *Mat2 ) MakeScale( x, y float32 ) *Mat2{
	
	this[0], this[2] = x, 0
	this[1], this[3] = 0, y
	
	return this
}

// returns rotation matrix
func ( this *Mat2 ) MakeRotation( angle float32 ) *Mat2{
	s, c := float32( math.Sin( float64( angle ) ) ), float32( math.Cos( float64( angle ) ) )
	
	this[0], this[2] = c, -s
	this[1], this[3] = s, c
	
	return this
}

// returns rotation of this matrix
func ( this *Mat2 ) GetRotation() float32{
	
	return float32( math.Atan2( float64( this[1] ), float64( this[0] ) ) )
}

// rotates matrix by angle
func ( this *Mat2 ) Rotate( angle float32 ) *Mat2{
	
	m11, m12 := this[0], this[2]
	m21, m22 := this[1], this[3]
	
	s, c := float32( math.Sin( float64( angle ) ) ), float32( math.Cos( float64( angle ) ) )
	
	this[0] = m11 * c + m12 * s
	this[1] = m11 * -s + m12 * c
	this[2] = m21 * c + m22 * s
	this[3] = m21 * -s + m22 * c
	
	return this
}

// sets values from Mat32
func ( this *Mat2 ) FromMat32( m *Mat32 ) *Mat2{
	
	this[0] = m[0]
	this[1] = m[1]
	this[2] = m[2]
	this[3] = m[3]
	
	return this
}

// checks if this equals other
func ( this *Mat2 ) Equals( other *Mat2 ) bool{
	
	if this[0] != other[0] { return false }
	if this[1] != other[1] { return false }
	if this[2] != other[2] { return false }
	if this[3] != other[3] { return false }
	
	return true
}

// returns this as string
func ( this *Mat2 ) String() string{
	
	return fmt.Sprintf("Mat2[\n %f, %f,\n %f, %f\n ]", this[0], this[2], this[1], this[3] )
}