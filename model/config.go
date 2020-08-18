package model

type Config struct {
	EsaAPIKey string `env:"PIGEON_ESA_KEY"`
	EsaTeam   string `env:"PIGEON_ESA_TEAM"`
}
