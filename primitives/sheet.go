package primitives

import (
	"github.com/wwestgarth/remote-geometry/geometry/curve"
	"github.com/wwestgarth/remote-geometry/geometry/surface"
	"github.com/wwestgarth/remote-geometry/topology"
	"github.com/wwestgarth/remote-geometry/vector"
)

func CreateSheetCircle(radius float64, centre, normal vector.Vector) *topology.Body {

	b := topology.NewSheetBody()
	f := b.FaceAny()

	p := surface.NewPlane(centre, normal)
	f.Attach(p)

	c := curve.NewCircle(radius, centre, normal)
	f.EdgeAny().Attach(c)

	return b
}
