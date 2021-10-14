package topology

import (
	"github.com/google/uuid"
)

type Face struct {
	id        uuid.UUID
	surface   Geometry
	edges     map[string]*Edge
	edgeOrder []string // an ordering for the edges in the map
}

func NewFace() *Face {
	return &Face{
		id:        uuid.New(),
		edges:     map[string]*Edge{},
		edgeOrder: []string{},
	}
}

func (f *Face) EdgeAny() *Edge {
	if len(f.edges) == 0 {
		return nil
	}

	return f.edges[f.edgeOrder[0]]
}

func (f *Face) Split(v1, v2 *Vertex) {

	// TODO errors if number of vertices are wrong for split or nil etc etc
	e := NewEdge()
	if len(f.edges) == 0 {
		// single closed face, add ring edge
		f.edges[e.id.String()] = e
		f.edgeOrder = append(f.edgeOrder, e.id.String())
	}

	e.vertices[v1.id.String()] = v1
	e.vertices[v2.id.String()] = v1

	// new face is creates
	//newFace := NewFace()
	//newFace.surface = f.surface.Clone() // copy surface into new face

	// TODO reassign edges to new face

}

// Attach sticks the give geometry to the Face, and existing geometry will be removed
func (f *Face) Attach(g Geometry) *Face {
	f.surface = g
	return f
}

// Detach removes geometry from the Face
func (f *Face) Detach() *Face {
	f.surface = nil
	return f
}

// Geometry returns attached geometry
func (f *Face) Geometry() Geometry {
	return f.surface
}
