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

	// Initialize GCS client
	gcsClient, err := ConnectToStorage(config)
	if err != nil {
		log.Fatalf("Can't connect GCS. Stopping to launch: %+v", err)
	}

	// Initialize infra
	esaInfra := infra.NewEsaInfra(esaClient)
	gitInfra := infra.NewGitInfra(config)
	imageInfra := infra.NewImageStorageInfra(config.BucketID, gcsClient)

	// Initialize Domain
	gitRepoUseCase := domain.NewGitRepositoryUseCase(gitInfra)
	imageStoreKeeperUseCase := domain.NewImageStoreKeeperUseCase(imageInfra)
	articleBuilderUseCase := domain.NewArticleBuilderUseCase(imageStoreKeeperUseCase, esaInfra)

	// Initialize Application
	transferApplication := application.NewTransferApplication(articleBuilderUseCase, gitRepoUseCase)

	// Run
	transferApplication.TransferArticle(config, postId)
}
