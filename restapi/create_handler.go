package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/wwestgarth/remote-geometry/geometry"
	"github.com/wwestgarth/remote-geometry/models"
	"github.com/wwestgarth/remote-geometry/restapi/geometry_server"
)

func createSphereHandler(params geometry_server.CreatesphereParams) (resp middleware.Responder) {

	var centre [3]float64
	for i, v := range params.Request.Centre {
		centre[i] = v
	}
	s := geometry.NewSphere(*params.Request.Radius, centre)

	s.Validate()

	return &geometry_server.CreatesphereCreated{Payload: &models.CreateGeometryResult{
		ID: s.ID(),
	}}
}
