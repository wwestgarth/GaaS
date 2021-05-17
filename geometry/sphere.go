package geometry

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidRadius = errors.New("invalid radius")
)

type Sphere struct {
	id     string
	centre Vector
	radius float64
}

// NewSphere create a new Geometry that represents a sphere
func NewSphere(radius float64, centre [3]float64) Geometry {

	return &Sphere{
		id:     uuid.New().String(),
		radius: radius,
		centre: NewVector(centre[0], centre[1], centre[2]),
	}
}

// ID return the ID of the geometry
func (s *Sphere) ID() (id string) {
	return s.id
}

// Validate return whether the geometry is valid
func (s *Sphere) Validate() (err error) {

	if distZero(s.radius) {
		return ErrInvalidRadius
	}

	if s.radius <= 0.0 {
		return ErrInvalidRadius
	}

	// TODO sizebox check
	return
}
