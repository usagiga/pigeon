package model

import (
	"golang.org/x/xerrors"
	"path"
	"strings"
)

type GitRepoDir struct {
	ProjectRootDir        string
	relativeArticleDir    string
	relativeImageStoreDir string
	relativeImageViewDir  string
}

func NewGitRepoDir(projectRootDir, articleDir, imageDir, imageViewDir string) (dir *GitRepoDir, err error) {
	if path.IsAbs(articleDir) || path.IsAbs(imageDir) {
		return nil, xerrors.New("article or image dir must be absolute path")
	}
	if strings.HasPrefix(articleDir, "~") || strings.HasPrefix(imageDir, "~") {
		return nil, xerrors.New(`article or image dir must NOT be beginning from "~"`)
	}

	return &GitRepoDir{
		ProjectRootDir:        projectRootDir,
		relativeArticleDir:    articleDir,
		relativeImageStoreDir: imageDir,
		relativeImageViewDir:  imageViewDir,
	}, nil
}

func (m *GitRepoDir) ArticleDir() (dir string) {
	return path.Join(m.ProjectRootDir, m.relativeArticleDir)
}

func (m *GitRepoDir) RelativeArticleDir() (dir string) {
	return m.relativeArticleDir
}

func (m *GitRepoDir) ImageDir() (dir string) {
	return path.Join(m.ProjectRootDir, m.relativeImageStoreDir)
}

func (m *GitRepoDir) RelativeImageStoreDir() (dir string) {
	return m.relativeImageStoreDir
}

func (m *GitRepoDir) RelativeImageViewDir() (dir string) {
	return m.relativeImageViewDir
}
