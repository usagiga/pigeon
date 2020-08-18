package infra

import "github.com/usagiga/pigeon/model"

type EsaInfra interface {
	GetArticle(id int) (article *model.Article, err error)
}

type ImageInfra interface {
	Fetch(dstPath, srcUrl string) (err error)
}
