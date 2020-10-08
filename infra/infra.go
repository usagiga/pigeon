package infra

import "github.com/usagiga/pigeon/model"

type EsaInfra interface {
	GetArticle(id int) (article *model.Article, err error)
}

type ImageStorageInfra interface {
	Fetch(repoDir *model.GitRepoDir, srcUrl string) (skipped bool, err error)
	Exists(repoDir *model.GitRepoDir, fileName string) (exists bool, err error)
}

type GitInfra interface {
	Clone(baseDir, repoUrl string) (projectRootDir string, err error)
	CommitUnStaged(projectRootDir, message string) (err error)
	Push(projectRootDir string) (err error)
}
