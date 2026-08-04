package loadbalancer

import (
	"loadbalancer/internal/backendmanager"
)

type LeastConnectionLoadBalancer struct {
}

func NewLeastConnectionLoadBalancer() *LeastConnectionLoadBalancer {
	return &LeastConnectionLoadBalancer{}
}

func (lc *LeastConnectionLoadBalancer) NextBackend(bp backendmanager.BackendPool) *backendmanager.Backend {
	backends := bp.Backends()
	if len(backends) == 0 {
		return nil
	}

	var leastConnBackend *backendmanager.Backend
	for _, backend := range backends {
		if leastConnBackend == nil || backend.GetConnectionCount() < leastConnBackend.GetConnectionCount() {
			leastConnBackend = backend
		}
	}
	return leastConnBackend
}
