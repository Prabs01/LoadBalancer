package tests

import (
	"os"
	"path/filepath"
	"testing"

	"loadbalancer/configs"
)

func TestLoadConfig_LoadsExpectedValues(t *testing.T) {
	configPath := filepath.Join("testdata", "config.yaml")

	cfg, err := configs.NewConfigManager(configPath).LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig returned error: %v", err)
	}

	if cfg.Mode != 1 {
		t.Fatalf("expected mode 1, got %d", cfg.Mode)
	}
	if cfg.Listen != ":8080" {
		t.Fatalf("expected listen :8080, got %q", cfg.Listen)
	}
	if cfg.Balancer != "round_robin" {
		t.Fatalf("expected balancer round_robin, got %q", cfg.Balancer)
	}
	if cfg.Timeout.Read != 5 {
		t.Fatalf("expected timeout.read 5, got %d", cfg.Timeout.Read)
	}
	if cfg.Health.Passive.Fall != 3 {
		t.Fatalf("expected health.passive.fall 3, got %d", cfg.Health.Passive.Fall)
	}
	if len(cfg.Routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(cfg.Routes))
	}
	if cfg.Routes[0].Match.PathPrefix != "/api" {
		t.Fatalf("expected route path_prefix /api, got %q", cfg.Routes[0].Match.PathPrefix)
	}
	if len(cfg.Pools["public"]) != 2 {
		t.Fatalf("expected 2 pools for public, got %d", len(cfg.Pools["public"]))
	}
	if cfg.Pools["public"][0].Addr != "10.0.0.1:80" {
		t.Fatalf("expected first pool addr to be 10.0.0.1:80, got %q", cfg.Pools["public"][0].Addr)
	}
	if cfg.Pools["public"][0].Weight != 3 {
		t.Fatalf("expected first pool weight to be 3, got %d", cfg.Pools["public"][0].Weight)
	}
	if cfg.Pools["public"][1].Addr != "10.0.0.2:80" {
		t.Fatalf("expected second pool addr to be 10.0.0.2:80, got %q", cfg.Pools["public"][1].Addr)
	}
	if cfg.Pools["public"][1].Weight != 1 {
		t.Fatalf("expected second pool weight to be 1, got %d", cfg.Pools["public"][1].Weight)
	}
}

func TestLoadConfig_ReturnsErrorWhenFileMissing(t *testing.T) {
	configPath := filepath.Join("testdata", "missing.yaml")

	_, err := configs.NewConfigManager(configPath).LoadConfig()
	if err == nil {
		t.Fatal("expected error for missing config file")
	}
	if !os.IsNotExist(err) && !os.IsNotExist(os.ErrNotExist) {
		t.Fatalf("expected not-exist error, got %v", err)
	}
}
