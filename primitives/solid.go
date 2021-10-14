package primitives

import (
	"github.com/wwestgarth/remote-geometry/geometry/surface"
	"github.com/wwestgarth/remote-geometry/topology"
	"github.com/wwestgarth/remote-geometry/vector"
)

func CreateSphere(radius float64, centre vector.Vector) *topology.Body {

	b := topology.NewBody()
	g := surface.NewSphere(radius, centre.Unpack())
	b.FaceAny().Attach(g)
	return b
}
