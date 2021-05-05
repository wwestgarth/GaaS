package restapi

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/wwestgarth/remote-geometry/geometry"
	"github.com/wwestgarth/remote-geometry/models"
	"github.com/wwestgarth/remote-geometry/restapi/geometry_server"
)

type Sphere struct {
	centre mgl64.Vec3
	radius float64
}

func createSphereHandler(params geometry_server.CreatesphereParams) (resp middleware.Responder) {

	var centre [3]float64
	for i, v := range params.Request.Centre {
		centre[i] = v
	}
	s := geometry.NewSphere(*params.Request.Radius, centre)

	s.Validate()

	return &geometry_server.CreatesphereCreated{Payload: &models.CreateGeometryResult{
		ID: uuid.New().String(),
	}}
}
