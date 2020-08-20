package domain

import (
	"github.com/usagiga/pigeon/infra"
	"github.com/usagiga/pigeon/model"
	"golang.org/x/xerrors"
	"os"
)

type GitRepositoryUseCaseImpl struct {
	gitInfra infra.GitInfra
}

func NewGitRepositoryUseCase() (domain GitRepositoryUseCase) {
	return &GitRepositoryUseCaseImpl{}
}

func (d *GitRepositoryUseCaseImpl) Initialize(repoUrl, articleDir, imageDir string) (repoDir *model.GitRepoDir, err error) {
	// Clone
	dstDir := os.TempDir()
	projectRootDir, err := d.gitInfra.Clone(dstDir, repoUrl)
	if err != nil {
		return nil, xerrors.Errorf("can't clone target repository(DestDir: %s, RepoURL: %s) : %w", dstDir, repoUrl, err)
	}

	// Return its dir info
	repoDir, err = model.NewGitRepoDir(projectRootDir, articleDir, imageDir)
	if err != nil {
		return nil, xerrors.Errorf("specified article or image dir aren't suited for it(ArticleDir: %s, ImageDir: %s) : %w", articleDir, imageDir, err)
	}

	return repoDir, nil
}

func (d *GitRepositoryUseCaseImpl) CommitAndPush(repoDir *model.GitRepoDir) (err error) {
	repoRootDir := repoDir.ProjectRootDir

	// Commit
	err = d.gitInfra.CommitUnStaged(repoRootDir, "pigeon auto post")
	if err != nil {
		return xerrors.Errorf("can't commit changes : %w", err)
	}

	// Push
	err = d.gitInfra.Push(repoDir.ProjectRootDir)
	if err != nil {
		return xerrors.Errorf("can't push commits : %w", err)
	}

	return nil
}

func (d *GitRepositoryUseCaseImpl) Dispose(repoDir *model.GitRepoDir) (err error) {
	repoRootDir := repoDir.ProjectRootDir

	err = os.RemoveAll(repoRootDir)
	if err != nil {
		return xerrors.Errorf("can't remove project root dir(Dir:%s) : %w", repoRootDir, err)
	}

	return nil
}
