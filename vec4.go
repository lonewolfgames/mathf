package mathf

import (
	"math"
	"fmt"
)

// float32 Array representing a 4D Vector
type Vec4 [4]float32

// returns new Vec4
func NewVec4( x, y, z, w float32 ) *Vec4{
	this := new( Vec4 )
	
	this[0], this[1], this[2], this[3] = x, y, z, w
	
	return this
}

// returns a copy of this
func ( this *Vec4 ) Clone() *Vec4{
	
	return new( Vec4 ).Copy( this )
}

// copies other
func ( this *Vec4 ) Copy( other *Vec4 ) *Vec4{
	
	this[0], this[1], this[2], this[3] = other[0], other[1], other[2], other[3]
	
	return this
}

// sets this from values
func ( this *Vec4 ) Set( x, y, z, w float32 ) *Vec4{
	
	this[0], this[1], this[2], this[3] = x, y, z, w
	
	return this
}

// adds other to this
func ( this *Vec4 ) Add( other *Vec4 ) *Vec4{
	
	this[0] += other[0]
	this[1] += other[1]
	this[2] += other[2]
	this[3] += other[3]
	
	return this
}

// adds a and b saves in this
func ( this *Vec4 ) VAdd( a, b *Vec4 ) *Vec4{
	
	this[0] = a[0] + b[0]
	this[1] = a[1] + b[1]
	this[2] = a[2] + b[2]
	this[3] = a[3] + b[3]
	
	return this
}

// adds scalar to this
func ( this *Vec4 ) SAdd( s float32 ) *Vec4{
	
	this[0] += s
	this[1] += s
	this[2] += s
	this[3] += s
	
	return this
}

// subtracts other from this
func ( this *Vec4 ) Sub( other *Vec4 ) *Vec4{
	
	this[0] -= other[0]
	this[1] -= other[1]
	this[2] -= other[2]
	this[3] -= other[3]
	
	return this
}

// subtracts a and b saves in this
func ( this *Vec4 ) VSub( a, b *Vec4 ) *Vec4{
	
	this[0] = a[0] - b[0]
	this[1] = a[1] - b[1]
	this[2] = a[2] - b[2]
	this[3] = a[3] - b[3]
	
	return this
}

// subtracts scalar from this
func ( this *Vec4 ) SSub( s float32 ) *Vec4{
	
	this[0] -= s
	this[1] -= s
	this[2] -= s
	this[3] -= s
	
	return this
}

// mutiples this by other
func ( this *Vec4 ) Mul( other *Vec4 ) *Vec4{
	
	this[0] *= other[0]
	this[1] *= other[1]
	this[2] *= other[2]
	this[3] *= other[3]
	
	return this
}

// mutiples a and b saves in this
func ( this *Vec4 ) VMul( a, b *Vec4 ) *Vec4{
	
	this[0] = a[0] * b[0]
	this[1] = a[1] * b[1]
	this[2] = a[2] * b[2]
	this[3] = a[3] * b[3]
	
	return this
}

// mutiples this by scalar
func ( this *Vec4 ) SMul( s float32 ) *Vec4{
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	
	return this
}

// divides this by other
func ( this *Vec4 ) Div( other *Vec4 ) *Vec4{
	x, y, z, w := other[0], other[1], other[2], other[3]
	
	if x != 0 { x = 1 / x }
	if y != 0 { y = 1 / y }
	if z != 0 { z = 1 / z }
	if w != 0 { w = 1 / w }
	
	this[0] *= x
	this[1] *= y
	this[2] *= z
	this[3] *= w
	
	return this
}

// divides a and b saves in this
func ( this *Vec4 ) VDiv( a, b *Vec4 ) *Vec4{
	x, y, z, w := b[0], b[1], b[2], b[3]
	
	if x != 0 { x = 1 / x }
	if y != 0 { y = 1 / y }
	if z != 0 { z = 1 / z }
	if w != 0 { w = 1 / w }
	
	this[0] = a[0] * x
	this[1] = a[1] * y
	this[2] = a[2] * z
	this[3] = a[3] * w
	
	return this
}

// divides this by scalar
func ( this *Vec4 ) SDiv( s float32 ) *Vec4{
	if s != 0 { s = 1 / s }
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	this[3] *= s
	
	return this
}

// normalize vector, makes length of 1
func ( this *Vec4 ) Normalize() *Vec4{
	l := this[0] * this[0] + this[1] * this[1] + this[2] * this[2] + this[3] * this[3]

	if l == 0 { return this }

	l = 1 / float32( math.Sqrt( float64( l ) ) )
	this[0] *= l
	this[1] *= l
	this[2] *= l
	this[3] *= l
	
	return this
}

// returns length of this
func ( this *Vec4 ) Length() float32{
	l := this[0] * this[0] + this[1] * this[1] + this[2] * this[2] + this[3] * this[3]

	if l == 0 { return 0 }
	
	return float32( math.Sqrt( float64( l ) ) )
}

// returns length squared of this
func ( this *Vec4 ) LengthSq() float32{
	
	return this[0] * this[0] + this[1] * this[1] + this[2] * this[2] + this[3] * this[3]
}

