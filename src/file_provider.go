package app_config_provider

import (
	"os"
	"strings"
)

type FileProvider struct {
	file map[string]string
}

func (e *FileProvider) Getenv(name string) string {
	return e.file[name]
}

func NewFileProvider(fileName string) *FileProvider {
	content, err := os.ReadFile(fileName)
	check(err)

	strRep := string(content)
	lines := strings.Split(strRep, ENDL)
	envMap := make(map[string]string, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "=")

		if len(parts) != 2 {
			panic("incorrect env file")
		}

		envMap[parts[0]] = parts[1]
	}

	return &FileProvider{
		file: envMap,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
