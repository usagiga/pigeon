package infra

import (
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"os"
	"path"
)

type GitInfraImpl struct{
	 projectRootDir string
}

func NewGitInfra(baseDir, repoUrl string) (infra GitInfra, err error) {
	err = runGit(baseDir, "clone", repoUrl)
	if err != nil {
		return nil, xerrors.Errorf("Can't run `git clone`", err)
	}

	repoDir, err := getRepoDir(baseDir, repoUrl)
	if err != nil {
		return nil, xerrors.Errorf("Can't get repo dir", err)
	}

	return &GitInfraImpl{
		projectRootDir: repoDir,
	}, nil
}

func (i *GitInfraImpl) CommitUnStaged(message string) (err error) {
	err = runGit(i.projectRootDir, "add", ".")
	if err != nil {
		return xerrors.Errorf("Can't run `git add`", err)
	}

	err = runGit(i.projectRootDir, "commit", "-m", message)
	if err != nil {
		return xerrors.Errorf("Can't run `git commit`", err)
	}

	return nil
}

func (i *GitInfraImpl) Push() (err error) {
	err = runGit(i.projectRootDir, "push", "origin", "head")
	if err != nil {
		return xerrors.Errorf("Can't run `git commit`", err)
	}

	return nil
}

func runGit(baseDir string, args ...string) (err error) {
	proc, err := os.StartProcess("git", args, &os.ProcAttr{
		Dir: baseDir,
	})
	if err != nil {
		return xerrors.Errorf("Can't start `git` process: %w", err)
	}

	_, err = proc.Wait()
	if err != nil {
		return xerrors.Errorf("Can't wait `git` process: %w", err)
	}

	return nil
}

func getRepoDir(baseDir, repoUrl string) (repoDir string, err error) {
	repoLastNode, err := urlnode.GetLastNodeFromString(repoUrl)
	if err != nil {
		return "", xerrors.Errorf("Can't split repo url: %w", err)
	}
	repoDir = path.Join(baseDir, urlnode.NodeWithoutExt(repoLastNode))

	return repoDir, nil
}
