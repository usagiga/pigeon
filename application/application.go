package application

import "github.com/usagiga/pigeon/model"

type TransferApplication interface {
	TransferArticle(config *model.Config, postId int)
}
