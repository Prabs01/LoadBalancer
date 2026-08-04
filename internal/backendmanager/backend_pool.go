package backendmanager

type BackendPool struct {
	Name     string
	backends []*Backend
	poolSize int
}

func (bp *BackendPool) Backends() []*Backend {
	return bp.backends
}

func (bp *BackendPool) PoolSize() int {
	return bp.poolSize
}

func (bp *BackendPool) GetHealthyBackends() []*Backend {
	var healthyBackends []*Backend
	for i := range bp.backends {
		if bp.backends[i].IsHealthy() {
			healthyBackends = append(healthyBackends, bp.backends[i])
		}
	}
	return healthyBackends
}

func (bp *BackendPool) GetUnhealthyBackends() []*Backend {
	var unhealthyBackends []*Backend
	for i := range bp.backends {
		if bp.backends[i].IsUnhealthy() {
			unhealthyBackends = append(unhealthyBackends, bp.backends[i])
		}
	}
	return unhealthyBackends
}

func (bp *BackendPool) GetTrialBackends() []*Backend {
	var trialBackends []*Backend
	for i := range bp.backends {
		if bp.backends[i].IsTrial() {
			trialBackends = append(trialBackends, bp.backends[i])
		}
	}
	return trialBackends
}

func (bp *BackendPool) GetConnectionCounts() map[string]int {
	connectionCounts := make(map[string]int)
	for _, b := range bp.backends {
		connectionCounts[b.Addr] = int(b.connectionCount.Load())
	}
	return connectionCounts
}

func NewBackendPool(name string, backends []*Backend) *BackendPool {
	return &BackendPool{
		Name:     name,
		backends: backends,
		poolSize: len(backends),
	}
}
