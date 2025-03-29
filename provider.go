package app_config_provider

const (
	Prod = 'p'
	Test = 't'
	Dev  = 'd'
)

type ConfigProvider interface {
	Getenv(string) string
}

func GetConfigProvider(envirinment rune) ConfigProvider {
	switch envirinment {
	case Prod, Test:
		return NewOsProvider()

	default:
		return NewFileProvider(".env")
	}
}
