package mathf

import (
	"math"
	"fmt"
)

// float32 Array representing a 3D Vector
type Vec3 [3]float32

// returns new Vec3
func NewVec3( x, y, z float32 ) *Vec3{
	this := new( Vec3 )
	
	this[0], this[1], this[2] = x, y, z
	
	return this
}

// returns a copy of this
func ( this *Vec3 ) Clone() *Vec3{
	
	return new( Vec3 ).Copy( this )
}

// copies other
func ( this *Vec3 ) Copy( other *Vec3 ) *Vec3{
	
	this[0], this[1], this[2] = other[0], other[1], other[2]
	
	return this
}

// sets this from values
func ( this *Vec3 ) Set( x, y, z float32 ) *Vec3{
	
	this[0], this[1], this[2] = x, y, z
	
	return this
}

// adds other to this
func ( this *Vec3 ) Add( other *Vec3 ) *Vec3{
	
	this[0] += other[0]
	this[1] += other[1]
	this[2] += other[2]
	
	return this
}

// adds a and b saves in this
func ( this *Vec3 ) VAdd( a, b *Vec3 ) *Vec3{
	
	this[0] = a[0] + b[0]
	this[1] = a[1] + b[1]
	this[2] = a[2] + b[2]
	
	return this
}

// adds scalar to this
func ( this *Vec3 ) SAdd( s float32 ) *Vec3{
	
	this[0] += s
	this[1] += s
	this[2] += s
	
	return this
}

// subtracts other from this
func ( this *Vec3 ) Sub( other *Vec3 ) *Vec3{
	
	this[0] -= other[0]
	this[1] -= other[1]
	this[2] -= other[2]
	
	return this
}

// subtracts a and b saves in this
func ( this *Vec3 ) VSub( a, b *Vec3 ) *Vec3{
	
	this[0] = a[0] - b[0]
	this[1] = a[1] - b[1]
	this[2] = a[2] - b[2]
	
	return this
}

// subtracts scalar from this
func ( this *Vec3 ) SSub( s float32 ) *Vec3{
	
	this[0] -= s
	this[1] -= s
	this[2] -= s
	
	return this
}

// mutiples this by other
func ( this *Vec3 ) Mul( other *Vec3 ) *Vec3{
	
	this[0] *= other[0]
	this[1] *= other[1]
	this[2] *= other[2]
	
	return this
}

// mutiples a and b saves in this
func ( this *Vec3 ) VMul( a, b *Vec3 ) *Vec3{
	
	this[0] = a[0] * b[0]
	this[1] = a[1] * b[1]
	this[2] = a[2] * b[2]
	
	return this
}

// mutiples this by scalar
func ( this *Vec3 ) SMul( s float32 ) *Vec3{
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	
	return this
}

// divides this by other
func ( this *Vec3 ) Div( other *Vec3 ) *Vec3{
	x, y, z := other[0], other[1], other[2]
	
	if x != 0 { x = 1 / x }
	if y != 0 { y = 1 / y }
	if z != 0 { z = 1 / z }
	
	this[0] *= x
	this[1] *= y
	this[2] *= z
	
	return this
}

// divides a and b saves in this
func ( this *Vec3 ) VDiv( a, b *Vec3 ) *Vec3{
	x, y, z := b[0], b[1], b[2]
	
	if x != 0 { x = 1 / x }
	if y != 0 { y = 1 / y }
	if z != 0 { z = 1 / z }
	
	this[0] = a[0] * x
	this[1] = a[1] * y
	this[2] = a[2] * z
	
	return this
}

// divides this by scalar
func ( this *Vec3 ) SDiv( s float32 ) *Vec3{
	if s != 0 { s = 1 / s }
	
	this[0] *= s
	this[1] *= s
	this[2] *= s
	
	return this
}

// normalize vector, makes length of 1
func ( this *Vec3 ) Normalize() *Vec3{
	l := this[0] * this[0] + this[1] * this[1] + this[2] * this[2]

	if l == 0 { return this }

	l = 1 / float32( math.Sqrt( float64( l ) ) )
	this[0] *= l
	this[1] *= l
	this[2] *= l
	
	return this
}

// returns length of this
func ( this *Vec3 ) Length() float32{
	l := this[0] * this[0] + this[1] * this[1] + this[2] * this[2]

	if l == 0 { return 0 }
	
	return float32( math.Sqrt( float64( l ) ) )
}

// returns length squared of this
func ( this *Vec3 ) LengthSq() float32{
	
	return this[0] * this[0] + this[1] * this[1] + this[2] * this[2]
}

