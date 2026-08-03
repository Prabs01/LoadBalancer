package loadbalancer

import (
	"loadbalancer/internal/backend_pool"
)

type LeastConnectionLoadBalancer struct {
}

func NewLeastConnectionLoadBalancer() *LeastConnectionLoadBalancer {
	return &LeastConnectionLoadBalancer{}
}

func (lc *LeastConnectionLoadBalancer) NextBackend(bp backend_pool.BackendPool) *backend_pool.Backend {
	backends := bp.Backends()
	if len(backends) == 0 {
		return nil
	}

	var leastConnBackend *backend_pool.Backend
	for _, backend := range backends {
		if leastConnBackend == nil || backend.GetConnectionCount() < leastConnBackend.GetConnectionCount() {
			leastConnBackend = backend
		}
	}
	return leastConnBackend
}