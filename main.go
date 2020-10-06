package main

import (
	"github.com/usagiga/envs-go"
	"github.com/usagiga/pigeon/application"
	"github.com/usagiga/pigeon/domain"
	"github.com/usagiga/pigeon/infra"
	"github.com/usagiga/pigeon/model"
	"log"
)

func main() {
	// Parse args
	postId := ParseArgs()

	// Load config from envs
	config := &model.Config{}
	if err := envs.Load(config); err != nil {
		log.Fatalf("Can't load config: %+v", err)
	}

	// Initialize esa.io client
	esaClient := ConnectToEsa(config)

	// Initialize infra
	esaInfra := infra.NewEsaInfra(esaClient)
	gitInfra := infra.NewGitInfra(config)

	var imageStorageInfra infra.ImageStorageInfra
	switch config.GetStoreImageMode() {
	case model.None:
		imageStorageInfra = nil
	case model.File:
		imageStorageInfra = infra.NewImageFileStorageInfra()
	case model.GCS:
		// Initialize GCS client
		gcsClient, err := ConnectToStorage(config)
		if err != nil {
			log.Fatalf("Can't connect GCS. Stopping to launch: %+v", err)
		}

		imageStorageInfra = infra.NewImageGCSStorageInfra(config.BucketID, gcsClient)
	}

	// Initialize Domain
	var imageStoreKeeperUseCase domain.ImageStoreKeeperUseCase
	switch config.GetStoreImageMode() {
	case model.None:
		imageStoreKeeperUseCase = domain.NewNOPStoreKeeperUseCase()
	default:
		imageStoreKeeperUseCase = domain.NewImageStoreKeeperUseCase(imageStorageInfra)
	}

	gitRepoUseCase := domain.NewGitRepositoryUseCase(gitInfra)
	articleBuilderUseCase := domain.NewArticleBuilderUseCase(imageStoreKeeperUseCase, esaInfra)

	// Initialize Application
	transferApplication := application.NewTransferApplication(articleBuilderUseCase, gitRepoUseCase)

	// Run
	transferApplication.TransferArticle(config, postId)
}
