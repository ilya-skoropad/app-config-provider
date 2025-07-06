package internal

import "os"

type OsProvider struct {
}

func (e *OsProvider) Getenv(name string) string {
	return os.Getenv(name)
}

func NewOsProvider() *OsProvider {
	return &OsProvider{}
}
