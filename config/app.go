package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Name string `split_words:"true" json:"APP_NAME"`
	Env  string `split_words:"true" json:"APP_ENV"`
	Port int    `split_words:"true" json:"APP_PORT"`
}

var App *AppConfig

func loadAppConfig() {
	App = &AppConfig{}

	err := envconfig.Process("app", App)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
