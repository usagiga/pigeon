package domain

import (
	"github.com/usagiga/pigeon/infra"
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"net/url"
	"path"
)

type ImageStoreKeeperUseCaseImpl struct {
	imageInfra infra.ImageStorageInfra
}

func NewImageStoreKeeperUseCase(imageInfra infra.ImageStorageInfra) (domain ImageStoreKeeperUseCase) {
	return &ImageStoreKeeperUseCaseImpl{
		imageInfra: imageInfra,
	}
}

func (d *ImageStoreKeeperUseCaseImpl) Store(repoDir *model.GitRepoDir, srcUrl string) (storedUrl string, err error) {
	// Download specified image
	_, err = d.imageInfra.Fetch(repoDir, srcUrl)
	if err != nil {
		return "", xerrors.Errorf("can't download image(URL: %s) : %w", srcUrl, err)
	}

	// Get file name from URL
	fileName, err := urlnode.GetLastNodeFromString(srcUrl)
	if err != nil {
		return "", xerrors.Errorf("can't get file name from URL(URL: %s) : %w", srcUrl, err)
	}

	// Return storedUrl
	imageViewBaseDir := repoDir.RelativeImageViewDir()
	urlInfo, err := url.Parse(imageViewBaseDir)
	if err != nil {
		// imageViewBaseDir is not URL, so, treat it as filepath
		return path.Join(imageViewBaseDir, fileName), nil
	}

	urlInfo.Path = path.Join(urlInfo.Path, fileName)

	return urlInfo.String(), nil
}
