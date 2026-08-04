package backendmanager

import (
	"sync/atomic"
)

type HealthState uint8

const (
	Healthy HealthState = iota
	Unhealthy
	Trial
)

func (hs HealthState) String() string {
	switch hs {
	case Healthy:
		return "Healthy"
	case Unhealthy:
		return "Unhealthy"
	case Trial:
		return "Trial"
	default:
		return "Unknown"
	}
}

type Backend struct {
	Addr            string
	weight          atomic.Int64
	healthState     atomic.Uint32 
	connectionCount atomic.Int64
	failureCount    atomic.Int64
	cooldownUntil   atomic.Int64 // Store as Unix timestamp in nanoseconds
}

func NewBackend(addr string, weight int, healthState HealthState) *Backend {
	var backend Backend
	backend.Addr = addr
	backend.weight.Store(int64(weight))
	backend.healthState.Store(uint32(healthState))
	backend.connectionCount.Store(0)
	backend.failureCount.Store(0)
	backend.cooldownUntil.Store(0)
	return &backend
}

func (b *Backend) GetWeight() int {
	return int(b.weight.Load())
}

func (b *Backend) SetWeight(weight int) {
	b.weight.Store(int64(weight))
}

func (b *Backend) GetHealthState() HealthState {
	return HealthState(b.healthState.Load())
}

func (b *Backend) SetHealthState(state HealthState) {
	b.healthState.Store(uint32(state))
}

func (b *Backend) IsHealthy() bool {
	return b.healthState.Load() == uint32(Healthy)
}

func (b *Backend) IsTrial() bool {
	return b.healthState.Load() == uint32(Trial)
}

func (b *Backend) IsUnhealthy() bool {
	return b.healthState.Load() == uint32(Unhealthy)
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

func (b *Backend) IncrementFailureCount() {
	b.failureCount.Add(1)
}

func (b *Backend) ResetFailureCount() {
	b.failureCount.Store(0)
}

func (b *Backend) GetFailureCount() int {
	return int(b.failureCount.Load())
}

func (b *Backend) SetCooldownUntil(t int64) {
	b.cooldownUntil.Store(t)
}

func (b *Backend) GetCooldownUntil() int64 {
	return b.cooldownUntil.Load()
}