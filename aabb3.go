package mathf

import "fmt"

// float32 Array representing an 3D axis aligned bounding box
type AABB3 struct {
	Min, Max *Vec3
}

// returns new AABB3
func NewAABB3(min, max *Vec3) *AABB3 {
	this := new(AABB3)

	this.Min = NewVec3(Inf, Inf, Inf)
	this.Max = NewVec3(-Inf, -Inf, -Inf)

	return this
}

// returns a copy of this
func (this *AABB3) Clone() *AABB3 {

	return new(AABB3).Copy(this)
}

// copies other
func (this *AABB3) Copy(other *AABB3) *AABB3 {

	this.Min.Copy(other.Min)
	this.Max.Copy(other.Max)

	return this
}

// sets this from values
func (this *AABB3) Set(min, max *Vec3) *AABB3 {

	this.Min.Copy(min)
	this.Max.Copy(max)

	return this
}

// expands this by point
func (this *AABB3) ExpandPoint(v *Vec3) *AABB3 {

	this.Min.Min(v)
	this.Max.Max(v)

	return this
}

// expands this by vector
func (this *AABB3) ExpandVec(v *Vec3) *AABB3 {

	this.Min.Sub(v)
	this.Max.Add(v)

	return this
}

// expands this by scalar
func (this *AABB3) ExpandScalar(s float32) *AABB3 {

	this.Min.SSub(s)
	this.Max.SAdd(s)

	return this
}

// merges this and other
func (this *AABB3) Union(other *AABB3) *AABB3 {

	this.Min.Min(other.Min)
	this.Max.Max(other.Max)

	return this
}

// returns this as empty
func (this *AABB3) Empty() *AABB3 {

	this.Min.Set(Inf, Inf, Inf)
	this.Max.Set(-Inf, -Inf, -Inf)

	return this
}

// checks if AABB contains point
func (this *AABB3) Contains(v *Vec3) bool {

	if v[0] < this.Min[0] {
		return false
	}
	if v[1] < this.Min[1] {
		return false
	}
	if v[2] < this.Min[2] {
		return false
	}

	if v[0] > this.Max[0] {
		return false
	}
	if v[1] > this.Max[1] {
		return false
	}
	if v[2] > this.Max[2] {
		return false
	}

	return true
}

// checks if this and other intersect
func (this *AABB3) Intersects(other *AABB3) bool {

	if other.Max[0] < this.Min[0] {
		return false
	}
	if other.Max[1] < this.Min[1] {
		return false
	}
	if other.Max[2] < this.Min[2] {
		return false
	}

	if other.Min[0] > this.Max[0] {
		return false
	}
	if other.Min[1] > this.Max[1] {
		return false
	}
	if other.Min[2] > this.Max[2] {
		return false
	}

	return true
}

// sets min and max from array of points
func (this *AABB3) FromPoints(array []*Vec3) *AABB3 {
	l := len(array)

	if l == 0 {
		return this.Empty()
	}

	v := array[0]
	minx := v[0]
	miny := v[1]
	minz := v[2]
	maxx := v[0]
	maxy := v[1]
	maxz := v[2]

	for i := 1; i < l; i++ {
		v = array[i]

		if minx > v[0] {
			minx = v[0]
		}
		if miny > v[1] {
			miny = v[1]
		}
		if minz > v[2] {
			minz = v[2]
		}
		if maxx < v[0] {
			maxx = v[0]
		}
		if maxy < v[1] {
			maxy = v[1]
		}
		if maxz < v[2] {
			maxz = v[2]
		}
	}

	this.Min[0] = minx
	this.Min[1] = miny
	this.Min[2] = minz
	this.Max[0] = maxx
	this.Max[1] = maxy
	this.Max[2] = maxz

	return this
}

// sets min and max from center and size
func (this *AABB3) FromCenterSize(center, size *Vec3) *AABB3 {
	hx := size[0] * 0.5
	hy := size[1] * 0.5

	this.Min[0] = center[0] - hx
	this.Min[1] = center[1] - hy

	this.Max[0] = center[0] + hx
	this.Max[1] = center[1] + hy

	return this
}

// checks if this equals other
func (this *AABB3) Equals(other *AABB3) bool {

	this.Min.Equals(other.Min)
	this.Max.Equals(other.Max)

	return true
}

// returns this as string type
func (this *AABB3) String() string {

	return fmt.Sprintf("AABB3[ Min: %s, Max: %s ]", this.Min, this.Max)
}
