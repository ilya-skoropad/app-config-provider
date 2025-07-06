package app_config_provider

import (
	"errors"
	"os"
	"slices"

	ie "github.com/ilya-skoropad/app-config-provider/errors"
	"github.com/ilya-skoropad/app-config-provider/internal"
)

const (
	Prod = 'p'
	Test = 't'
	Dev  = 'd'

	Env  = 'e'
	File = 'f'
)

type ConfigProvider interface {
	Getenv(string) string
}

type ConfigManager struct {
	rules       map[byte]byte
	envFileName string
}

func (cm *ConfigManager) GetConfigProvider(envName string) (ConfigProvider, error) {
	src := os.Getenv(envName)

	if len(src) > 1 {
		return nil, errors.New(ie.ErrorWrongEnvName)
	}

	var envirinment byte = src[0]
	if len(src) == 0 {
		envirinment = Dev
	}

	if cm.rules[envirinment] == Env {
		return internal.NewOsProvider(), nil
	}

	return internal.NewFileProvider(cm.envFileName)
}

func NewConfigManager(rules map[byte]byte, envFileName string) (*ConfigManager, error) {
	err := validateRules(rules)
	if err != nil {
		return nil, err
	}

	return &ConfigManager{
		rules:       rules,
		envFileName: envFileName,
	}, nil
}

func validateRules(rules map[byte]byte) error {
	for envName, provName := range rules {
		if !slices.Contains([]byte{Prod, Test, Dev}, envName) {
			return errors.New(ie.ErrorWrongEnv)
		}

		if !slices.Contains([]byte{Env, File}, provName) {
			return errors.New(ie.ErrorWrongProvider)
		}
	}

	return nil
}
