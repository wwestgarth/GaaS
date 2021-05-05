package geometry

type Geometry interface {
	Validate() (err error)
}
