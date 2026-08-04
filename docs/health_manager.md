**Health Manager**

Documentation for the `health` package `HealthManager`.

**Overview**

`HealthManager` is responsible for tracking backend failures, moving backends between
`Healthy`, `Unhealthy`, and `Trial` states, and enforcing cooldown windows after
thresholded failures.

**Constructor**

`NewHealthManager(pool *backendmanager.BackendPool, threshold int, cooldown time.Duration) *HealthManager`

- `pool`: the `BackendPool` the manager observes (may be `nil` if manager does not need to enumerate pool members)
- `threshold`: number of consecutive failures required to mark a `Healthy` backend `Unhealthy`
- `cooldown`: duration used to set the cooldown timestamp when a backend becomes `Unhealthy`

**Primary Methods**

- `RecordSuccess(backend *backendmanager.Backend) error` : resets the backend failure count.
- `RecordFailure(backend *backendmanager.Backend) error` : increments failure count; if count >= threshold and state is `Healthy`, marks `Unhealthy` and sets cooldown; if state is `Trial` and a failure occurs, marks `Unhealthy` and sets cooldown.
- `CoolDownExpired(backend *backendmanager.Backend) error` : when the cooldown timestamp has passed and the backend is `Unhealthy`, transitions it to `Trial`, resets failure count and cooldown timestamp.

**Behavior Notes**

- Failure counts and cooldown timestamps are stored on the `Backend` and managed via
  atomic operations; callers can safely call these methods concurrently.
- `CoolDownExpired` does not force a transition unless the cooldown has actually expired.

**Example**

```go
import (
    "time"
    "loadbalancer/internal/backendmanager"
    "loadbalancer/internal/health"
)

b := backendmanager.NewBackend("10.0.0.1:80", 1, backendmanager.Healthy)
pool := backendmanager.NewBackendPool("p", []*backendmanager.Backend{b})
hm := health.NewHealthManager(pool, 3, 30*time.Second)

// record a failure
_ = hm.RecordFailure(b)

// after cooldown expires
_ = hm.CoolDownExpired(b)
```
