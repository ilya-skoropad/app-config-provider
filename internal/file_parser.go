package internal

import (
	"errors"
	"os"
	"strings"

	ie "github.com/ilya-skoropad/app-config-provider/errors"
)

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ParseFile(data string) (map[string]string, error) {
	lines := strings.Split(data, ENDL)
	envMap := make(map[string]string, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "=")

		if len(parts) != 2 {
			return nil, errors.New(ie.ErrorIncorrectEnvFile)
		}

		envMap[parts[0]] = parts[1]
	}

	return envMap, nil
}
