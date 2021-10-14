package curve

import (
	"github.com/google/uuid"
	"github.com/wwestgarth/remote-geometry/vector"
)

type Line struct {
	id        uuid.UUID
	point     vector.Vector
	direction vector.Vector
	//axis Vector
}

func NewLine(point, direction vector.Vector) *Line {
	return &Line{
		id:        uuid.New(),
		point:     point,
		direction: direction,
	}
}

func (l *Line) ID() string {
	return l.id.String()
}

func (l *Line) Validate() error {
	return nil // TODO
}
