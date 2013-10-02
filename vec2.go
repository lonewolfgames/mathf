package mathf

import (
	"math"
	"fmt"
)

// float32 Array representing a 2D Vector
type Vec2 [2]float32

// returns new Vec2
func NewVec2( x, y float32 ) *Vec2{
	this := new( Vec2 )
	
	this[0], this[1] = x, y
	
	return this
}

// returns a copy of this
func ( this *Vec2 ) Clone() *Vec2{
	
	return new( Vec2 ).Copy( this )
}

// copies other
func ( this *Vec2 ) Copy( other *Vec2 ) *Vec2{
	
	this[0], this[1] = other[0], other[1]
	
	return this
}

// sets this from values
func ( this *Vec2 ) Set( x, y float32 ) *Vec2{
	
	this[0], this[1] = x, y
	
	return this
}

// adds other to this
func ( this *Vec2 ) Add( other *Vec2 ) *Vec2{
	
	this[0] += other[0]
	this[1] += other[1]
	
	return this
}

// adds a and b saves in this
func ( this *Vec2 ) VAdd( a, b *Vec2 ) *Vec2{
	
	this[0] = a[0] + b[0]
	this[1] = a[1] + b[1]
	
	return this
}

// adds scalar to this
func ( this *Vec2 ) SAdd( s float32 ) *Vec2{
	
	this[0] += s
	this[1] += s
	
	return this
}

// subtracts other from this
func ( this *Vec2 ) Sub( other *Vec2 ) *Vec2{
	
	this[0] -= other[0]
	this[1] -= other[1]
	
	return this
}

// subtracts a and b saves in this
func ( this *Vec2 ) VSub( a, b *Vec2 ) *Vec2{
	
	this[0] = a[0] - b[0]
	this[1] = a[1] - b[1]
	
	return this
}

// adds scalar to this
func ( this *Vec2 ) SSub( s float32 ) *Vec2{
	
	this[0] -= s
	this[1] -= s
	
	return this
}

// mutiples this by other
func ( this *Vec2 ) Mul( other *Vec2 ) *Vec2{
	
	this[0] *= other[0]
	this[1] *= other[1]
	
	return this
}

// mutiples a and b saves in this
func ( this *Vec2 ) VMul( a, b *Vec2 ) *Vec2{
	
	this[0] = a[0] * b[0]
	this[1] = a[1] * b[1]
	
	return this
}

// mutiples this by scalar
func ( this *Vec2 ) SMul( s float32 ) *Vec2{
	
	this[0] *= s
	this[1] *= s
	
	return this
}

// divides this by other
func ( this *Vec2 ) Div( other *Vec2 ) *Vec2{
	x, y := other[0], other[1]
	
	if x != 0 { x = 1 / x }
	if y != 0 { y = 1 / y }
	
	this[0] *= x
	this[1] *= y
	
	return this
}

// divides a and b saves in this
func ( this *Vec2 ) VDiv( a, b *Vec2 ) *Vec2{
	x, y := b[0], b[1]
	
	if x != 0 { x = 1 / x }
	if y != 0 { y = 1 / y }
	
	this[0] = a[0] * x
	this[1] = a[1] * y
	
	return this
}

// divides this by scalar
func ( this *Vec2 ) SDiv( s float32 ) *Vec2{
	if s != 0 { s = 1 / s }
	
	this[0] *= s
	this[1] *= s
	
	return this
}

// normalize vector, makes length of 1
func ( this *Vec2 ) Normalize() *Vec2{
	l := this[0] * this[0] + this[1] * this[1]

	if l == 0 { return this }

	l = 1 / float32( math.Sqrt( float64( l ) ) )
	this[0] *= l
	this[1] *= l
	
	return this
}

// returns length of this
func ( this *Vec2 ) Length() float32{
	l := this[0] * this[0] + this[1] * this[1]

	if l == 0 { return 0 }
	
	return float32( math.Sqrt( float64( l ) ) )
}

