package geometry

import "errors"

var (
	ErrNotImplemented = errors.New("not implemented")
)

type Geometry interface {

	// ID return the ID of the geometry
	ID() (id string)

	// Validate return whether the geometry is valid
	Validate() (err error)
}
