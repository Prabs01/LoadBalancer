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

func NewBackend(addr string, weight int, healthState health.HealthState) *Backend {
	var backend Backend
	backend.Addr = addr
	backend.weight.Store(int64(weight))
	backend.healthState = healthState
	backend.connectionCount.Store(0)
	return &backend
}

func (b *Backend) GetWeight() int {
	return int(b.weight.Load())
}

func (b *Backend) SetWeight(weight int) {
	b.weight.Store(int64(weight))
}

func (b *Backend) GetHealthState() health.HealthState {
	return b.healthState
}

func (b *Backend) SetHealthState(state health.HealthState) {
	b.healthState = state
}

func (b *Backend) IncrementConnectionCount() {
	b.connectionCount.Add(1)
}

func (b *Backend) DecrementConnectionCount() {
	b.connectionCount.Add(-1)
}	

func (b *Backend) GetConnectionCount() int {
	return int(b.connectionCount.Load())
}


