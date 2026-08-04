package health

import (
	"fmt"
	"loadbalancer/internal/backendmanager"
	"time"
)

type HealthManager struct {
	backendPool *backendmanager.BackendPool	
	failureThreshold int
	cooldownTime time.Duration
}

func NewHealthManager(pool *backendmanager.BackendPool, threshold int, cooldown time.Duration) *HealthManager {
	hm := &HealthManager{
		backendPool: pool,
		failureThreshold: threshold,
		cooldownTime: cooldown,
	}
	return hm	
}

func (hm *HealthManager) RecordSuccess(backend_ *backendmanager.Backend) error {
	if backend_ == nil {
		return fmt.Errorf("backend is nil")
	}

	backend_.ResetFailureCount()
	return nil
}


func (hm *HealthManager) RecordFailure(backend_ *backendmanager.Backend) error {
	if backend_ == nil {
		return fmt.Errorf("backend is nil")
	}

	backend_.IncrementFailureCount()

	if (backend_.GetFailureCount() >= int(hm.failureThreshold) && backend_.GetHealthState() == backendmanager.Healthy) {
		backend_.SetHealthState(backendmanager.Unhealthy)

		//cooldown the backend for a specified duration
		backend_.SetCooldownUntil(time.Now().Add(hm.cooldownTime).UnixNano())

	}else if backend_.GetHealthState() == backendmanager.Trial {
		backend_.SetHealthState(backendmanager.Unhealthy)
		backend_.SetCooldownUntil(time.Now().Add(hm.cooldownTime).UnixNano())
	}

	return nil
}

func (hm *HealthManager) CoolDownExpired(backend_ *backendmanager.Backend) error {
	if backend_ == nil {
		return fmt.Errorf("backend is nil")
	}

	if backend_.GetHealthState() == backendmanager.Unhealthy && time.Now().UnixNano() >= backend_.GetCooldownUntil() {
		backend_.SetHealthState(backendmanager.Trial)
		backend_.ResetFailureCount()
		backend_.SetCooldownUntil(0) // Reset cooldown
		return nil
	}
	return nil
}