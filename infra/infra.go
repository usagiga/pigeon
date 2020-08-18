package infra

import "github.com/usagiga/pigeon/model"

type EsaInfra interface {
	GetArticle(id int) (article *model.Article, err error)
}

type ImageInfra interface {
	Fetch(dstPath, srcUrl string) (err error)
}

type GitInfra interface {
	CommitUnStaged(message string) (err error)
	Push() (err error)
}