// returns length squared of this
func ( this *Vec2 ) LengthSq() float32{
	
	return this[0] * this[0] + this[1] * this[1]
}

// sets length of vector, if length is zero returns zero vector
func ( this *Vec2 ) SetLength( length float32 ) *Vec2{
	l := this[0] * this[0] + this[1] * this[1]

	if l == 0 { return this }

	l = 1 / float32( math.Sqrt( float64( l ) ) )
	this[0] *= length * l
	this[1] *= length * l
	
	return this
}

// returns dot product of this and other
func ( this *Vec2 ) Dot( other *Vec2 ) float32{
	
	return this[0] * other[0] + this[1] * other[1]
}

// returns dot product of a and b
func ( this *Vec2 ) VDot( a, b *Vec2 ) float32{
	
	return a[0] * b[0] + a[1] * b[1]
}

// returns cross product of this and other
func ( this *Vec2 ) Cross( other *Vec2 ) float32{
	
	return this[0] * other[1] - this[1] * other[0]
}

// returns cross product of a and b
func ( this *Vec2 ) VCross( a, b *Vec2 ) float32{
	
	return a[0] * b[1] - a[1] * b[0]
}

// returns inverse of this
func ( this *Vec2 ) Inverse() *Vec2{
	
	this[0] *= -1
	this[1] *= -1
	
	return this
}

// saves inverse of other in this
func ( this *Vec2 ) VInverse( other *Vec2 ) *Vec2{
	
	this[0] = -other[0]
	this[1] = -other[1]
	
	return this
}

// lerps this and other by x
func ( this *Vec2 ) Lerp( other *Vec2, x float32 ) *Vec2{

	this[0] += ( other[0] - this[0] ) * x
	this[1] += ( other[1] - this[1] ) * x

	return this
}

// lerps a and b by x, saves in this
func ( this *Vec2 ) VLerp( a, b *Vec2, x float32 ) *Vec2{
	
	this[0] = a[0] + ( b[0] - a[0] ) * x
	this[1] = a[1] + ( b[1] - a[1] ) * x

	return this
}

// returns distance to other
func ( this *Vec2 ) DistanceTo( other *Vec2 ) float32{
	x := this[0] - other[0]
	y := this[1] - other[1]
	l := x * x + y * y
	
	if l == 0 { return 0 }
	
	return float32( math.Sqrt( float64( l ) ) )
}

// returns distance squared to other
func ( this *Vec2 ) DistanceToSq( other *Vec2 ) float32{
	x := this[0] - other[0]
	y := this[1] - other[1]
	
	return x * x + y * y
}

// sets this values from  min values of this and other
func ( this *Vec2 ) Min( other *Vec2 ) *Vec2{
	
	if this[0] > other[0] { this[0] = other[0] }
	if this[1] > other[1] { this[1] = other[1] }
	
	return this
}

// sets this values from max values of this and other
func ( this *Vec2 ) Max( other *Vec2 ) *Vec2{
	
	if this[0] < other[0] { this[0] = other[0] }
	if this[1] < other[1] { this[1] = other[1] }
	
	return this
}

// clamps this between min and max
func ( this *Vec2 ) Clamp( min, max *Vec2 ) *Vec2{
	
	if this[0] < min[0] { this[0] = min[0] }
	if this[1] < min[1] { this[1] = min[1] }
	
	if this[0] > max[0] { this[0] = max[0] }
	if this[1] > max[1] { this[1] = max[1] }
	
	return this
}

// clamps each element between 0 and 1
func ( this *Vec2 ) Clamp01() *Vec2{
	
	this[0] = Clamp01( this[0] )
	this[1] = Clamp01( this[1] )
	
	return this
}

// checks if this equals other
func ( this *Vec2 ) Equals( other *Vec2 ) bool{
	
	if this[0] != other[0] { return false }
	if this[1] != other[1] { return false }
	
	return true
}

// returns this as string type
func ( this *Vec2 ) String() string{
	
	return fmt.Sprintf("Vec2[ %f, %f ]", this[0], this[1] )
}