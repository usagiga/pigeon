package infra

import (
	"github.com/usagiga/pigeon/model"
	"os"
	"path"
	"testing"
)

var (
	tmpGitRepoDir *model.GitRepoDir
)

const (
	tmpFile = "existing.png"
)

func initTestFile(t *testing.T) {
	tmpDir := os.TempDir()
	tmpGitRepoDir, _ = model.NewGitRepoDir(tmpDir, "test", "test", "test")

	err := os.MkdirAll(tmpGitRepoDir.ImageDir(), 0666)
	if err != nil {
		t.Fatalf("Can't generate tmp dir: %v", err)
	}

	testFilePath := path.Join(tmpGitRepoDir.ImageDir(), tmpFile)
	_, err = os.Create(testFilePath)
	if err != nil {
		t.Fatalf("Can't generate tmp file: %v", err)
	}
}

func finalizeTestFile(t *testing.T) {
	err := os.RemoveAll(tmpGitRepoDir.ImageDir())
	if err != nil {
		t.Errorf("Can't remove test files: %v", err)
	}
}

func TestImageFileStorageInfraImpl_Exists(t *testing.T) {
	// Initialize
	initTestFile(t)
	defer finalizeTestFile(t)

	imageInfra := NewImageFileStorageInfra()

	// Declare test cases
	type Arg struct {
		repoDir  *model.GitRepoDir
		fileName string
	}

	type Result struct {
		exists bool
	}

	testCases := []struct {
		isExpectedError bool
		arg             Arg
		result          Result
	}{
		// Nominal scenario (existing)
		{
			isExpectedError: false,
			arg: Arg{
				repoDir:  tmpGitRepoDir,
				fileName: "existing.png",
			},
			result: Result{
				exists: true,
			},
		},
		// Nominal scenario (not existing)
		{
			isExpectedError: false,
			arg: Arg{
				repoDir:  tmpGitRepoDir,
				fileName: "not-existing.png",
			},
			result: Result{
				exists: false,
			},
		},
	}

	// Run test
	for i, v := range testCases {
		caseNum := i + 1
		exists, err := imageInfra.Exists(v.arg.repoDir, v.arg.fileName)

		// When raising NOT expected error
		if err != nil && !v.isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && v.isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When returns NOT expected result
		if exists != v.result.exists {
			t.Errorf(
				"Case %d: Returned 'exists' is not expected.\nExpected:\t%v\nActual:\t%v",
				caseNum,
				v.result.exists,
				exists,
			)
		}
	}
}

func TestImageFileStorageInfraImpl_Fetch(t *testing.T) {
	// Initialize
	initTestFile(t)
	defer finalizeTestFile(t)

	imageInfra := NewImageFileStorageInfra()

	// Declare test cases
	type Arg struct {
		repoDir *model.GitRepoDir
		srcUrl  string
	}
	type Result struct {
		skipped bool
	}

	testCases := []struct {
		isExpectedError bool
		arg             Arg
		result          Result
	}{
		// Nominal scenario
		{
			isExpectedError: false,
			arg: Arg{
				repoDir: tmpGitRepoDir,
				srcUrl:  "https://example.com/index.html",
			},
			result: Result{
				skipped: false,
			},
		},
		// Nominal scenario (Skipped)
		{
			isExpectedError: false,
			arg: Arg{
				repoDir: tmpGitRepoDir,
				srcUrl:  "https://example.com/index.html",
			},
			result: Result{
				skipped: true,
			},
		},
		// On error on fetch image
		{
			isExpectedError: true,
			arg: Arg{
				repoDir: tmpGitRepoDir,
				srcUrl:  "error",
			},
			result: Result{
				skipped: false,
			},
		},
	}

	// Run test
	for i, v := range testCases {
		caseNum := i + 1
		skipped, err := imageInfra.Fetch(v.arg.repoDir, v.arg.srcUrl)

		// When raising NOT expected error
		if err != nil && !v.isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && v.isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When returns NOT expected result
		if skipped != v.result.skipped {
			t.Errorf(
				"Case %d: Returned 'skipped' is not expected.\nExpected:\t%v\nActual:\t%v",
				caseNum,
				v.result.skipped,
				skipped,
			)
		}
	}
}
