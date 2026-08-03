package backend_pool

import (
	"loadbalancer/internal/health"
	"sync/atomic"
)

type Backend struct {
	Addr        string
	weight      atomic.Int64
	healthState health.HealthState
	connectionCount atomic.Int64
}

type BackendPool struct {
	Name     string
	Backends []*Backend
}

func (bp *BackendPool) GetHealthyBackends() []*Backend {
	var healthyBackends []*Backend
	for i := range bp.Backends {
		if bp.Backends[i].healthState == health.Healthy {
			healthyBackends = append(healthyBackends, bp.Backends[i])
		}
	}
	return healthyBackends
}

func (bp *BackendPool) GetUnhealthyBackends() []*Backend {
	var unhealthyBackends []*Backend
	for i := range bp.Backends {
		if bp.Backends[i].healthState == health.Unhealthy {
			unhealthyBackends = append(unhealthyBackends, bp.Backends[i])
		}
	}
	return unhealthyBackends
}

func (bp *BackendPool) GetTrialBackends() []*Backend {
	var trialBackends []*Backend
	for i := range bp.Backends {
		if bp.Backends[i].healthState == health.Trial {
			trialBackends = append(trialBackends, bp.Backends[i])
		}
	}
	return trialBackends
}

func (bp *BackendPool) GetConnectionCounts() map[string]int {
	connectionCounts := make(map[string]int)
	for _, b := range bp.Backends {
		connectionCounts[b.Addr] = int(b.connectionCount.Load())
	}
	return connectionCounts
}


func NewBackend(addr string, weight int, healthState health.HealthState) *Backend {
	var backend Backend
	backend.Addr = addr
	backend.weight.Store(int64(weight))
	backend.healthState = healthState
	backend.connectionCount.Store(0)
	return &backend
}

func NewBackendPool(name string, backends []*Backend) *BackendPool {
	return &BackendPool{
		Name:     name,
		Backends: backends,
	}
}


