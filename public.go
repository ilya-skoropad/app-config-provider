package app_config_provider

import (
	"errors"
	"slices"

	"github.com/ilya-skoropad/app-config-provider/internal"
)

const (
	Prod = 'p'
	Test = 't'
	Dev  = 'd'

	Env  = 'e'
	File = 'f'

	ErrorWrongEnv      = "E-ML6DZY8C"
	ErrorWrongProvider = "E-JZ5GEWG6"
)

type ConfigProvider interface {
	Getenv(string) string
}

type ConfigManager struct {
	rules       map[rune]rune
	envFileName string
}

func (cm *ConfigManager) GetConfigProvider(envirinment rune) ConfigProvider {
	if cm.rules[envirinment] == Env {
		return internal.NewOsProvider()
	}

	return internal.NewFileProvider(cm.envFileName)
}

func NewConfigManager(rules map[rune]rune, envFileName string) (*ConfigManager, error) {
	err := validateRules(rules)
	if err != nil {
		return nil, err
	}

	return &ConfigManager{
		rules:       rules,
		envFileName: envFileName,
	}, nil
}

func validateRules(rules map[rune]rune) error {
	for envName, provName := range rules {
		if !slices.Contains([]rune{Prod, Test, Dev}, envName) {
			return errors.New(ErrorWrongEnv)
		}

		if !slices.Contains([]rune{Env, File}, provName) {
			return errors.New(ErrorWrongProvider)
		}
	}

	return nil
}
