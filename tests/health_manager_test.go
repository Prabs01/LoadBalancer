package tests

import (
	"testing"
	"time"

	"loadbalancer/internal/backendmanager"
	"loadbalancer/internal/health"
)

func TestRecordSuccessAndFailure(t *testing.T) {
	b := backendmanager.NewBackend("10.0.0.1:80", 1, backendmanager.Healthy)
	pool := backendmanager.NewBackendPool("p", []*backendmanager.Backend{b})
	hm := health.NewHealthManager(pool, 3, time.Second)

	if err := hm.RecordFailure(b); err != nil {
		t.Fatalf("unexpected error on RecordFailure: %v", err)
	}
	if b.GetFailureCount() != 1 {
		t.Fatalf("expected failure count 1, got %d", b.GetFailureCount())
	}

	if err := hm.RecordSuccess(b); err != nil {
		t.Fatalf("unexpected error on RecordSuccess: %v", err)
	}
	if b.GetFailureCount() != 0 {
		t.Fatalf("expected failure count reset to 0, got %d", b.GetFailureCount())
	}
}

func TestRecordFailure_ThresholdTriggersUnhealthy(t *testing.T) {
	b := backendmanager.NewBackend("10.0.0.2:80", 1, backendmanager.Healthy)
	pool := backendmanager.NewBackendPool("p", []*backendmanager.Backend{b})
	hm := health.NewHealthManager(pool, 2, time.Second)

	if err := hm.RecordFailure(b); err != nil {
		t.Fatalf("unexpected error on first RecordFailure: %v", err)
	}
	if b.GetHealthState() != backendmanager.Healthy {
		t.Fatalf("expected backend to remain Healthy after 1 failure")
	}

	if err := hm.RecordFailure(b); err != nil {
		t.Fatalf("unexpected error on second RecordFailure: %v", err)
	}
	if b.GetHealthState() != backendmanager.Unhealthy {
		t.Fatalf("expected backend to become Unhealthy after threshold reached")
	}
	if b.GetCooldownUntil() == 0 {
		t.Fatalf("expected cooldown to be set")
	}
}

func TestCoolDownExpired_TransitionsToTrial(t *testing.T) {
	b := backendmanager.NewBackend("10.0.0.3:80", 1, backendmanager.Unhealthy)
	pool := backendmanager.NewBackendPool("p", []*backendmanager.Backend{b})
	hm := health.NewHealthManager(pool, 1, time.Millisecond)

	// simulate cooldown in the past
	b.SetCooldownUntil(time.Now().Add(-time.Second).UnixNano())
	b.IncrementFailureCount()

	if err := hm.CoolDownExpired(b); err != nil {
		t.Fatalf("unexpected error on CoolDownExpired: %v", err)
	}
	if b.GetHealthState() != backendmanager.Trial {
		t.Fatalf("expected backend to be Trial after cooldown expired")
	}
	if b.GetFailureCount() != 0 {
		t.Fatalf("expected failure count to be reset to 0")
	}
	if b.GetCooldownUntil() != 0 {
		t.Fatalf("expected cooldown to be reset to 0")
	}
}
