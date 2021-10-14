package topology

type Geometry interface {

	// ID return the ID of the geometry
	ID() (id string)

	// Validate return whether the geometry is valid
	Validate() (err error)
}
