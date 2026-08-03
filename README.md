# LoadBalancer


## Structure

- `cmd/loadbalancer/` - entrypoint for the executable
- `internal/config/` - configuration loading and validation
- `internal/health/` - health check logic
- `internal/http/` - HTTP-specific helpers and handlers
- `internal/loadbalancer/` - core balancing logic
- `internal/logging/` - logging setup and adapters
- `internal/service/` - backend service models and registry
- `internal/transport/` - transport adapters for future protocols
- `pkg/` - reusable packages intended for external use later
- `configs/` - example or local config files
- `scripts/` - helper scripts
- `tests/` - integration or end-to-end tests
- `docs/` - design notes and docs
