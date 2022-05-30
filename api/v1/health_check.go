package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"privy-test/api/http/request"
	"privy-test/service"
)

type healthCheckHandler struct {
	requestParser      request.Parser
	healthCheckService service.HealthCheckService
}

func NewHealthCheckHandler(
	requestParser request.Parser,
	healthCheckService service.HealthCheckService,
) *healthCheckHandler {
	return &healthCheckHandler{
		requestParser:      requestParser,
		healthCheckService: healthCheckService,
	}
}

// HealthCheck godoc
// @Summary HealthCheck endpoint
// @Description Health Check endpoint
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} healthcheck.HTTPHealthCheckResponse
// @Router /ping [get]
func (hc *healthCheckHandler) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, hc.healthCheckService.HealthCheck(ctx.Request().Context()))
}