// sets length of vector, if length is zero returns zero vector
func ( this *Vec3 ) SetLength( length float32 ) *Vec3{
	l := this[0] * this[0] + this[1] * this[1] + this[2] * this[2]

	if l == 0 { return this }

	l = 1 / float32( math.Sqrt( float64( l ) ) )
	this[0] *= length * l
	this[1] *= length * l
	this[2] *= length * l
	
	return this
}

// returns dot product of this and other
func ( this *Vec3 ) Dot( other *Vec3 ) float32{
	
	return this[0] * other[0] + this[1] * other[1] + this[2] * other[2]
}

// returns dot product of a and b
func ( this *Vec3 ) VDot( a, b *Vec3 ) float32{
	
	return a[0] * b[0] + a[1] * b[1] + a[2] * b[2]
}

// returns cross product of this and other
func ( this *Vec3 ) Cross( other *Vec3 ) *Vec3{
	
	this[0] = this[1] * other[2] - this[2] * other[1]
	this[1] = this[2] * other[0] - this[0] * other[2]
	this[2] = this[0] * other[1] - this[1] * other[0]
	
	return this
}

// returns cross product of a and b saves it in this
func ( this *Vec3 ) VCross( a, b *Vec3 ) *Vec3{
	
	this[0] = a[1] * b[2] - a[2] * b[1]
	this[1] = a[2] * b[0] - a[0] * b[2]
	this[2] = a[0] * b[1] - a[1] * b[0]
	
	return this
}

// returns inverse of this
func ( this *Vec3 ) Inverse() *Vec3{
	
	this[0] *= -1
	this[1] *= -1
	this[2] *= -1
	
	return this
}

// saves inverse of other in this
func ( this *Vec3 ) VInverse( other *Vec3 ) *Vec3{
	
	this[0] = -other[0]
	this[1] = -other[1]
	this[2] = -other[2]
	
	return this
}

// lerps this and other by x
func ( this *Vec3 ) Lerp( other *Vec3, x float32 ) *Vec3{

	this[0] += ( other[0] - this[0] ) * x
	this[1] += ( other[1] - this[1] ) * x
	this[2] += ( other[2] - this[2] ) * x

	return this
}

// lerps a and b by x, saves in this
func ( this *Vec3 ) VLerp( a, b *Vec3, x float32 ) *Vec3{
	
	this[0] = a[0] + ( b[0] - a[0] ) * x
	this[1] = a[1] + ( b[1] - a[1] ) * x
	this[2] = a[2] + ( b[2] - a[2] ) * x

	return this
}

// returns distance to other
func ( this *Vec3 ) DistanceTo( other *Vec3 ) float32{
	x := this[0] - other[0]
	y := this[1] - other[1]
	z := this[2] - other[2]
	l := x * x + y * y + z * z
	
	if l == 0 { return 0 }
	
	return float32( math.Sqrt( float64( l ) ) )
}

// returns distance squared to other
func ( this *Vec3 ) DistanceToSq( other *Vec3 ) float32{
	x := this[0] - other[0]
	y := this[1] - other[1]
	z := this[2] - other[2]
	
	return x * x + y * y + z * z
}

// sets this values from  min values of this and other
func ( this *Vec3 ) Min( other *Vec3 ) *Vec3{
	
	if this[0] > other[0] { this[0] = other[0] }
	if this[1] > other[1] { this[1] = other[1] }
	if this[2] > other[2] { this[2] = other[2] }
	
	return this
}

// sets this values from max values of this and other
func ( this *Vec3 ) Max( other *Vec3 ) *Vec3{
	
	if this[0] < other[0] { this[0] = other[0] }
	if this[1] < other[1] { this[1] = other[1] }
	if this[2] < other[2] { this[2] = other[2] }
	
	return this
}

// clamps this between min and max
func ( this *Vec3 ) Clamp( min, max *Vec3 ) *Vec3{
	
	if this[0] < min[0] { this[0] = min[0] }
	if this[1] < min[1] { this[1] = min[1] }
	if this[2] < min[2] { this[2] = min[2] }
	
	if this[0] > max[0] { this[0] = max[0] }
	if this[1] > max[1] { this[1] = max[1] }
	if this[2] > max[2] { this[2] = max[2] }
	
	return this
}

// clamps each element between 0 and 1
func ( this *Vec3 ) Clamp01() *Vec3{
	
	this[0] = Clamp01( this[0] )
	this[1] = Clamp01( this[1] )
	this[2] = Clamp01( this[2] )
	
	return this
}

// checks if this equals other
func ( this *Vec3 ) Equals( other *Vec3 ) bool{
	
	if this[0] != other[0] { return false }
	if this[1] != other[1] { return false }
	if this[2] != other[2] { return false }
	
	return true
}

// returns this as string type
func ( this *Vec3 ) String() string{
	
	return fmt.Sprintf("Vec3[ %f, %f, %f ]", this[0], this[1], this[2] )
}