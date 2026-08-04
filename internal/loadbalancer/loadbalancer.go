package loadbalancer

import (
	"loadbalancer/internal/backendmanager"
)

type LoadBalancer interface {
	NextBackend(*backendmanager.BackendPool) *backendmanager.Backend
}
