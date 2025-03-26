# App Config Provider

Provides very basic functionality for environment variables for go.
Allows to get configs from file or from environment variables.

By default will try to get configs from file, unless there is `env` variable in the environment variables and in set to `prod`

## Install

``` bash
go get -v github.com/ilya-skoropad/app-config-provider@master
```

## Example

``` go
package config

import (
    acp "github.com/ilya-skoropad/app-config-provider"
)

type Config struct {
    AppPort string
    AppHost string
    DBUser  string
    DBName  string
    DBPass  string
    DBPort  string
    DBHost  string
}

func Fetch() Config {
    provider := acp.GetConfigProvider()

    return Config{
        AppPort: provider.Getenv("USER_API_APP_PORT"),
        AppHost: provider.Getenv("USER_API_APP_HOST"),
        DBName:  provider.Getenv("USER_API_DB_NAME"),
        DBUser:  provider.Getenv("USER_API_DB_USER"),
        DBPass:  provider.Getenv("USER_API_DB_PASS"),
        DBPort:  provider.Getenv("USER_API_DB_PORT"),
        DBHost:  provider.Getenv("USER_API_DB_HOST"),
    }
}
```
