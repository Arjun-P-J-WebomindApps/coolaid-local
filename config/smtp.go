package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type SMTPConfig struct {
	Host     string `split_words:"true" json:"SMTP_HOST"`
	Port     string `split_words:"true" json:"SMTP_PORT"`
	Auth     string `split_words:"true" json:"SMTP_AUTH"`
	Username string `split_words:"true" json:"SMTP_USERNAME"`
	Password string `split_words:"true" json:"SMTP_PASSWORD"`
	Secure   string `split_words:"true" json:"SMTP_SECURE"`
}

var SMTP *SMTPConfig

func loadSMTPConfig() {

	SMTP = &SMTPConfig{}

	err := envconfig.Process("smtp", SMTP)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
