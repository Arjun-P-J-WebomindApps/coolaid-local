package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host       string `split_words:"true" json:"HOST"`
	Port       string `split_words:"true" json:"PORT"`
	User       string `split_words:"true" json:"USER"`
	Password   string `split_words:"true" json:"PASSWORD"`
	SslMode    string `split_words:"true" json:"SSL_MODE"`
	Name       string `split_words:"true" json:"NAME"`
	SearchPath string `split_words:"true" json:"SEARCH_PATH"`
}

var DB *DBConfig

func loadDBConfig() {

	DB = &DBConfig{}

	err := envconfig.Process("db", DB)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
