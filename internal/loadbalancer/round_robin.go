package loadbalancer

import (
	"loadbalancer/internal/backendmanager"
	"sync/atomic"
)

type RoundRobinLoadBalancer struct {
	currentIndex atomic.Uint64
}

func NewRoundRobinLoadBalancer() *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{
		currentIndex: atomic.Uint64{},
	}
}

func (rr *RoundRobinLoadBalancer) NextBackend(bp backendmanager.BackendPool) *backendmanager.Backend {
	backends := bp.Backends()
	if len(backends) == 0 {
		return nil
	}
	index := rr.currentIndex.Load() % uint64(len(backends))
	backend := backends[index]
	rr.currentIndex.Add(1)
	return backend
}
