package internal

type FileProvider struct {
	envs map[string]string
}

func (e *FileProvider) Getenv(name string) string {
	return e.envs[name]
}

func NewFileProvider(fileName string) (*FileProvider, error) {
	data, err := ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	envs, err := ParseFile(data)
	if err != nil {
		return nil, err
	}

	return &FileProvider{
		envs: envs,
	}, nil
}
