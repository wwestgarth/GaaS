package topology

import "github.com/google/uuid"

type Body struct {
	id      uuid.UUID
	faceAny *Face
}

// NewBody returns a single closed faced body
func NewBody() *Body {
	return &Body{
		id:      uuid.New(),
		faceAny: NewFace(),
	}
}

func (b *Body) FaceAny() *Face {
	return b.faceAny
}

// NewSheetBody returns a single faces single edged sheet body
func NewSheetBody() *Body {

	// One face one edge
	f := NewFace()
	e := NewEdge()

	// connect them
	e.faces[f.id.String()] = f
	f.edges[e.id.String()] = e
	f.edgeOrder = []string{e.id.String()}

	// Return
	return &Body{
		id:      uuid.New(),
		faceAny: f,
	}
}

func (b *Body) ID() string {
	return b.id.String()
}

func (b *Body) Validate() error {

	// TODO properly

	f := b.FaceAny()
	if err := f.Geometry().Validate(); err != nil {
		return err
	}

	e := f.EdgeAny()
	if err := e.Geometry().Validate(); err != nil {
		return err
	}

	return nil
}
