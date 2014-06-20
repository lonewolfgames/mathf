package mathf

import "fmt"

// float32 Array representing an 2D axis aligned bounding box
type AABB2 struct {
	Min, Max *Vec2
}

// returns new AABB2
func NewAABB2() *AABB2 {
	this := new(AABB2)

	this.Min = NewVec2(Inf, Inf)
	this.Max = NewVec2(-Inf, -Inf)

	return this
}

// returns a copy of this
func (this *AABB2) Clone() *AABB2 {

	return new(AABB2).Copy(this)
}

// copies other
func (this *AABB2) Copy(other *AABB2) *AABB2 {

	this.Min.Copy(other.Min)
	this.Max.Copy(other.Max)

	return this
}

// sets this from values
func (this *AABB2) Set(min, max *Vec2) *AABB2 {

	this.Min.Copy(min)
	this.Max.Copy(max)

	return this
}

// expands this by point
func (this *AABB2) ExpandPoint(v *Vec2) *AABB2 {

	this.Min.Min(v)
	this.Max.Max(v)

	return this
}

// expands this by vector
func (this *AABB2) ExpandVec(v *Vec2) *AABB2 {

	this.Min.Sub(v)
	this.Max.Add(v)

	return this
}

// expands this by scalar
func (this *AABB2) ExpandScalar(s float32) *AABB2 {

	this.Min.SSub(s)
	this.Max.SAdd(s)

	return this
}

// merges this and other
func (this *AABB2) Union(other *AABB2) *AABB2 {

	this.Min.Min(other.Min)
	this.Max.Max(other.Max)

	return this
}

// returns this as empty
func (this *AABB2) Empty() *AABB2 {

	this.Min.Set(Inf, Inf)
	this.Max.Set(-Inf, -Inf)

	return this
}

// checks if AABB contains point
func (this *AABB2) Contains(v *Vec2) bool {

	if v[0] < this.Min[0] {
		return false
	}
	if v[1] < this.Min[1] {
		return false
	}

	if v[0] > this.Max[0] {
		return false
	}
	if v[1] > this.Max[1] {
		return false
	}

	return true
}

// checks if this and other intersect
func (this *AABB2) Intersects(other *AABB2) bool {

	if other.Max[0] < this.Min[0] {
		return false
	}
	if other.Max[1] < this.Min[1] {
		return false
	}

	if other.Min[0] > this.Max[0] {
		return false
	}
	if other.Min[1] > this.Max[1] {
		return false
	}

	return true
}

// sets min and max from array of points
func (this *AABB2) FromPoints(array []*Vec2) *AABB2 {
	l := len(array)

	if l == 0 {
		return this.Empty()
	}

	v := array[0]
	minx := v[0]
	miny := v[1]
	maxx := v[0]
	maxy := v[1]

	for i := 1; i < l; i++ {
		v = array[i]

		if minx > v[0] {
			minx = v[0]
		}
		if miny > v[1] {
			miny = v[1]
		}
		if maxx < v[0] {
			maxx = v[0]
		}
		if maxy < v[1] {
			maxy = v[1]
		}
	}

	this.Min[0] = minx
	this.Min[1] = miny
	this.Max[0] = maxx
	this.Max[1] = maxy

	return this
}

// sets min and max from center and size
func (this *AABB2) FromCenterSize(center, size *Vec2) *AABB2 {
	hx := size[0] * 0.5
	hy := size[1] * 0.5

	this.Min[0] = center[0] - hx
	this.Min[1] = center[1] - hy

	this.Max[0] = center[0] + hx
	this.Max[1] = center[1] + hy

	return this
}

// checks if this equals other
func (this *AABB2) Equals(other *AABB2) bool {

	this.Min.Equals(other.Min)
	this.Max.Equals(other.Max)

	return true
}

// returns this as string type
func (this *AABB2) String() string {

	return fmt.Sprintf("AABB2[ Min: %s, Max: %s ]", this.Min, this.Max)
}
