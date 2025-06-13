package config

type AppConfig struct {
	Name            string `json:"name"`
	Debug           bool   `json:"debug"`
	Env             string `json:"env"`
	Url             string `json:"url"`
	Version         string `json:"version"`
	Timezone        string `json:"timezone"`
	Key             string `json:"key"`
	ResourcesByPage int    `json:"resources_by_page"`
}
