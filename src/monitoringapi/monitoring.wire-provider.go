package monitoringapi

import "github.com/google/wire"

func ProvideMonitoringController() MonitoringController {
  return NewMonitoringController()
}

var MonitoringApiSet = wire.NewSet(ProvideMonitoringController)
