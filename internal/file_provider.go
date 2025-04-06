package internal

type FileProvider struct {
	file     map[string]string
	fileName string
}

func (e *FileProvider) Getenv(name string) string {
	return e.file[name]
}

func NewFileProvider(fileName string) *FileProvider {
	return &FileProvider{
		fileName: fileName,
	}
}
