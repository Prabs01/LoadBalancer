package loadbalancer

import (
	backendpool "loadbalancer/internal/backend_pool"
)

type LoadBalancer interface {
	NextBackend(*backendpool.BackendPool) *backendpool.Backend
}
