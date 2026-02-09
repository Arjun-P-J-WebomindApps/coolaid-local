package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AuthConfig struct {
	OtpExpiryMinutes         int    `split_words:"true" json:"AUTH_OTP_EXPIRY_MINUTES"`
	AccessTokenExpiryMinutes int    `split_words:"true" json:"AUTH_ACCESS_TOKEN_EXPIRY_MINUTES"`
	RefreshTokenExpiryHours  int    `split_words:"true" json:"AUTH_REFRESH_TOKEN_EXPIRY_HOURS"`
	SessionExpiryHours       int    `split_words:"true" json:"AUTH_SESSION_EXPIRY_HOURS"`
	JwtSecret                string `split_words:"true" json:"AUTH_JWT_SECRET"`
}

var Auth *AuthConfig

func loadAuthConfig() {

	Auth = &AuthConfig{}

	err := envconfig.Process("auth", Auth)
	if err != nil {
		log.Fatalf("auth config invalid fields %s", err.Error())
	}

}
