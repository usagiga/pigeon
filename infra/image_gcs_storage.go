package infra

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"net/http"
)

type ImageGCSStorageInfraImpl struct {
	bucketName string
	gcsClient  *storage.Client
}

func NewImageGCSStorageInfra(bucketName string, gcsClient *storage.Client) (infra ImageStorageInfra) {
	return &ImageGCSStorageInfraImpl{
		bucketName: bucketName,
		gcsClient: gcsClient,
	}
}

func (i *ImageGCSStorageInfraImpl) Fetch(repoDir *model.GitRepoDir, srcUrl string) (skipped bool, err error) {
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

func (i *ImageGCSStorageInfraImpl) fetch(srcUrl string) (imageBytes []byte, err error) {
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

func (i *ImageGCSStorageInfraImpl) storeIntoFile(_ *model.GitRepoDir, fileName string, imageBytes []byte) (err error) {
	br := bytes.NewReader(imageBytes)

	wc := i.gcsClient.Bucket(i.bucketName).Object(fileName).NewWriter(context.TODO())
	if _, err = io.Copy(wc, br); err != nil {
		return xerrors.Errorf("can't write image into stream: %w", err)
	}
	if err := wc.Close(); err != nil {
		return xerrors.Errorf("can't close stream: %w", err)
	}

	return nil
}

func (i *ImageGCSStorageInfraImpl) Exists(_ *model.GitRepoDir, fileName string) (exists bool, err error) {
	_, err = i.gcsClient.Bucket(i.bucketName).Object(fileName).Attrs(context.TODO())
	if err == storage.ErrObjectNotExist {
		return false, nil
	}
	if err != nil {
		return false, xerrors.Errorf("can't get GCS object attrs: %w", err)
	}

	return true, nil
}
