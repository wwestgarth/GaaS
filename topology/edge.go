package topology

import (
	"github.com/google/uuid"
	"github.com/wwestgarth/remote-geometry/vector"
)

type Vertex struct {
	id    uuid.UUID
	point vector.Vector
}

func NewVertex() *Vertex {
	return &Vertex{
		id: uuid.New(),
	}
}

func (v *Vertex) Attach(p vector.Vector) *Vertex {
	v.point = p
	return v
}

type Edge struct {
	id       uuid.UUID
	curve    Geometry
	vertices map[string]*Vertex
	faces    map[string]*Face
}

func NewEdge() *Edge {
	return &Edge{
		id:       uuid.New(),
		vertices: map[string]*Vertex{},
		faces:    map[string]*Face{},
	}
}

func (e *Edge) Geometry() Geometry {
	return e.curve
}

func (e *Edge) Attach(g Geometry) *Edge {
	e.curve = g
	return e
}

// Split an edge by creating a new vertex, a new edge might be created
func (e *Edge) Split() (*Vertex, *Edge) {

	v := NewVertex()

	if len(e.vertices) == 0 {
		// ring edge, split it, no new edge
		e.vertices[v.id.String()] = v
		return v, nil
	}

	newEdge := NewEdge()
	//newEdge.curve = e.curve.Clone()

	newEdge.faces = e.faces // TODO: share map or not? will need to lock for concurrency
	e.vertices[v.id.String()] = v
	newEdge.vertices[v.id.String()] = v

	return nil, nil
}
