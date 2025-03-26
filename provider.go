package app_config_provider

import "os"

type ConfigProvider interface {
	Getenv(string) string
}

func GetConfigProvider() ConfigProvider {
	// First let's decide do we have env variable or not.
	env := os.Getenv("env")

	switch env {
	case "prod":
		return NewOsProvider()

	default:
		return NewFileProvider(".env")
	}
}
