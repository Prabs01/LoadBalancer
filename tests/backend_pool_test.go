package tests

import (
	"loadbalancer/internal/backend_pool"
	"loadbalancer/internal/health"
	"testing"
)

func TestNewBackendPool(t *testing.T) {
	backends := []*backend_pool.Backend{backend_pool.NewBackend("10.0.0.1:80", 1, health.Healthy)}
	bp := backend_pool.NewBackendPool("public", backends)
	if bp.Name != "public" {
		t.Fatalf("expected name public, got %q", bp.Name)
	}
	if len(bp.Backends) != 1 {
		t.Fatalf("expected 1 backend, got %d", len(bp.Backends))
	}		
	if bp.Backends[0].Addr != "10.0.0.1:80" {
		t.Fatalf("expected backend addr 10.0.0.1:80, got %q", bp.Backends[0].Addr)
	}
}

func TestGetHealthyBackends(t *testing.T) {
	backends := []*backend_pool.Backend{
		backend_pool.NewBackend("10.0.0.1:80", 1, health.Healthy),
		backend_pool.NewBackend("10.0.0.2:80", 1, health.Unhealthy),
		backend_pool.NewBackend("10.0.0.3:80", 1, health.Healthy),
	}

	bp := backend_pool.NewBackendPool("public", backends)
	healthy := bp.GetHealthyBackends()

	if len(healthy) != 2 {
		t.Fatalf("expected 2 healthy backends, got %d", len(healthy))
	}
	if healthy[0] != bp.Backends[0] {
		t.Fatalf("expected first healthy backend to point to the first slice element")
	}
	if healthy[1] != bp.Backends[2] {
		t.Fatalf("expected second healthy backend to point to the third slice element")
	}
	if healthy[0].Addr != "10.0.0.1:80" || healthy[1].Addr != "10.0.0.3:80" {
		t.Fatalf("unexpected healthy backend addresses: %+v", healthy)
	}
}

func TestGetUnhealthyBackends(t *testing.T) {
	backends := []*backend_pool.Backend{
		backend_pool.NewBackend("10.0.0.1:80", 1, health.Healthy),
		backend_pool.NewBackend("10.0.0.2:80", 1, health.Unhealthy),
		backend_pool.NewBackend("10.0.0.3:80", 1, health.Unhealthy),
	}

	bp := backend_pool.NewBackendPool("public", backends)
	unhealthy := bp.GetUnhealthyBackends()

	if len(unhealthy) != 2 {
		t.Fatalf("expected 2 unhealthy backends, got %d", len(unhealthy))
	}
	if unhealthy[0] != bp.Backends[1] {
		t.Fatalf("expected first unhealthy backend to point to the second slice element")
	}
	if unhealthy[1] != bp.Backends[2] {
		t.Fatalf("expected second unhealthy backend to point to the third slice element")
	}
}

func TestGetTrialBackends(t *testing.T) {
	backends := []*backend_pool.Backend{
		backend_pool.NewBackend("10.0.0.1:80", 1, health.Trial),
		backend_pool.NewBackend("10.0.0.2:80", 1, health.Healthy),
		backend_pool.NewBackend("10.0.0.3:80", 1, health.Trial),
	}

	bp := backend_pool.NewBackendPool("public", backends)
	trial := bp.GetTrialBackends()

	if len(trial) != 2 {
		t.Fatalf("expected 2 trial backends, got %d", len(trial))
	}
	if trial[0] != bp.Backends[0] {
		t.Fatalf("expected first trial backend to point to the first slice element")
	}
	if trial[1] != bp.Backends[2] {
		t.Fatalf("expected second trial backend to point to the third slice element")
	}
}

func TestGetConnectionCounts(t *testing.T) {
	backends := []*backend_pool.Backend{
		backend_pool.NewBackend("10.0.0.1:80", 1, health.Healthy),
		backend_pool.NewBackend("10.0.0.2:80", 1, health.Healthy),
	}

	bp := backend_pool.NewBackendPool("public", backends)
	counts := bp.GetConnectionCounts()

	if len(counts) != 2 {
		t.Fatalf("expected 2 connection count entries, got %d", len(counts))
	}
	if counts["10.0.0.1:80"] != 0 {
		t.Fatalf("expected initial count for 10.0.0.1:80 to be 0, got %d", counts["10.0.0.1:80"])
	}
	if counts["10.0.0.2:80"] != 0 {
		t.Fatalf("expected initial count for 10.0.0.2:80 to be 0, got %d", counts["10.0.0.2:80"])
	}
}
