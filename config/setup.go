package config

import (
	"platform/pkg/env"
	"strconv"
)

type config struct {
	serverConfig ServerConfig
	appConfig    AppConfig
}

type Config interface {
	ServerConfig() ServerConfig
	AppConfig() AppConfig
}

func NewConfig() Config {
	resourcesByPage, err := strconv.Atoi(env.GetEnv("APP_RESOURCES_BY_PAGE", "10"))
	if err != nil {
		panic(err)
	}

	appName := env.GetEnv("APP_NAME", "another-idea")

	return &config{
		serverConfig: ServerConfig{
			Port: ":8080",
		},
		appConfig: AppConfig{
			Name:            appName,
			Version:         env.GetEnv("APP_VERSION", "1.0.0"),
			Debug:           env.GetEnv("APP_DEBUG", "false") == "true",
			Env:             env.GetEnv("APP_ENV", "development"),
			Url:             env.GetEnv("APP_URL", "http://localhost:8080"),
			Timezone:        env.GetEnv("APP_TIMEZONE", "Asia/Tehran"),
			Key:             env.GetEnv("APP_KEY", "key"),
			ResourcesByPage: resourcesByPage,
		},
	}
}

func (c *config) ServerConfig() ServerConfig {
	return c.serverConfig
}

func (c *config) AppConfig() AppConfig {
	return c.appConfig
}
