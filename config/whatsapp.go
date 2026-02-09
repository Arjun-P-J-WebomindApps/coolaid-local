package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type WhatsappConfig struct {
	ApiVersion    string `split_words:"true" json:"WHATSAPP_API_VERSION"`
	AccessToken   string `split_words:"true" json:"WHATSAPP_ACCESS_TOKEN"`
	PhoneNumberID string `split_words:"true" json:"WHATSAPP_PHONE_NUMBER_ID"`
}

var Whatsapp *WhatsappConfig

/* Deprecated*/
func loadWhatsappConfig() {
	Whatsapp = &WhatsappConfig{}
	err := envconfig.Process("whatsapp", Whatsapp)

	if err != nil {
		log.Fatalf("Error loading Whatsapp env %s", err)
	}

}
