package monitoringapi

import "github.com/gin-gonic/gin"

type MonitoringController struct {
}

func NewMonitoringController() MonitoringController {
  return MonitoringController{}
}

func (m *MonitoringController) ConfigureRoutes(e *gin.Engine) {
  e.GET("/healthcheck", m.Healthcheck)
  e.GET("/metrics", m.Metrics)
}

func (m *MonitoringController) Healthcheck(c *gin.Context) {
  c.Status(200)
}

func (m *MonitoringController) Metrics(c *gin.Context) {
  // TODO: implement Prometheus
  // for now, just return empty
  c.Header("Content-Type", "application/openmetrics-text")
  c.String(200, "")
}
