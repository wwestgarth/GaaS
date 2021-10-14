package curve

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wwestgarth/remote-geometry/vector"
)

var (
	ErrInvalidRadius = errors.New("invalid radius")
)

type Circle struct {
	id     uuid.UUID
	radius float64
	centre vector.Vector
	normal vector.Vector
	//axis Vector
}

func NewCircle(radius float64, centre, normal vector.Vector) *Circle {
	return &Circle{
		id:     uuid.New(),
		radius: radius,
		centre: centre,
		normal: normal,
	}
}

func (c *Circle) ID() string {
	return c.id.String()
}

func (c *Circle) Validate() error {

	if vector.DistZero(c.radius) {
		return ErrInvalidRadius
	}

	if c.radius <= 0.0 {
		return ErrInvalidRadius
	}

	return nil // TODO
}
