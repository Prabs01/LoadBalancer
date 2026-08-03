package backendpool

import (
	"loadbalancer/internal/health"
)

type Backend struct {
	Addr string
	Weight int
	HealthState health.HealthState
}



type BackendPool struct {
	Name     string
	Backends []Backend
}
func (bp *BackendPool) GetHealthyBackends() []*Backend {
	var healthyBackends []*Backend
	for _, b := range bp.Backends {
		if b.HealthState == health.Healthy {
			healthyBackends = append(healthyBackends, &b)
		}
	}
	return healthyBackends
}



func (bp *BackendPool) GetUnhealthyBackends() []*Backend {
	var unhealthyBackends []*Backend
	for _, b := range bp.Backends {
		if b.HealthState == health.Unhealthy {
			unhealthyBackends = append(unhealthyBackends, &b)
		}
	}
	return unhealthyBackends
}

func (bp *BackendPool) GetTrialBackends() []*Backend {
	var trialBackends []*Backend
	for _, b := range bp.Backends {
		if b.HealthState == health.Trial {
			trialBackends = append(trialBackends, &b)
		}
	}
	return trialBackends
}	

func (bp *BackendPool) GetConnectionCounts() map[string]int {
	connectionCounts := make(map[string]int)
	for _, b := range bp.Backends {
		connectionCounts[b.Addr] = 0
	}
	return connectionCounts
}



func NewBackendPool(name string, backends []Backend) *BackendPool {
	return &BackendPool{
		Name:     name,
		Backends: backends,
	}
}		




