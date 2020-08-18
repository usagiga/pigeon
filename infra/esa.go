package infra

import (
	"github.com/hiroakis/esa-go"
	"github.com/usagiga/pigeon/model"
	"golang.org/x/xerrors"
)

type EsaInfraImpl struct {
	client *esa.EsaClient
}

func NewEsaInfra(client *esa.EsaClient) (infra EsaInfra) {
	return &EsaInfraImpl{client: client}
}

func (i *EsaInfraImpl) GetArticle(id int) (article *model.Article, err error) {
	post, err := i.client.GetPost(id)
	if err != nil {
		return nil, xerrors.Errorf("Can't get post through API: %w", err)
	}

	article = &model.Article{
		Title:    post.Name,
		Contents: post.BodyMd,
	}

	return article, nil
}
