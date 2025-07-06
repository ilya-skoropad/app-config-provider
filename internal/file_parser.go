package internal

import (
	"errors"
	"os"
	"regexp"
	"strings"

	ie "github.com/ilya-skoropad/app-config-provider/errors"
)

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.New(ie.ErrorCannotReadFile)
	}

	return string(data), nil
}

func ParseFile(data string) (map[string]string, error) {
	lines := strings.Split(data, ENDL)
	envMap := make(map[string]string, len(lines))

	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		r, _ := regexp.Compile("(^[A-Z_]{1,})=(.*)")
		matches := r.FindStringSubmatch(line)
		if len(matches) != 3 {
			return map[string]string{}, errors.New(ie.ErrorIncorrectFile)
		}

		envMap[matches[1]] = matches[2]
	}

	return envMap, nil
}
