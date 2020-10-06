package infra

import (
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type ImageFileStorageInfraImpl struct {
}

func NewImageFileStorageInfra() (infra ImageStorageInfra) {
	return &ImageFileStorageInfraImpl{}
}

func (i *ImageFileStorageInfraImpl) Fetch(repoDir *model.GitRepoDir, srcUrl string) (skipped bool, err error) {
	// Get name from URL
	fileName, err := urlnode.GetLastNodeFromString(srcUrl)
	if err != nil {
		return false, xerrors.Errorf("Can't get file name from URL(URL: %s): %w", srcUrl, err)
	}

	// Check redundant upload
	exists, err := i.Exists(repoDir, fileName)
	if err != nil {
		return false, xerrors.Errorf("Can't check image has already uploaded (Name: %s): %w", fileName, err)
	}
	if exists {
		return true, nil
	}

	// Fetch from URL
	imageBytes, err := i.fetch(srcUrl)
	if err != nil {
		return false, xerrors.Errorf("Can't fetch file from URL(URL: %s): %w", srcUrl, err)
	}

	// Save into dir
	err = i.storeIntoFile(repoDir, fileName, imageBytes)
	if err != nil {
		return false, xerrors.Errorf("Can't store file(Name: %s): %w", fileName, err)
	}

	return false, nil
}

func (i *ImageFileStorageInfraImpl) fetch(srcUrl string) (imageBytes []byte, err error) {
	res, err := http.Get(srcUrl)
	if err != nil {
		return nil, xerrors.Errorf("Can't download image(URL: %s): %w", srcUrl, err)
	}

	//noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	imageBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.Errorf("Can't read response body stream(URL: %s): %w", srcUrl, err)
	}

	return imageBytes, nil
}

func (i *ImageFileStorageInfraImpl) storeIntoFile(repoDir *model.GitRepoDir, fileName string, imageBytes []byte) (err error) {
	filePath := path.Join(repoDir.ImageDir(), fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return xerrors.Errorf("Can't create file(%s): %w", filePath, err)
	}

	//noinspection GoUnhandledErrorResult
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		return xerrors.Errorf("Can't write data into file(%s): %w", filePath, err)
	}

	return nil
}

func (i *ImageFileStorageInfraImpl) Exists(repoDir *model.GitRepoDir, fileName string) (exists bool, err error) {
	filePath := path.Join(repoDir.ImageDir(), fileName)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, xerrors.Errorf("Can't detect file existence(%s): %w", filePath, err)
	}

	return true, nil
}
