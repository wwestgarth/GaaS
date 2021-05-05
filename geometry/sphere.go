package geometry

import (
	"errors"

	"github.com/go-gl/mathgl/mgl64"
)

var (
	ErrNegativeRadius = errors.New("negative radius")
)

type Sphere struct {
	centre mgl64.Vec3
	radius float64
}

func NewSphere(radius float64, centre [3]float64) Geometry {

	return &Sphere{
		radius: radius,
		centre: mgl64.Vec3(centre), //TODO look for vector package
	}
}

func (s *Sphere) Validate() (err error) {

	// TODO tolerances
	if s.radius <= 0.0 {
		return ErrNegativeRadius
	}

	// TODO sizebox check
	return
}
