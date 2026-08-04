**BackendManager**

A concise guide to the `backendmanager` package.

**Overview**

The `backendmanager` package models individual backend servers and backend pools used
by the load balancer and health manager. It provides small, concurrency-safe helpers
for tracking weights, connection counts, failure counts and cooldown state.

**Core Types**

`Backend` : represents a single backend server and exposes getters/setters for:
  - address (`Addr`)
  - weight (`GetWeight`, `SetWeight`)
  - connection count (`IncrementConnectionCount`, `DecrementConnectionCount`, `GetConnectionCount`)
  - failure count (`IncrementFailureCount`, `ResetFailureCount`, `GetFailureCount`)
  - cooldown timestamp (`SetCooldownUntil`, `GetCooldownUntil`)

`BackendPool` : a named collection of `Backend` instances with helpers:
  - `Backends()` : returns the underlying slice
  - `PoolSize()` : returns number of backends
  - `GetHealthyBackends()`, `GetUnhealthyBackends()`, `GetTrialBackends()`
  - `GetConnectionCounts()` : maps addresses to their active connection counts

**HealthState**

Constants representing a backend health state:

- `Healthy`
- `Unhealthy`
- `Trial`

**Usage Example**

Create a backend and pool:

```go
import "loadbalancer/internal/backendmanager"

b := backendmanager.NewBackend("10.0.0.1:80", 1, backendmanager.Healthy)
pool := backendmanager.NewBackendPool("public", []*backendmanager.Backend{b})
```

Query connection counts:

```go
counts := pool.GetConnectionCounts()
```

Notes

- Methods that update counters use atomic primitives; they are safe for concurrent use.
- Cooldown times are stored as Unix nanoseconds via `SetCooldownUntil` / `GetCooldownUntil`.
