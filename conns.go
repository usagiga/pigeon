package main

import (
	"github.com/hiroakis/esa-go"
	"github.com/usagiga/pigeon/model"
)

func ConnectToEsa(config *model.Config) (esaClient *esa.EsaClient) {
	return esa.NewEsaClient(config.EsaAPIKey, config.EsaTeam)
}
