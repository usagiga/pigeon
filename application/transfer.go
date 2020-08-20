package application

import (
	"github.com/usagiga/pigeon/domain"
	"github.com/usagiga/pigeon/model"
	"log"
)

type TransferApplicationImpl struct {
	articleBuilderUseCase domain.ArticleBuilderUseCase
	gitRepositoryUseCase  domain.GitRepositoryUseCase
}

func NewTransferApplication(
	articleBuilderUseCase domain.ArticleBuilderUseCase,
	gitRepositoryUseCase domain.GitRepositoryUseCase,
) (app TransferApplication) {
	return &TransferApplicationImpl{
		articleBuilderUseCase: articleBuilderUseCase,
		gitRepositoryUseCase:  gitRepositoryUseCase,
	}
}

func (a *TransferApplicationImpl) TransferArticle(config *model.Config, postId int) {
	// Initialize git repository
	repoDir, err := a.gitRepositoryUseCase.Initialize(config.DiaryRepoURL, config.ArticleDir, config.ImageDir)
	if err != nil {
		log.Fatalf("Can't initialize git repository: %+v", err)
	}

	defer func() {
		err = a.gitRepositoryUseCase.Dispose(repoDir)
		if err != nil {
			log.Fatalf("Can't dispose git repository: %+v", err)
		}
	}()

	// Get raw article
	rawArticle, err := a.articleBuilderUseCase.GetRawArticle(postId)
	if err != nil {
		log.Fatalf("Can't get article: %+v", err)
	}

	// Format article
	// (replace variables, image urls hosted on esa.io, ...)
	formattedArticle, err := a.articleBuilderUseCase.FormatArticle(repoDir, rawArticle)
	if err != nil {
		log.Fatalf("Can't format article: %+v", err)
	}

	// Store article
	err = a.articleBuilderUseCase.Store(repoDir, formattedArticle)
	if err != nil {
		log.Fatalf("Can't store article: %+v", err)
	}

	// Push
	err = a.gitRepositoryUseCase.CommitAndPush(repoDir)
	if err != nil {
		log.Fatalf("Can't push this changes: %+v", err)
	}
}
