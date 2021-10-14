package vector

import "math"

// All vectors will be immutable and any changes will dervice a new instance
// As such all receivers shall not be pointers

type Vector struct {
	x float64
	y float64
	z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func NewZeroVector() Vector {
	return Vector{}
}

func (v Vector) Unpack() [3]float64 {
	return [3]float64{v.x, v.y, v.z}
}

func (v Vector) Len() float64 {
	return math.Sqrt(Dot(v, v))
}

func Dot(v, u Vector) float64 {
	return (v.x * u.x) + (v.y * u.y) + (v.z * u.z)
}

func Sub(v, u Vector) Vector {
	return Vector{
		x: v.x - u.x,
		y: u.y - v.y,
		z: u.z - u.z,
	}
}
