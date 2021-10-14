package surface

import (
	"github.com/google/uuid"
	"github.com/wwestgarth/remote-geometry/vector"
)

type Plane struct {
	id     uuid.UUID
	normal vector.Vector
	point  vector.Vector
	//axis   Vector
}

func NewPlane(point, normal vector.Vector) *Plane {
	return &Plane{
		id:     uuid.New(),
		normal: normal,
		point:  point,
	}
}

func (p *Plane) ID() string {
	return p.id.String()
}

func (p *Plane) Validate() error {
	return nil // TODO
}
