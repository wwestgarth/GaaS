package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/wwestgarth/remote-geometry/models"
	"github.com/wwestgarth/remote-geometry/restapi/geometry_server"
)

func handleHealth(params geometry_server.HealthParams) (resp middleware.Responder) {

	// Hardcode for now
	status := "pass"
	return &geometry_server.HealthOK{
		Payload: &models.HealthResult{
			Status: &status,
		},
	}
}
