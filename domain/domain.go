package domain

import "github.com/usagiga/pigeon/model"

type GitRepositoryUseCase interface {
	Initialize(repoUrl, articleDir, imageDir string) (repoDir *model.GitRepoDir, err error)
	CommitAndPush(repoDir *model.GitRepoDir) (err error)
	Dispose(repoDir *model.GitRepoDir) (err error)
}

type ArticleBuilderUseCase interface {
	GetRawArticle(postId int) (rawArticle *model.Article, err error)
	FormatArticle(repoDir *model.GitRepoDir, rawArticle *model.Article) (formattedArticle *model.Article, err error)
	Store(repoDir *model.GitRepoDir, article *model.Article) (err error)
}

type ImageStoreKeeperUseCase interface {
	Store(repoDir *model.GitRepoDir, url string) (storedUrl string, err error)
}
