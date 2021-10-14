package surface

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wwestgarth/remote-geometry/vector"
)

var (
	ErrInvalidRadius = errors.New("invalid radius")
)

type Sphere struct {
	id     uuid.UUID
	centre vector.Vector
	radius float64
}

// NewSphere create a new Geometry that represents a sphere
func NewSphere(radius float64, centre [3]float64) *Sphere {

	return &Sphere{
		id:     uuid.New(),
		radius: radius,
		centre: vector.NewVector(centre[0], centre[1], centre[2]),
	}
}

// ID return the ID of the geometry
func (s *Sphere) ID() string {
	return s.id.String()
}

// Validate return whether the geometry is valid
func (s *Sphere) Validate() (err error) {

	if vector.DistZero(s.radius) {
		return ErrInvalidRadius
	}

	if s.radius <= 0.0 {
		return ErrInvalidRadius
	}

	// TODO sizebox check
	return
}
