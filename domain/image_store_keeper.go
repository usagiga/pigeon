package domain

import (
	"github.com/usagiga/pigeon/infra"
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"path"
)

type ImageStoreKeeperUseCaseImpl struct {
	imageInfra infra.ImageInfra
}

func NewImageStoreKeeperUseCase(imageInfra infra.ImageInfra) (domain ImageStoreKeeperUseCase) {
	return &ImageStoreKeeperUseCaseImpl{
		imageInfra: imageInfra,
	}
}

func (d *ImageStoreKeeperUseCaseImpl) Store(repoDir *model.GitRepoDir, url string) (storedUrl string, err error) {
	// Download specified image
	err = d.imageInfra.Fetch(repoDir.ImageDir(), url)
	if err != nil {
		return "", xerrors.Errorf("can't download image(URL: %s) : %w", url, err)
	}

	// Get file name from URL
	fileName, err := urlnode.GetLastNodeFromString(url)
	if err != nil {
		return "", xerrors.Errorf("can't get file name from URL(URL: %s) : %w", url, err)
	}

	return path.Join(repoDir.RelativeImageViewDir(), fileName), nil
}
