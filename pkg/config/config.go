package config

import (
	"os"

	root "github.com/c-o-l-o-r/watchtower/pkg"
)

func GetConfig() *root.Config {
	return &root.Config{
		Server:     &root.ServerConfig{Port: envOrDefaultString("watchtower:server:port", ":8080")},
		Kubernetes: &root.KubernetesConfig{Namespace: envOrDefaultString("watchtower:namespace", "watchtower")},
	}
}

func envOrDefaultString(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
