package domain

import (
	"github.com/usagiga/pigeon/model"
)

type NOPStoreKeeperUseCaseImpl struct{}

func NewNOPStoreKeeperUseCase() (domain ImageStoreKeeperUseCase) {
	return &NOPStoreKeeperUseCaseImpl{}
}

func (d *NOPStoreKeeperUseCaseImpl) Store(repoDir *model.GitRepoDir, srcUrl string) (storedUrl string, err error) {
	return srcUrl, nil
}