// sets length of vector, if length is zero returns zero vector
func ( this *Vec4 ) SetLength( length float32 ) *Vec4{
	l := this[0] * this[0] + this[1] * this[1] + this[2] * this[2] + this[3] * this[3]

	if l == 0 { return this }

	l = 1 / float32( math.Sqrt( float64( l ) ) )
	this[0] *= length * l
	this[1] *= length * l
	this[2] *= length * l
	this[3] *= length * l
	
	return this
}

// returns dot product of this and other
func ( this *Vec4 ) Dot( other *Vec4 ) float32{
	
	return this[0] * other[0] + this[1] * other[1] + this[2] * other[2] + this[3] * other[3]
}

// returns dot product of a and b
func ( this *Vec4 ) VDot( a, b *Vec4 ) float32{
	
	return a[0] * b[0] + a[1] * b[1] + a[2] * b[2] + a[3] * b[3]
}

// returns inverse of this
func ( this *Vec4 ) Inverse() *Vec4{
	
	this[0] *= -1
	this[1] *= -1
	this[2] *= -1
	this[3] *= -1
	
	return this
}

// saves inverse of other in this
func ( this *Vec4 ) VInverse( other *Vec4 ) *Vec4{
	
	this[0] = -other[0]
	this[1] = -other[1]
	this[2] = -other[2]
	this[3] = -other[3]
	
	return this
}

// lerps this and other by x
func ( this *Vec4 ) Lerp( other *Vec4, x float32 ) *Vec4{

	this[0] += ( other[0] - this[0] ) * x
	this[1] += ( other[1] - this[1] ) * x
	this[2] += ( other[2] - this[2] ) * x
	this[3] += ( other[3] - this[3] ) * x

	return this
}

// lerps a and b by x, saves in this
func ( this *Vec4 ) VLerp( a, b *Vec4, x float32 ) *Vec4{
	
	this[0] = a[0] + ( b[0] - a[0] ) * x
	this[1] = a[1] + ( b[1] - a[1] ) * x
	this[2] = a[2] + ( b[2] - a[2] ) * x
	this[3] = a[3] + ( b[3] - a[3] ) * x

	return this
}

// returns distance to other
func ( this *Vec4 ) DistanceTo( other *Vec4 ) float32{
	x := this[0] - other[0]
	y := this[1] - other[1]
	z := this[2] - other[2]
	w := this[3] - other[3]
	l := x * x + y * y + z * z + w * w
	
	if l == 0 { return 0 }
	
	return float32( math.Sqrt( float64( l ) ) )
}

// returns distance squared to other
func ( this *Vec4 ) DistanceToSq( other *Vec4 ) float32{
	x := this[0] - other[0]
	y := this[1] - other[1]
	z := this[2] - other[2]
	w := this[3] - other[3]
	
	return x * x + y * y + z * z + w * w
}

// sets this values from  min values of this and other
func ( this *Vec4 ) Min( other *Vec4 ) *Vec4{
	
	if this[0] > other[0] { this[0] = other[0] }
	if this[1] > other[1] { this[1] = other[1] }
	if this[2] > other[2] { this[2] = other[2] }
	if this[3] > other[3] { this[3] = other[3] }
	
	return this
}

// sets this values from max values of this and other
func ( this *Vec4 ) Max( other *Vec4 ) *Vec4{
	
	if this[0] < other[0] { this[0] = other[0] }
	if this[1] < other[1] { this[1] = other[1] }
	if this[2] < other[2] { this[2] = other[2] }
	if this[3] < other[3] { this[3] = other[3] }
	
	return this
}

// clamps this between min and max
func ( this *Vec4 ) Clamp( min, max *Vec4 ) *Vec4{
	
	if this[0] < min[0] { this[0] = min[0] }
	if this[1] < min[1] { this[1] = min[1] }
	if this[2] < min[2] { this[2] = min[2] }
	if this[3] < min[3] { this[3] = min[3] }
	
	if this[0] > max[0] { this[0] = max[0] }
	if this[1] > max[1] { this[1] = max[1] }
	if this[2] > max[2] { this[2] = max[2] }
	if this[3] > max[3] { this[3] = max[3] }
	
	return this
}

// clamps each element between 0 and 1
func ( this *Vec4 ) Clamp01() *Vec4{
	
	this[0] = Clamp01( this[0] )
	this[1] = Clamp01( this[1] )
	this[2] = Clamp01( this[2] )
	this[3] = Clamp01( this[3] )
	
	return this
}

// sets values from Vec2
func ( this *Vec4 ) FromVec2( v *Vec2 ) *Vec4{
	
	this[0] = v[0]
	this[1] = v[1]
	this[2] = 0
	this[3] = 1
	
	return this
}

// sets values from Vec3
func ( this *Vec4 ) FromVec3( v *Vec3 ) *Vec4{
	
	this[0] = v[0]
	this[1] = v[1]
	this[2] = v[2]
	this[3] = 1
	
	return this
}

// checks if this equals other
func ( this *Vec4 ) Equals( other *Vec4 ) bool{
	
	if this[0] != other[0] { return false }
	if this[1] != other[1] { return false }
	if this[2] != other[2] { return false }
	if this[3] != other[3] { return false }
	
	return true
}

// returns this as string type
func ( this *Vec4 ) String() string{
	
	return fmt.Sprintf("Vec4[ %f, %f, %f, %f ]", this[0], this[1], this[2], this[3] )
}