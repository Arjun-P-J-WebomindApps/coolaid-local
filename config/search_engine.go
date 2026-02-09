package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type SearchEngineConfig struct {
	TypesenseAPIEndpoint string `split_words:"true"`
	TypesenseAPIKey      string `split_words:"true"`
}

var SearchEngine *SearchEngineConfig

func loadSearchEngineConfig() {
	SearchEngine = &SearchEngineConfig{}

	err := envconfig.Process("SEARCH_ENGINE", SearchEngine)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
