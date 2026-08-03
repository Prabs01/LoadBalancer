# Config Manager

Overview

The Config Manager loads application configuration from a YAML file and returns a `Config` struct used by the application. It is implemented in the `configs` package.

Usage

- Initialize: `configs.NewConfigManager(path)`
- Load: call `LoadConfig()` on the returned manager to retrieve a `*configs.Config` or an error.

Config schema

Example YAML (see `tests/testdata/config.yaml` for a concrete fixture used by tests):

```yaml
mode: 1
listen: ":8080"
balancer: round_robin
timeout:
  read: 5
  dial: 2
  write: 5
  idle: 30
health:
  passive:
    fall: 3
    cooldown: 60
routes:
  - match:
      path_prefix: /api
    pool: public
pools:
  public:
    - addr: "10.0.0.1:80"
    - addr: "10.0.0.2:80"
```

Field mapping (struct -> YAML)

- `Config.Mode` -> `mode` (int)
- `Config.Listen` -> `listen` (string)
- `Config.Balancer` -> `balancer` (string)
- `Config.Timeout.Read` -> `timeout.read` (int)
- `Config.Timeout.Dial` -> `timeout.dial` (int)
- `Config.Timeout.Write` -> `timeout.write` (int)
- `Config.Timeout.Idle` -> `timeout.idle` (int)
- `Config.Health.Passive.Fall` -> `health.passive.fall` (int)
- `Config.Health.Passive.Cooldown` -> `health.passive.cooldown` (int)
- `Config.Routes` -> `routes` (array of objects with `match.path_prefix` and `pool`)
- `Config.Pools` -> `pools` (map of pool name to list of `addr` entries)

Error behavior

- If the config file cannot be read, `LoadConfig()` returns an error indicating the file could not be opened.
- If the file is present but cannot be parsed as YAML into the `Config` struct, `LoadConfig()` returns an unmarshal error.

Testing

The test fixture used by the package is at `tests/testdata/config.yaml`. Tests exercise successful loading and the missing-file error case.

Notes

- The YAML tags and exported struct fields in `configs/config_manager.go` must match the YAML keys shown above for `yaml.Unmarshal` to populate the struct correctly.
- If you change the YAML layout, update the struct tags and this doc accordingly.
