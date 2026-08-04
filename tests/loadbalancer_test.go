package tests

import (
	"testing"

	"loadbalancer/internal/backendmanager"
	lb "loadbalancer/internal/loadbalancer"
)

// simple round-robin implementation used only for testing
type rrLB struct {
	last int
}

func (r *rrLB) NextBackend(pool *backendmanager.BackendPool) *backendmanager.Backend {
	if pool.PoolSize() == 0 {
		return nil
	}
	r.last = (r.last + 1) % pool.PoolSize()
	return pool.Backends()[r.last]
}

func TestRRImplementsInterface(t *testing.T) {
	var _ lb.LoadBalancer = (*rrLB)(nil)
}

func TestRRSelectionOrder(t *testing.T) {

	backends := []*backendmanager.Backend{
		backendmanager.NewBackend("a:1", 1, backendmanager.Healthy),
		backendmanager.NewBackend("b:2", 1, backendmanager.Healthy),
		backendmanager.NewBackend("c:3", 1, backendmanager.Healthy),
	}
	pool := backendmanager.NewBackendPool("test", backends)

	rl := &rrLB{last: -1}

	want := []string{"a:1", "b:2", "c:3", "a:1", "b:2"}
	for i, w := range want {
		b := rl.NextBackend(pool)
		if b == nil {
			t.Fatalf("expected backend, got nil at iteration %d", i)
		}
		if b.Addr != w {
			t.Fatalf("iteration %d: expected %s, got %s", i, w, b.Addr)
		}
	}
}
