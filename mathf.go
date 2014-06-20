package mathf

import (
	"math"
	"math/rand"
)

const (
	Epsilon = float32(0.000001)
	Pi      = float32(math.Pi)
	TWO_PI  = Pi * 2
	HALF_PI = Pi * 0.5
	TO_RADS = Pi / 180
	TO_DEGS = 180 / Pi

	Inf = math.MaxFloat32

	Pi_1_16  = TWO_PI * 1 / 16
	Pi_3_16  = TWO_PI * 3 / 16
	Pi_5_16  = TWO_PI * 5 / 16
	Pi_7_16  = TWO_PI * 7 / 16
	Pi_9_16  = TWO_PI * 9 / 16
	Pi_11_16 = TWO_PI * 11 / 16
	Pi_13_16 = TWO_PI * 13 / 16
	Pi_15_16 = TWO_PI * 15 / 16

	RIGHT      = "right"
	UP_RIGHT   = "up_right"
	UP         = "up"
	UP_LEFT    = "up_left"
	LEFT       = "left"
	DOWN_LEFT  = "down_left"
	DOWN       = "down"
	DOWN_RIGHT = "down_right"
)

// returns absolute value of x
func Abs(x float32) float32 {
	if x < 0 {
		return -x
	}

	return x
}

// checks if a and b are equal
func Equals(a, b float32) bool {

	return Abs(a-b) < Epsilon
}

// returns x as a standard radian 0 <= x < PI
func StandardRadian(x float32) float32 {

	return float32(math.Mod(float64(x), float64(TWO_PI)))
}

// returns x as a standard angle 0 <= x < 360
func StandardAngle(x float32) float32 {

	return float32(math.Mod(float64(x), 360))
}

// returns sign of x
func Sign(x float32) float32 {
	if x < 0 {
		return -1
	}

	return 1
}

// clamp x between min and max
func Clamp(x, min, max float32) float32 {

	if x < min {
		return min
	}
	if x > max {
		return max
	}

	return x
}

// clamp x between 0 and 1
func Clamp01(x float32) float32 {

	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}

	return x
}

// truncate x
func Truncate(x, n float32) float32 {
	p := float32(math.Pow(10, float64(n)))
	num := x * p

	if num < 0 {
		return float32(math.Ceil(float64(num))) / p
	}

	return float32(math.Floor(float64(num))) / p
}

// lerps between a and b by x
func Lerp(x, a, b float32) float32 {

	return a + (b-a)*x
}

// lerps a and b by x, insures standrad radian
func LerpRadian(x, a, b float32) float32 {

	return StandardRadian(a + (b-a)*x)
}

// lerps a and b by x, insures standrad angle
func LerpAngle(x, a, b float32) float32 {

	return StandardAngle(a + (b-a)*x)
}

// smooth step
func SmoothStep(x, min, max float32) float32 {
	if x <= min {
		return 0
	}
	if x >= max {
		return 1
	}

	x = (x - min) / (max - min)

	return x * x * (3 - 2*x)
}

// smoother step
func SmootherStep(x, min, max float32) float32 {
	if x <= min {
		return 0
	}
	if x >= max {
		return 1
	}

	x = (x - min) / (max - min)

	return x * x * x * (x*(x*6-15) + 10)
}

// returns x in radians
func DegsToRads(x float32) float32 {

	return StandardRadian(x * TO_RADS)
}

// returns x in degrees
func RadsToDegs(x float32) float32 {

	return StandardAngle(x * TO_DEGS)
}

// ping pongs x
func PingPong(x, length float32) float32 {

	return length - float32(math.Abs(math.Mod(float64(x), float64((2*length)-length))))
}

// returns random int
func RandInt(min, max int) int {

	return min + rand.Intn(max-min)
}

// returns random int32
func RandInt32(min, max int32) int32 {

	return min + rand.Int31n(max-min)
}

// returns random int64
func RandInt64(min, max int64) int64 {

	return min + rand.Int63n(max-min)
}

// returns random float32
func RandFloat(min, max float32) float32 {

	return min + (rand.Float32() * (max - min))
}

// checks if x is a power of two
func IsPOTInt(x int) bool {

	if (x & (x - 1)) == 0 {
		return true
	}

	return false
}

// checks if x is a power of two
func IsPOTInt32(x int32) bool {

	if (x & (x - 1)) == 0 {
		return true
	}

	return false
}

// checks if x is a power of two
func IsPOTInt64(x int64) bool {

	if (x & (x - 1)) == 0 {
		return true
	}

	return false
}

// returns x to its next power of two
func ToPOTInt(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++

	return x
}

// returns x to its next power of two
func ToPOTInt32(x int32) int32 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++

	return x
}

// returns x to its next power of two
func ToPOTInt64(x int64) int64 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++

	return x
}

// returns string name of angle
func DirectionAngle(a float32) string {
	a = StandardAngle(a)

	if a > 337.5 && a < 22.5 {
		return RIGHT
	}
	if a > 22.5 && a < 67.5 {
		return UP_RIGHT
	}
	if a > 67.5 && a < 112.5 {
		return UP
	}
	if a > 112.5 && a < 157.5 {
		return UP_LEFT
	}
	if a > 157.5 && a < 202.5 {
		return LEFT
	}
	if a > 202.5 && a < 247.5 {
		return DOWN_LEFT
	}
	if a > 247.5 && a < 292.5 {
		return DOWN
	}
	if a > 292.5 && a < 337.5 {
		return DOWN_RIGHT
	}

	return RIGHT
}

// returns string name of radian
func DirectionRads(a float32) string {
	a = StandardRadian(a)

	if a > Pi_15_16 && a < Pi_1_16 {
		return RIGHT
	}
	if a > Pi_1_16 && a < Pi_3_16 {
		return UP_RIGHT
	}
	if a > Pi_3_16 && a < Pi_5_16 {
		return UP
	}
	if a > Pi_5_16 && a < Pi_7_16 {
		return UP_LEFT
	}
	if a > Pi_7_16 && a < Pi_9_16 {
		return LEFT
	}
	if a > Pi_9_16 && a < Pi_11_16 {
		return DOWN_LEFT
	}
	if a > Pi_11_16 && a < Pi_13_16 {
		return DOWN
	}
	if a > Pi_13_16 && a < Pi_15_16 {
		return DOWN_RIGHT
	}

	return RIGHT
}

// returns smallest number in arguments
func Min(nums ...float32) float32 {
	len := len(nums)
	if len == 0 {
		return 0
	}

	min := nums[0]
	a, b := min, min

	for i := 1; i < len; i++ {
		a = nums[i]

		if i+1 > len-1 {
			b = nums[0]
		} else {
			b = nums[i+1]
		}
		if a < b {
			min = a
		}
	}

	return min
}

// returns largest number in arguments
func Max(nums ...float32) float32 {
	len := len(nums)
	if len == 0 {
		return 0
	}

	min := nums[0]
	a, b := min, min

	for i := 1; i < len; i++ {
		a = nums[i]

		if i+1 > len-1 {
			b = nums[0]
		} else {
			b = nums[i+1]
		}
		if a > b {
			min = a
		}
	}

	return min
}
