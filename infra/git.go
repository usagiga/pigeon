package infra

import (
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"os/exec"
	"path"
)

type GitInfraImpl struct{
	gitBinPath string
}

func NewGitInfra(config *model.Config) (infra GitInfra) {
	return &GitInfraImpl{
		gitBinPath: config.GitBinPath,
	}
}

func (i *GitInfraImpl) Clone(baseDir, repoUrl string) (projectRootDir string, err error) {
	err = i.runGit(baseDir, "clone", repoUrl)
	if err != nil {
		return "", xerrors.Errorf("Can't run `git clone`: %w", err)
	}

	repoDir, err := i.getRepoDir(baseDir, repoUrl)
	if err != nil {
		return "", xerrors.Errorf("Can't get repo dir: %w", err)
	}

	return repoDir, nil
}

func (i *GitInfraImpl) CommitUnStaged(projectRootDir, message string) (err error) {
	err = i.runGit(projectRootDir, "add", ".")
	if err != nil {
		return xerrors.Errorf("Can't run `git add`: %w", err)
	}

	err = i.runGit(projectRootDir, "commit", "-m", message)
	if err != nil {
		return xerrors.Errorf("Can't run `git commit`: %w", err)
	}

	return nil
}

func (i *GitInfraImpl) Push(projectRootDir string) (err error) {
	err = i.runGit(projectRootDir, "push", "origin", "HEAD")
	if err != nil {
		return xerrors.Errorf("Can't run `git commit`: %w", err)
	}

	return nil
}

func (i *GitInfraImpl) runGit(baseDir string, args ...string) (err error) {
	cmd := exec.Command(i.gitBinPath, args...)
	cmd.Dir = baseDir

	err = cmd.Run()
	if err != nil {
		return xerrors.Errorf("Can't start `git` process: %w", err)
	}

	return nil
}

func (i *GitInfraImpl) getRepoDir(baseDir, repoUrl string) (repoDir string, err error) {
	repoLastNode, err := urlnode.GetLastNodeFromString(repoUrl)
	if err != nil {
		return "", xerrors.Errorf("Can't split repo url: %w", err)
	}
	repoDir = path.Join(baseDir, urlnode.NodeWithoutExt(repoLastNode))

	return repoDir, nil
}
