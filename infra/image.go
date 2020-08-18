package infra

import (
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type ImageInfraImpl struct{}

func NewImageInfra() (infra ImageInfra) {
	return &ImageInfraImpl{}
}

func (i *ImageInfraImpl) Fetch(dstDir, srcUrl string) (err error) {
	// Fetch from URL
	imageBytes, err := i.fetch(srcUrl)
	if err != nil {
		return xerrors.Errorf("Can't fetch file from URL(URL: %s): %w", srcUrl, err)
	}

	// Get name from URL
	fileName, err := urlnode.GetLastNodeFromString(srcUrl)
	if err != nil {
		return xerrors.Errorf("Can't get file name from URL(URL: %s): %w", srcUrl, err)
	}

	// Save into dir
	dstPath := path.Join(dstDir, fileName)
	err = i.storeIntoFile(dstPath, imageBytes)
	if err != nil {
		return xerrors.Errorf("Can't store file(Path: %s): %w", dstPath, err)
	}

	return nil
}

func (i *ImageInfraImpl) fetch(srcUrl string) (imageBytes []byte, err error) {
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

func (i *ImageInfraImpl) storeIntoFile(dstPath string, imageBytes []byte) (err error) {
	file, err := os.Create(dstPath)
	if err != nil {
		return xerrors.Errorf("Can't create file(%s): %w", dstPath, err)
	}

	//noinspection GoUnhandledErrorResult
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		return xerrors.Errorf("Can't write data into file(%s): %w", dstPath, err)
	}

	return nil
}
