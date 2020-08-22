package main

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/hiroakis/esa-go"
	"github.com/usagiga/pigeon/model"
	"golang.org/x/xerrors"
)

func ConnectToEsa(config *model.Config) (esaClient *esa.EsaClient) {
	return esa.NewEsaClient(config.EsaAPIKey, config.EsaTeam)
}

func ConnectToStorage(config *model.Config) (storageClient *storage.Client, err error) {
	if config.ProjectID == "" || config.BucketID == "" {
		return nil, xerrors.Errorf("can't initialize GCS client. there's no project or bucket id")
	}

	storageClient, err = storage.NewClient(context.TODO())
	if err != nil {
		return nil, xerrors.Errorf("can't initialize GCS client: %w", err)
	}

	return storageClient, err
}
